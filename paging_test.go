package coc_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"sync"
	"testing"

	coc "github.com/natebabyak/clashofclans-go"
)

func TestIteratorPagination(t *testing.T) {
	for _, backward := range []bool{false, true} {
		t.Run(fmt.Sprint("backward=", backward), func(t *testing.T) {
			calls := 0
			direction, opposite := "after", "before"
			if backward {
				direction, opposite = opposite, direction
			}
			client := testClient(t, func(r *http.Request) (*http.Response, error) {
				calls++
				q := r.URL.Query()
				if q.Get("name") != "Example" || q.Get("labelIds") != "1,2" || q.Get("limit") != "2" {
					t.Errorf("filters lost: %v", q)
				}
				wantCursor := ""
				if backward {
					wantCursor = "start"
				}
				if calls == 2 {
					wantCursor = "next+/=&"
				}
				if q.Get(direction) != wantCursor || q.Get(opposite) != "" {
					t.Errorf("wrong cursor: %v", q)
				}
				if calls > 2 {
					t.Fatal("iterator failed to stop")
				}
				if calls == 1 {
					return response(200, fmt.Sprintf(`{"items":[{"name":"one"},{"name":"two"}],"paging":{"cursors":{%q:"next+/=&"}}}`, direction)), nil
				}
				return response(200, `{"items":[{"name":"three"}]}`), nil
			})
			opts := []coc.ListOption{coc.Limit(2)}
			if backward {
				opts = append(opts, coc.Before("start"))
			}
			sequence := client.Clans(coc.ClanFilter{Name: "Example", LabelIDs: []int{1, 2}}, opts...)
			if calls != 0 {
				t.Fatal("iterator fetched before being used")
			}
			var names []string
			for clan, err := range sequence {
				if err != nil {
					t.Fatal(err)
				}
				names = append(names, clan.Name)
			}
			if !reflect.DeepEqual(names, []string{"one", "two", "three"}) || calls != 2 {
				t.Fatalf("items=%v, calls=%d", names, calls)
			}
		})
	}
}

func TestIteratorEarlyStop(t *testing.T) {
	calls := 0
	client := testClient(t, func(*http.Request) (*http.Response, error) {
		calls++
		return response(200, `{"items":[{"name":"one"},{"name":"two"}],"paging":{"cursors":{"after":"next"}}}`), nil
	})
	for _, err := range client.ClanMembers("#2PP") {
		if err != nil {
			t.Fatal(err)
		}
		break
	}
	if calls != 1 {
		t.Fatalf("early break made %d calls", calls)
	}
}

func TestIteratorErrors(t *testing.T) {
	t.Run("later page", func(t *testing.T) {
		calls, items, failures := 0, 0, 0
		client := testClient(t, func(*http.Request) (*http.Response, error) {
			calls++
			if calls == 1 {
				return response(200, `{"items":[{"name":"one"}],"paging":{"cursors":{"after":"next"}}}`), nil
			}
			return response(429, `{"reason":"requestThrottled"}`), nil
		})
		for item, err := range client.PlayerRankings(coc.Global) {
			if err != nil {
				var apiErr *coc.APIError
				if !errors.As(err, &apiErr) || apiErr.StatusCode != 429 || item.Name != "" {
					t.Fatalf("item=%v, err=%v", item, err)
				}
				failures++
			} else {
				items++
			}
		}
		if calls != 2 || items != 1 || failures != 1 {
			t.Fatalf("calls=%d, items=%d, failures=%d", calls, items, failures)
		}
	})
	t.Run("cursor cycle", func(t *testing.T) {
		calls, failures := 0, 0
		client := testClient(t, func(*http.Request) (*http.Response, error) {
			calls++
			if calls > 3 {
				t.Fatal("cursor loop was not stopped")
			}
			cursor := "A"
			if calls == 2 {
				cursor = "B"
			}
			return response(200, fmt.Sprintf(`{"items":[],"paging":{"cursors":{"after":%q}}}`, cursor)), nil
		})
		for _, err := range client.Locations() {
			if err == nil {
				t.Fatal("unexpected item")
			}
			failures++
		}
		if failures != 1 || calls != 3 {
			t.Fatalf("calls=%d, failures=%d", calls, failures)
		}
	})
	t.Run("invalid options", func(t *testing.T) {
		client := testClient(t, func(*http.Request) (*http.Response, error) {
			t.Fatal("invalid iterator made a request")
			return nil, nil
		})
		failures := 0
		for _, err := range client.Locations(coc.Limit(0)) {
			if err == nil {
				t.Fatal("missing option error")
			}
			failures++
		}
		if failures != 1 {
			t.Fatalf("failures=%d", failures)
		}
	})
	t.Run("cancellation between items", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		calls, items, failures := 0, 0, 0
		client := testClient(t, func(*http.Request) (*http.Response, error) {
			calls++
			return response(200, `{"items":[{"name":"one"},{"name":"two"}],"paging":{"cursors":{"after":"next"}}}`), nil
		})
		for _, err := range client.Locations(coc.WithContext(ctx)) {
			if err != nil {
				if !errors.Is(err, context.Canceled) {
					t.Fatal(err)
				}
				failures++
			} else {
				items++
				cancel()
			}
		}
		if calls != 1 || items != 1 || failures != 1 {
			t.Fatalf("calls=%d, items=%d, failures=%d", calls, items, failures)
		}
	})
}

func TestDefaultPagingAndSharedOptions(t *testing.T) {
	client := testClient(t, func(r *http.Request) (*http.Response, error) {
		if len(r.URL.Query()) != 0 {
			t.Errorf("default query = %v", r.URL.Query())
		}
		return response(200, `{"items":[]}`), nil
	})
	if _, err := client.GetLocations(); err != nil {
		t.Fatal(err)
	}
	// Options may be reused concurrently without sharing mutable cursor state.
	client = testClient(t, func(r *http.Request) (*http.Response, error) {
		if r.URL.Query().Get("limit") != "2" {
			t.Errorf("limit lost: %v", r.URL.Query())
		}
		if r.URL.Query().Get("after") == "" {
			return response(200, `{"items":[{"name":"one"}],"paging":{"cursors":{"after":"next"}}}`), nil
		}
		return response(200, `{"items":[{"name":"two"}]}`), nil
	})
	opts := []coc.ListOption{coc.Limit(2)}
	seq := client.Locations(opts...)
	var wg sync.WaitGroup
	for range 2 {
		wg.Go(func() {
			items := 0
			for _, err := range seq {
				if err != nil {
					t.Error(err)
					return
				}
				items++
			}
			if items != 2 {
				t.Errorf("items=%d", items)
			}
		})
	}
	wg.Wait()
}
