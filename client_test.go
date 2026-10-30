package coc_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	coc "github.com/natebabyak/clashofclans-go"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func response(status int, body string) *http.Response {
	return &http.Response{StatusCode: status, Header: make(http.Header), Body: io.NopCloser(strings.NewReader(body))}
}

func testClient(t *testing.T, transport roundTripFunc) *coc.Client {
	t.Helper()
	client, err := coc.NewClient("test-token", coc.WithBaseURL("https://example.invalid/v1/"), coc.WithHTTPClient(&http.Client{Transport: transport}))
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func discard[T any](_ *T, err error) error { return err }

func TestRouteAccess(t *testing.T) {
	tests := []struct {
		name, path, query string
		call              func(*coc.Client) error
	}{
		{"clan", "/clans/%232PP", "", func(c *coc.Client) error { return discard(c.GetClan(" 2pp ")) }},
		{"members", "/clans/%232PP/members", "limit=10", func(c *coc.Client) error { return discard(c.GetClanMembers("#2PP", coc.Limit(10))) }},
		{"war", "/clans/%232PP/currentwar", "", func(c *coc.Client) error { return discard(c.GetClanCurrentWar("2PP")) }},
		{"war log", "/clans/%232PP/warlog", "limit=10", func(c *coc.Client) error { return discard(c.GetClanWarLog("2PP", coc.Limit(10))) }},
		{"CWL group", "/clans/%232PP/currentwar/leaguegroup", "", func(c *coc.Client) error { return discard(c.GetClanCurrentWarLeagueGroup("2PP")) }},
		{"CWL war", "/clanwarleagues/wars/%232PP", "", func(c *coc.Client) error { return discard(c.GetClanWarLeagueWar("2PP")) }},
		{"raids", "/clans/%232PP/capitalraidseasons", "limit=10", func(c *coc.Client) error { return discard(c.GetClanCapitalRaidSeasons("2PP", coc.Limit(10))) }},
		{"search", "/clans", "name=Example&limit=10", func(c *coc.Client) error {
			return discard(c.SearchClans(coc.ClanFilter{Name: "Example"}, coc.Limit(10)))
		}},
		{"player", "/players/%232PP", "", func(c *coc.Client) error { return discard(c.GetPlayer("2PP")) }},
		{"battle log", "/players/%232PP/battlelog", "", func(c *coc.Client) error { return discard(c.GetPlayerBattleLog("2PP")) }},
		{"verify", "/players/%232PP/verifytoken", "", func(c *coc.Client) error { return discard(c.VerifyPlayerToken("2PP", "player-token")) }},
		{"league history", "/players/%232PP/leaguehistory", "", func(c *coc.Client) error { return discard(c.GetPlayerLeagueHistory("2PP")) }},
		{"league", "/leagues/1", "", func(c *coc.Client) error { return discard(c.GetLeague(1)) }},
		{"leagues", "/leagues", "limit=10", func(c *coc.Client) error { return discard(c.GetLeagues(coc.Limit(10))) }},
		{"tier", "/leaguetiers/1", "", func(c *coc.Client) error { return discard(c.GetLeagueTier(1)) }},
		{"tiers", "/leaguetiers", "limit=10", func(c *coc.Client) error { return discard(c.GetLeagueTiers(coc.Limit(10))) }},
		{"capital league", "/capitalleagues/1", "", func(c *coc.Client) error { return discard(c.GetCapitalLeague(1)) }},
		{"capital leagues", "/capitalleagues", "limit=10", func(c *coc.Client) error { return discard(c.GetCapitalLeagues(coc.Limit(10))) }},
		{"builder league", "/builderbaseleagues/1", "", func(c *coc.Client) error { return discard(c.GetBuilderBaseLeague(1)) }},
		{"builder leagues", "/builderbaseleagues", "limit=10", func(c *coc.Client) error { return discard(c.GetBuilderBaseLeagues(coc.Limit(10))) }},
		{"war league", "/warleagues/1", "", func(c *coc.Client) error { return discard(c.GetWarLeague(1)) }},
		{"war leagues", "/warleagues", "limit=10", func(c *coc.Client) error { return discard(c.GetWarLeagues(coc.Limit(10))) }},
		{"seasons", "/leagues/1/seasons", "limit=10", func(c *coc.Client) error { return discard(c.GetLeagueSeasons(1, coc.Limit(10))) }},
		{"season rankings", "/leagues/1/seasons/2026-10", "limit=10", func(c *coc.Client) error { return discard(c.GetLeagueSeasonRankings(1, "2026-10", coc.Limit(10))) }},
		{"league group", "/leaguegroup/%232PP/1", "playerTag=%232PP", func(c *coc.Client) error { return discard(c.GetLeagueGroup("2PP", 1, "2pp")) }},
		{"locations", "/locations", "limit=10", func(c *coc.Client) error { return discard(c.GetLocations(coc.Limit(10))) }},
		{"location", "/locations/1", "", func(c *coc.Client) error { return discard(c.GetLocation(1)) }},
		{"clan rankings", "/locations/global/rankings/clans", "limit=10", func(c *coc.Client) error { return discard(c.GetClanRankings(coc.Global, coc.Limit(10))) }},
		{"player rankings", "/locations/32000006/rankings/players", "limit=10", func(c *coc.Client) error {
			return discard(c.GetPlayerRankings(coc.LocationID(32000006), coc.Limit(10)))
		}},
		{"builder clan rankings", "/locations/global/rankings/clans-builder-base", "limit=10", func(c *coc.Client) error { return discard(c.GetClanBuilderBaseRankings(coc.Global, coc.Limit(10))) }},
		{"builder player rankings", "/locations/global/rankings/players-builder-base", "limit=10", func(c *coc.Client) error { return discard(c.GetPlayerBuilderBaseRankings(coc.Global, coc.Limit(10))) }},
		{"capital rankings", "/locations/global/rankings/capitals", "limit=10", func(c *coc.Client) error { return discard(c.GetClanCapitalRankings(coc.Global, coc.Limit(10))) }},
		{"player labels", "/labels/players", "limit=10", func(c *coc.Client) error { return discard(c.GetPlayerLabels(coc.Limit(10))) }},
		{"clan labels", "/labels/clans", "limit=10", func(c *coc.Client) error { return discard(c.GetClanLabels(coc.Limit(10))) }},
		{"gold pass", "/goldpass/seasons/current", "", func(c *coc.Client) error { return discard(c.GetCurrentGoldPassSeason()) }},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			client := testClient(t, func(r *http.Request) (*http.Response, error) {
				calls++
				if r.URL.EscapedPath() != "/v1"+tc.path {
					t.Errorf("path = %s, want /v1%s", r.URL.EscapedPath(), tc.path)
				}
				expected, _ := url.ParseQuery(tc.query)
				if !reflect.DeepEqual(r.URL.Query(), expected) {
					t.Errorf("query = %v, want %v", r.URL.Query(), expected)
				}
				if r.Header.Get("Authorization") != "Bearer test-token" || r.Header.Get("Accept") != "application/json" {
					t.Errorf("missing request headers: %v", r.Header)
				}
				method := http.MethodGet
				if tc.name == "verify" {
					method = http.MethodPost
					var body struct {
						Token string `json:"token"`
					}
					if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Token != "player-token" {
						t.Errorf("verify body = %+v, error = %v", body, err)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Error("missing content type")
					}
				}
				if r.Method != method {
					t.Errorf("method = %s, want %s", r.Method, method)
				}
				return response(http.StatusOK, `{}`), nil
			})
			if err := tc.call(client); err != nil {
				t.Fatal(err)
			}
			if calls != 1 {
				t.Errorf("requests = %d, want 1", calls)
			}
		})
	}
}

