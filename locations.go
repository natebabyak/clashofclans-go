package coc

import (
	"fmt"
	"net/url"
)

// GetClanRankings retrieves clan rankings for a location.
// Use Global or LocationID(id) to select the ranking location.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetClanRankings(locationID RankingLocation, opts ...ListOption) (*List[ClanRanking], error) {
	path := fmt.Sprintf("/locations/%s/rankings/clans", url.PathEscape(string(locationID)))
	return get[List[ClanRanking]](c, path, opts)
}

// GetPlayerRankings retrieves player rankings for a location.
// Use Global or LocationID(id) to select the ranking location.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetPlayerRankings(locationID RankingLocation, opts ...ListOption) (*List[PlayerRanking], error) {
	path := fmt.Sprintf("/locations/%s/rankings/players", url.PathEscape(string(locationID)))
	return get[List[PlayerRanking]](c, path, opts)
}

// GetPlayerBuilderBaseRankings retrieves player builder base rankings for a location.
// Use Global or LocationID(id) to select the ranking location.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetPlayerBuilderBaseRankings(locationID RankingLocation, opts ...ListOption) (*List[PlayerBuilderBaseRanking], error) {
	path := fmt.Sprintf("/locations/%s/rankings/players-builder-base", url.PathEscape(string(locationID)))
	return get[List[PlayerBuilderBaseRanking]](c, path, opts)
}

// GetClanBuilderBaseRankings retrieves clan builder base rankings for a location.
// Use Global or LocationID(id) to select the ranking location.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetClanBuilderBaseRankings(locationID RankingLocation, opts ...ListOption) (*List[ClanBuilderBaseRanking], error) {
	path := fmt.Sprintf("/locations/%s/rankings/clans-builder-base", url.PathEscape(string(locationID)))
	return get[List[ClanBuilderBaseRanking]](c, path, opts)
}

// GetLocations lists locations.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetLocations(opts ...ListOption) (*List[Location], error) {
	path := "/locations"
	return get[List[Location]](c, path, opts)
}

// GetClanCapitalRankings retrieves clan capital rankings for a location.
// Use Global or LocationID(id) to select the ranking location.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetClanCapitalRankings(locationID RankingLocation, opts ...ListOption) (*List[ClanCapitalRanking], error) {
	path := fmt.Sprintf("/locations/%s/rankings/capitals", url.PathEscape(string(locationID)))
	return get[List[ClanCapitalRanking]](c, path, opts)
}

// GetLocation retrieves information about a location.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetLocation(locationID int, opts ...RequestOption) (*Location, error) {
	path := fmt.Sprintf("/locations/%d", locationID)
	return get[Location](c, path, opts)
}
