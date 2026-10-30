package coc

import (
	"fmt"
	"iter"
)

func paginate[T any](fetch func(...ListOption) (*List[T], error), opts []ListOption) iter.Seq2[T, error] {
	// Copy the option slice so later caller changes do not affect iteration.
	opts = append([]ListOption(nil), opts...)
	return func(yield func(T, error) bool) {
		var zero T
		options, err := resolveOptions(opts)
		if err != nil {
			yield(zero, err)
			return
		}
		direction := "after"
		if options.query.Get("before") != "" {
			direction = "before"
		}
		seen := map[string]bool{}
		if cursor := options.query.Get(direction); cursor != "" {
			seen[cursor] = true
		}
		pageOpts := append([]ListOption(nil), opts...)
		for {
			page, err := fetch(pageOpts...)
			if err != nil {
				yield(zero, err)
				return
			}
			for _, item := range page.Items {
				if err := options.ctx.Err(); err != nil {
					yield(zero, err)
					return
				}
				if !yield(item, nil) {
					return
				}
			}
			next := page.Paging.Cursors.After
			if direction == "before" {
				next = page.Paging.Cursors.Before
			}
			if next == "" {
				return
			}
			if seen[next] {
				yield(zero, fmt.Errorf("coc: API repeated a paging cursor"))
				return
			}
			seen[next] = true
			pageOpts = append(pageOpts[:len(opts)], cursorOption(direction, next))
		}
	}
}
