package coc

// GetCurrentGoldPassSeason retrieves information about the current gold pass season.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetCurrentGoldPassSeason(opts ...RequestOption) (*GoldPassSeason, error) {
	path := "/goldpass/seasons/current"
	return get[GoldPassSeason](c, path, opts)
}
