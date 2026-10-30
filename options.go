package coc

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// ListOption configures a list request or iterator. Use Limit, After, Before,
// or WithContext. Omitting options uses the API's default page size and cursors.
// Options are immutable and may be reused across requests.
type ListOption interface {
	apply(*requestOptions) error
}

// RequestOption configures any route, including list routes. Use WithContext.
// Paging options implement only ListOption, so they cannot be passed to routes
// that do not support paging.
type RequestOption interface {
	ListOption
	requestOption()
}

type requestOptions struct {
	ctx   context.Context
	query url.Values
}

type requestOption func(*requestOptions) error

func (o requestOption) apply(r *requestOptions) error { return o(r) }
func (requestOption) requestOption()                  {}

type listOption func(*requestOptions) error

func (o listOption) apply(r *requestOptions) error { return o(r) }

// WithContext supplies cancellation and deadlines for a request or iterator.
// Without it, requests use context.Background and the client's HTTP timeout.
// Passing a nil context returns an error when the request is made.
func WithContext(ctx context.Context) RequestOption {
	return requestOption(func(r *requestOptions) error {
		if ctx == nil {
			return fmt.Errorf("coc: context must not be nil")
		}
		r.ctx = ctx
		return nil
	})
}

// Limit sets the maximum number of items requested per page. It must be positive.
// On iterators, this controls page size, not the total number of items yielded.
// Omit Limit to use the API's default page size.
func Limit(n int) ListOption {
	return listOption(func(r *requestOptions) error {
		if n <= 0 {
			return fmt.Errorf("coc: limit must be positive")
		}
		r.query.Set("limit", strconv.Itoa(n))
		return nil
	})
}

// After starts a page after an opaque cursor returned by the API.
// An empty cursor is omitted. After and Before cannot both be nonempty.
func After(cursor string) ListOption { return cursorOption("after", cursor) }

// Before starts a page before an opaque cursor returned by the API.
// An empty cursor is omitted. After and Before cannot both be nonempty.
// Iterators follow before cursors when started with Before.
func Before(cursor string) ListOption { return cursorOption("before", cursor) }

func cursorOption(key, cursor string) ListOption {
	return listOption(func(r *requestOptions) error {
		if cursor == "" {
			r.query.Del(key)
		} else {
			r.query.Set(key, cursor)
		}
		return nil
	})
}

func resolveOptions[O ListOption](opts []O) (requestOptions, error) {
	r := requestOptions{ctx: context.Background(), query: url.Values{}}
	for _, opt := range opts {
		if any(opt) == nil {
			return r, fmt.Errorf("coc: option must not be nil")
		}
		if err := opt.apply(&r); err != nil {
			return r, err
		}
	}
	if r.query.Get("after") != "" && r.query.Get("before") != "" {
		return r, fmt.Errorf("coc: after and before cannot be used together")
	}
	return r, nil
}