func TestSearchFilterAndCursorEncoding(t *testing.T) {
	client := testClient(t, func(r *http.Request) (*http.Response, error) {
		want := url.Values{"name": {"A & B"}, "warFrequency": {string(coc.ALWAYS)}, "locationId": {"1"}, "minMembers": {"20"}, "maxMembers": {"50"}, "minClanPoints": {"1000"}, "minClanLevel": {"5"}, "labelIds": {"1,2"}, "limit": {"10"}, "after": {"a+/=?&"}}
		if !reflect.DeepEqual(r.URL.Query(), want) {
			t.Errorf("query = %v, want %v", r.URL.Query(), want)
		}
		return response(200, `{"items":[{"tag":"#2PP","name":"Example"}],"paging":{"cursors":{"after":"next"}}}`), nil
	})
	page, err := client.SearchClans(coc.ClanFilter{Name: "A & B", WarFrequency: coc.ALWAYS, LocationID: 1, MinMembers: 20, MaxMembers: 50, MinClanPoints: 1000, MinClanLevel: 5, LabelIDs: []int{1, 2}}, coc.Limit(10), coc.After("a+/=?&"))
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Items) != 1 || page.Items[0].Name != "Example" || page.Paging.Cursors.After != "next" {
		t.Fatalf("decoded page = %+v", page)
	}
}

