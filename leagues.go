package coc

import (
	"context"
	"fmt"
	"net/url"
)

// GetLeagueTier retrieves information about a league tier.
func (c *Client) GetLeagueTier(ctx context.Context, leagueTierId int) (*LeagueTier, error) {
	path := fmt.Sprintf("/leaguetiers/%d", leagueTierId)
	var tier LeagueTier
	err := c.get(ctx, path, &tier)
	return &tier, err
}

// GetCapitalLeagues lists capital leagues.
func (c *Client) GetCapitalLeagues(ctx context.Context, opts *PagingOptions) (*List[CapitalLeague], error) {
	path := withQuery("/capitalleagues", opts.query())
	var res List[CapitalLeague]
	err := c.get(ctx, path, &res)
	return &res, err
}

// GetLeagueTiers lists league tiers.
func (c *Client) GetLeagueTiers(ctx context.Context, opts *PagingOptions) (*List[LeagueTier], error) {
	path := withQuery("/leaguetiers", opts.query())
	var res List[LeagueTier]
	err := c.get(ctx, path, &res)
	return &res, err
}

// GetLeagues lists leagues.
func (c *Client) GetLeagues(ctx context.Context, opts *PagingOptions) (*List[League], error) {
	path := withQuery("/leagues", opts.query())
	var res List[League]
	err := c.get(ctx, path, &res)
	return &res, err
}

// GetLeagueSeasonRankings retrieves league season rankings (Legend League only).
func (c *Client) GetLeagueSeasonRankings(ctx context.Context, leagueId int, seasonId string, opts *PagingOptions) (*List[PlayerRanking], error) {
	path := withQuery(fmt.Sprintf("/leagues/%d/seasons/%s", leagueId, seasonId), opts.query())
	var res List[PlayerRanking]
	err := c.get(ctx, path, &res)
	return &res, err
}

// GetCapitalLeague retrieves information about a capital league.
func (c *Client) GetCapitalLeague(ctx context.Context, leagueId int) (*CapitalLeague, error) {
	path := fmt.Sprintf("/capitalleagues/%d", leagueId)
	var league CapitalLeague
	err := c.get(ctx, path, &league)
	return &league, err
}

// GetBuilderBaseLeague retrieves information about a builder base league.
func (c *Client) GetBuilderBaseLeague(ctx context.Context, leagueId int) (*BuilderBaseLeague, error) {
	path := fmt.Sprintf("/builderbaseleagues/%d", leagueId)
	var league BuilderBaseLeague
	err := c.get(ctx, path, &league)
	return &league, err
}

// GetBuilderBaseLeagues lists builder base leagues.
func (c *Client) GetBuilderBaseLeagues(ctx context.Context, opts *PagingOptions) (*List[BuilderBaseLeague], error) {
	path := withQuery("/builderbaseleagues", opts.query())
	var res List[BuilderBaseLeague]
	err := c.get(ctx, path, &res)
	return &res, err
}

// GetLeague retrieves information about a league.
func (c *Client) GetLeague(ctx context.Context, leagueId int) (*League, error) {
	path := fmt.Sprintf("/leagues/%d", leagueId)
	var league League
	err := c.get(ctx, path, &league)
	return &league, err
}

// GetLeagueGroup retrieves a ranked league group.
func (c *Client) GetLeagueGroup(ctx context.Context, leagueGroupTag string, leagueSeasonId int, playerTag string) (*LeagueGroup, error) {
	q := url.Values{}
	q.Set("playerTag", playerTag)
	path := withQuery(fmt.Sprintf("/leaguegroup/%s/%d", FormatTag(leagueGroupTag), leagueSeasonId), q)
	var group LeagueGroup
	err := c.get(ctx, path, &group)
	return &group, err
}

// GetLeagueSeasons lists league seasons (Legend League only).
func (c *Client) GetLeagueSeasons(ctx context.Context, leagueId int, opts *PagingOptions) (*List[LeagueSeason], error) {
	path := withQuery(fmt.Sprintf("/leagues/%d/seasons", leagueId), opts.query())
	var res List[LeagueSeason]
	err := c.get(ctx, path, &res)
	return &res, err
}

// GetWarLeague retrieves information about a war league.
func (c *Client) GetWarLeague(ctx context.Context, leagueId int) (*WarLeague, error) {
	path := fmt.Sprintf("/warleagues/%d", leagueId)
	var league WarLeague
	err := c.get(ctx, path, &league)
	return &league, err
}

// GetWarLeagues lists war leagues.
func (c *Client) GetWarLeagues(ctx context.Context, opts *PagingOptions) (*List[WarLeague], error) {
	path := withQuery("/warleagues", opts.query())
	var res List[WarLeague]
	err := c.get(ctx, path, &res)
	return &res, err
}
