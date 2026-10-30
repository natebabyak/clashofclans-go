package coc

import (
	"net/url"
	"strconv"
	"strings"
)

func FormatTag(tag string) string {
	if !strings.HasPrefix(tag, "#") {
		tag = "#" + tag
	}
	return url.PathEscape(tag)
}

type PagingOptions struct {
	Limit  *int
	After  *string
	Before *string
}

func (o *PagingOptions) query() url.Values {
	q := url.Values{}
	if o == nil {
		return q
	}
	setPaging(q, o.Limit, o.After, o.Before)
	return q
}

func setPaging(q url.Values, limit *int, after, before *string) {
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if after != nil {
		q.Set("after", *after)
	}
	if before != nil {
		q.Set("before", *before)
	}
}

func withQuery(path string, q url.Values) string {
	if len(q) == 0 {
		return path
	}
	return path + "?" + q.Encode()
}
