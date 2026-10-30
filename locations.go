package coc

import (
	"context"
	"fmt"
)

// GetClanRankings retrieves clan rankings for a location.
func (c *Client) GetClanRankings(ctx context.Context, locationId string, opts *PagingOptions) (*List[ClanRanking], error) {
	path := withQuery(fmt.Sprintf("/locations/%s/rankings/clans", locationId), opts.query())
	var res List[ClanRanking]
	err := c.get(ctx, path, &res)
	return &res, err
}

// GetPlayerRankings retrieves player rankings for a location.
func (c *Client) GetPlayerRankings(ctx context.Context, locationId string, opts *PagingOptions) (*List[PlayerRanking], error) {
	path := withQuery(fmt.Sprintf("/locations/%s/rankings/players", locationId), opts.query())
	var res List[PlayerRanking]
	err := c.get(ctx, path, &res)
	return &res, err
}

// GetPlayerBuilderBaseRankings retrieves player builder base rankings for a location.
func (c *Client) GetPlayerBuilderBaseRankings(ctx context.Context, locationId string, opts *PagingOptions) (*List[PlayerBuilderBaseRanking], error) {
	path := withQuery(fmt.Sprintf("/locations/%s/rankings/players-builder-base", locationId), opts.query())
	var res List[PlayerBuilderBaseRanking]
	err := c.get(ctx, path, &res)
	return &res, err
}

// GetClanBuilderBaseRankings retrieves clan builder base rankings for a location.
func (c *Client) GetClanBuilderBaseRankings(ctx context.Context, locationId string, opts *PagingOptions) (*List[ClanBuilderBaseRanking], error) {
	path := withQuery(fmt.Sprintf("/locations/%s/rankings/clans-builder-base", locationId), opts.query())
	var res List[ClanBuilderBaseRanking]
	err := c.get(ctx, path, &res)
	return &res, err
}

// GetLocations lists locations.
func (c *Client) GetLocations(ctx context.Context, opts *PagingOptions) (*List[Location], error) {
	path := withQuery("/locations", opts.query())
	var res List[Location]
	err := c.get(ctx, path, &res)
	return &res, err
}

// GetClanCapitalRankings retrieves clan capital rankings for a location.
func (c *Client) GetClanCapitalRankings(ctx context.Context, locationId string, opts *PagingOptions) (*List[ClanCapitalRanking], error) {
	path := withQuery(fmt.Sprintf("/locations/%s/rankings/capitals", locationId), opts.query())
	var res List[ClanCapitalRanking]
	err := c.get(ctx, path, &res)
	return &res, err
}

// GetLocation retrieves information about a location.
func (c *Client) GetLocation(ctx context.Context, locationId int) (*Location, error) {
	path := fmt.Sprintf("/locations/%d", locationId)
	var location Location
	err := c.get(ctx, path, &location)
	return &location, err
}
