package coc

import "context"

// GetPlayerLabels lists player labels.
func (c *Client) GetPlayerLabels(ctx context.Context, opts *PagingOptions) (*List[Label], error) {
	path := withQuery("/labels/players", opts.query())
	var res List[Label]
	err := c.get(ctx, path, &res)
	return &res, err
}

// GetClanLabels lists clan labels.
func (c *Client) GetClanLabels(ctx context.Context, opts *PagingOptions) (*List[Label], error) {
	path := withQuery("/labels/clans", opts.query())
	var res List[Label]
	err := c.get(ctx, path, &res)
	return &res, err
}
