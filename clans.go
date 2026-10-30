package coc

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// ClanFilter selects clans to search for. Zero-valued fields are omitted.
// Supply at least one criterion; Name must contain at least three characters.
// Paging is configured separately with Limit, After, and Before.
type ClanFilter struct {
	// Name is the clan name to search for (at least three characters).
	Name string
	// WarFrequency restricts how often the clan participates in wars.
	WarFrequency WarFrequency
	// LocationID restricts results to a numeric location ID.
	LocationID int
	// MinMembers is the minimum number of clan members.
	MinMembers int
	// MaxMembers is the maximum number of clan members.
	MaxMembers int
	// MinClanPoints is the minimum number of clan trophy points.
	MinClanPoints int
	// MinClanLevel is the minimum clan level.
	MinClanLevel int
	// LabelIDs restricts results using label IDs from GetClanLabels.
	LabelIDs []int
}

func (f ClanFilter) query() url.Values {
	q := url.Values{}
	if f.Name != "" {
		q.Set("name", f.Name)
	}
	if f.WarFrequency != "" {
		q.Set("warFrequency", string(f.WarFrequency))
	}
	for key, value := range map[string]int{
		"locationId": f.LocationID, "minMembers": f.MinMembers,
		"maxMembers": f.MaxMembers, "minClanPoints": f.MinClanPoints,
		"minClanLevel": f.MinClanLevel,
	} {
		if value != 0 {
			q.Set(key, strconv.Itoa(value))
		}
	}
	if len(f.LabelIDs) > 0 {
		ids := make([]string, len(f.LabelIDs))
		for i, id := range f.LabelIDs {
			ids[i] = strconv.Itoa(id)
		}
		q.Set("labelIds", strings.Join(ids, ","))
	}
	return q
}

// GetClanCurrentWarLeagueGroup retrieves information about a clan's current clan war league group.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetClanCurrentWarLeagueGroup(clanTag string, opts ...RequestOption) (*ClanWarLeagueGroup, error) {
	path := fmt.Sprintf("/clans/%s/currentwar/leaguegroup", FormatTag(clanTag))
	return get[ClanWarLeagueGroup](c, path, opts)
}

// GetClanWarLeagueWar retrieves information about an individual clan war league war.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetClanWarLeagueWar(warTag string, opts ...RequestOption) (*ClanWarLeagueWar, error) {
	path := fmt.Sprintf("/clanwarleagues/wars/%s", FormatTag(warTag))
	return get[ClanWarLeagueWar](c, path, opts)
}

// GetClanWarLog retrieves a clan's clan war log.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetClanWarLog(clanTag string, opts ...ListOption) (*List[ClanWarLogEntry], error) {
	path := fmt.Sprintf("/clans/%s/warlog", FormatTag(clanTag))
	return get[List[ClanWarLogEntry]](c, path, opts)
}

// SearchClans searches for clans by name and/or filter criteria.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) SearchClans(filter ClanFilter, opts ...ListOption) (*List[Clan], error) {
	path := withQuery("/clans", filter.query())
	return get[List[Clan]](c, path, opts)
}

// GetClanCurrentWar retrieves information about a clan's current clan war.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetClanCurrentWar(clanTag string, opts ...RequestOption) (*ClanWar, error) {
	path := fmt.Sprintf("/clans/%s/currentwar", FormatTag(clanTag))
	return get[ClanWar](c, path, opts)
}

// GetClan retrieves information about a clan.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetClan(tag string, opts ...RequestOption) (*Clan, error) {
	path := fmt.Sprintf("/clans/%s", FormatTag(tag))
	return get[Clan](c, path, opts)
}

// GetClanMembers lists clan members.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetClanMembers(clanTag string, opts ...ListOption) (*List[ClanMember], error) {
	path := fmt.Sprintf("/clans/%s/members", FormatTag(clanTag))
	return get[List[ClanMember]](c, path, opts)
}

// GetClanCapitalRaidSeasons retrieves a clan's capital raid seasons.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetClanCapitalRaidSeasons(clanTag string, opts ...ListOption) (*List[ClanCapitalRaidSeason], error) {
	path := fmt.Sprintf("/clans/%s/capitalraidseasons", FormatTag(clanTag))
	return get[List[ClanCapitalRaidSeason]](c, path, opts)
}
