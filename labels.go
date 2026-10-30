package coc

// GetPlayerLabels lists player labels.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetPlayerLabels(opts ...ListOption) (*List[Label], error) {
	path := "/labels/players"
	return get[List[Label]](c, path, opts)
}

// GetClanLabels lists clan labels.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetClanLabels(opts ...ListOption) (*List[Label], error) {
	path := "/labels/clans"
	return get[List[Label]](c, path, opts)
}