func TestInvalidOptionsMakeNoRequests(t *testing.T) {
	client := testClient(t, func(*http.Request) (*http.Response, error) {
		t.Error("unexpected request")
		return response(200, `{}`), nil
	})
	for _, opts := range [][]coc.ListOption{
		{coc.Limit(0)}, {coc.Limit(-1)}, {coc.After("a"), coc.Before("b")}, {coc.WithContext(nil)}, {nil},
	} {
		page, err := client.GetLocations(opts...)
		if page != nil || err == nil {
			t.Fatalf("invalid options: page=%v, err=%v", page, err)
		}
	}
}

func TestClientConfiguration(t *testing.T) {
	t.Setenv("COC_API_KEY", "environment-token-must-not-be-used")
	if c, err := coc.NewClient(""); c != nil || err == nil {
		t.Fatal("empty explicit token must fail even with environment set")
	}
	for _, opt := range []coc.ClientOption{nil, coc.WithHTTPClient(nil), coc.WithBaseURL("/v1"), coc.WithBaseURL("https://example.invalid/v1?x=y"), coc.WithBaseURL("ftp://example.invalid"), coc.WithBaseURL("https://user:pass@example.invalid")} {
		if c, err := coc.NewClient("token", opt); c != nil || err == nil {
			t.Fatal("invalid client option accepted")
		}
	}
	if _, err := coc.NewClient("token"); err != nil {
		t.Fatal(err)
	}
}

func TestErrorResponses(t *testing.T) {
	for _, body := range []string{`{"reason":"accessDenied","message":"Private war log"}`, "bad gateway", ""} {
		t.Run(body, func(t *testing.T) {
			client := testClient(t, func(*http.Request) (*http.Response, error) { return response(http.StatusForbidden, body), nil })
			result, err := client.GetClanWarLog("#2PP")
			var apiErr *coc.APIError
			if result != nil || !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusForbidden {
				t.Fatalf("result=%v, err=%v", result, err)
			}
			if strings.HasPrefix(body, "{") && (apiErr.Reason != "accessDenied" || apiErr.Message != "Private war log") {
				t.Fatalf("error details = %+v", apiErr)
			}
		})
	}
	client := testClient(t, func(*http.Request) (*http.Response, error) { return response(200, `invalid`), nil })
	if player, err := client.GetPlayer("#2PP"); player != nil || err == nil {
		t.Fatal("malformed JSON must return nil and an error")
	}
	sentinel := errors.New("transport failed")
	client = testClient(t, func(*http.Request) (*http.Response, error) { return nil, sentinel })
	if _, err := client.GetPlayer("#2PP"); !errors.Is(err, sentinel) {
		t.Fatalf("transport error not preserved: %v", err)
	}
}

func TestRequestContext(t *testing.T) {
	key := struct{}{}
	ctx, cancel := context.WithCancel(context.WithValue(context.Background(), key, "request"))
	started := make(chan struct{})
	client := testClient(t, func(r *http.Request) (*http.Response, error) {
		if r.Context().Value(key) != "request" {
			t.Error("context value not propagated")
		}
		close(started)
		<-r.Context().Done()
		return nil, r.Context().Err()
	})
	done := make(chan error, 1)
	go func() { _, err := client.GetPlayer("#2PP", coc.WithContext(ctx)); done <- err }()
	<-started
	cancel()
	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("request did not stop after cancellation")
	}
	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(-time.Second))
	defer cancel()
	client = testClient(t, func(r *http.Request) (*http.Response, error) { return nil, r.Context().Err() })
	if _, err := client.GetLocations(coc.WithContext(ctx)); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("deadline error = %v", err)
	}
}

func TestPathSegmentsAreEscaped(t *testing.T) {
	tests := []struct {
		path string
		call func(*coc.Client) error
	}{
		{"/v1/leagues/1/seasons/2026%2F10%3Fx=y", func(c *coc.Client) error { return discard(c.GetLeagueSeasonRankings(1, "2026/10?x=y")) }},
		{"/v1/locations/a%2Fb%3Fx=y/rankings/players", func(c *coc.Client) error { return discard(c.GetPlayerRankings(coc.RankingLocation("a/b?x=y"))) }},
	}
	for _, tc := range tests {
		client := testClient(t, func(r *http.Request) (*http.Response, error) {
			if r.URL.EscapedPath() != tc.path || r.URL.RawQuery != "" {
				t.Errorf("URL = %s", r.URL)
			}
			return response(200, `{}`), nil
		})
		if err := tc.call(client); err != nil {
			t.Fatal(err)
		}
	}
}
