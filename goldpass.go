package coc

import "context"

// GetCurrentGoldPassSeason retrieves information about the current gold pass season.
func (c *Client) GetCurrentGoldPassSeason(ctx context.Context) (*GoldPassSeason, error) {
	var season GoldPassSeason
	err := c.get(ctx, "/goldpass/seasons/current", &season)
	return &season, err
}
