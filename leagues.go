package coc

import (
	"fmt"
	"net/url"
)

// GetLeagueTier retrieves information about a league tier.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetLeagueTier(leagueTierID int, opts ...RequestOption) (*LeagueTier, error) {
	path := fmt.Sprintf("/leaguetiers/%d", leagueTierID)
	return get[LeagueTier](c, path, opts)
}

// GetCapitalLeagues lists capital leagues.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetCapitalLeagues(opts ...ListOption) (*List[CapitalLeague], error) {
	path := "/capitalleagues"
	return get[List[CapitalLeague]](c, path, opts)
}

// GetLeagueTiers lists league tiers.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetLeagueTiers(opts ...ListOption) (*List[LeagueTier], error) {
	path := "/leaguetiers"
	return get[List[LeagueTier]](c, path, opts)
}

// GetLeagues lists leagues.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetLeagues(opts ...ListOption) (*List[League], error) {
	path := "/leagues"
	return get[List[League]](c, path, opts)
}

// GetLeagueSeasonRankings retrieves league season rankings (Legend League only).
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetLeagueSeasonRankings(leagueID int, seasonID string, opts ...ListOption) (*List[PlayerRanking], error) {
	path := fmt.Sprintf("/leagues/%d/seasons/%s", leagueID, url.PathEscape(seasonID))
	return get[List[PlayerRanking]](c, path, opts)
}

// GetCapitalLeague retrieves information about a capital league.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetCapitalLeague(leagueID int, opts ...RequestOption) (*CapitalLeague, error) {
	path := fmt.Sprintf("/capitalleagues/%d", leagueID)
	return get[CapitalLeague](c, path, opts)
}

// GetBuilderBaseLeague retrieves information about a builder base league.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetBuilderBaseLeague(leagueID int, opts ...RequestOption) (*BuilderBaseLeague, error) {
	path := fmt.Sprintf("/builderbaseleagues/%d", leagueID)
	return get[BuilderBaseLeague](c, path, opts)
}

// GetBuilderBaseLeagues lists builder base leagues.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetBuilderBaseLeagues(opts ...ListOption) (*List[BuilderBaseLeague], error) {
	path := "/builderbaseleagues"
	return get[List[BuilderBaseLeague]](c, path, opts)
}

// GetLeague retrieves information about a league.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetLeague(leagueID int, opts ...RequestOption) (*League, error) {
	path := fmt.Sprintf("/leagues/%d", leagueID)
	return get[League](c, path, opts)
}

// GetLeagueGroup retrieves a ranked league group.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetLeagueGroup(leagueGroupTag string, leagueSeasonID int, playerTag string, opts ...RequestOption) (*LeagueGroup, error) {
	q := url.Values{}
	q.Set("playerTag", normalizeTag(playerTag))
	path := withQuery(fmt.Sprintf("/leaguegroup/%s/%d", FormatTag(leagueGroupTag), leagueSeasonID), q)
	return get[LeagueGroup](c, path, opts)
}

// GetLeagueSeasons lists league seasons (Legend League only).
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetLeagueSeasons(leagueID int, opts ...ListOption) (*List[LeagueSeason], error) {
	path := fmt.Sprintf("/leagues/%d/seasons", leagueID)
	return get[List[LeagueSeason]](c, path, opts)
}

// GetWarLeague retrieves information about a war league.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetWarLeague(leagueID int, opts ...RequestOption) (*WarLeague, error) {
	path := fmt.Sprintf("/warleagues/%d", leagueID)
	return get[WarLeague](c, path, opts)
}

// GetWarLeagues lists war leagues.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetWarLeagues(opts ...ListOption) (*List[WarLeague], error) {
	path := "/warleagues"
	return get[List[WarLeague]](c, path, opts)
}
