package coc

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type SearchClansOptions struct {
	PagingOptions
	Name          *string
	WarFrequency  *WarFrequency
	LocationId    *int
	MinMembers    *int
	MaxMembers    *int
	MinClanPoints *int
	MinClanLevel  *int
	LabelIds      *string
}

func (o *SearchClansOptions) query() url.Values {
	q := url.Values{}
	if o == nil {
		return q
	}
	setPaging(q, o.Limit, o.After, o.Before)
	if o.Name != nil {
		q.Set("name", *o.Name)
	}
	if o.WarFrequency != nil {
		q.Set("warFrequency", string(*o.WarFrequency))
	}
	if o.LocationId != nil {
		q.Set("locationId", strconv.Itoa(*o.LocationId))
	}
	if o.MinMembers != nil {
		q.Set("minMembers", strconv.Itoa(*o.MinMembers))
	}
	if o.MaxMembers != nil {
		q.Set("maxMembers", strconv.Itoa(*o.MaxMembers))
	}
	if o.MinClanPoints != nil {
		q.Set("minClanPoints", strconv.Itoa(*o.MinClanPoints))
	}
	if o.MinClanLevel != nil {
		q.Set("minClanLevel", strconv.Itoa(*o.MinClanLevel))
	}
	if o.LabelIds != nil {
		q.Set("labelIds", *o.LabelIds)
	}
	return q
}

// GetClanCurrentWarLeagueGroup retrieves information about a clan's current clan war league group.
func (c *Client) GetClanCurrentWarLeagueGroup(ctx context.Context, clanTag string) (*ClanWarLeagueGroup, error) {
	path := fmt.Sprintf("/clans/%s/currentwar/leaguegroup", FormatTag(clanTag))
	var group ClanWarLeagueGroup
	err := c.get(ctx, path, &group)
	return &group, err
}

// GetClanWarLeagueWar retrieves information about an individual clan war league war.
func (c *Client) GetClanWarLeagueWar(ctx context.Context, warTag string) (*ClanWarLeagueWar, error) {
	path := fmt.Sprintf("/clanwarleagues/wars/%s", FormatTag(warTag))
	var war ClanWarLeagueWar
	err := c.get(ctx, path, &war)
	return &war, err
}

// GetClanWarLog retrieves a clan's clan war log.
func (c *Client) GetClanWarLog(ctx context.Context, clanTag string, opts *PagingOptions) (*List[ClanWarLogEntry], error) {
	path := withQuery(fmt.Sprintf("/clans/%s/warlog", FormatTag(clanTag)), opts.query())
	var res List[ClanWarLogEntry]
	err := c.get(ctx, path, &res)
	return &res, err
}

// SearchClans searches for clans by name and/or filter criteria.
func (c *Client) SearchClans(ctx context.Context, opts *SearchClansOptions) (*List[Clan], error) {
	path := withQuery("/clans", opts.query())
	var res List[Clan]
	err := c.get(ctx, path, &res)
	return &res, err
}

// GetClanCurrentWar retrieves information about a clan's current clan war.
func (c *Client) GetClanCurrentWar(ctx context.Context, clanTag string) (*ClanWar, error) {
	path := fmt.Sprintf("/clans/%s/currentwar", FormatTag(clanTag))
	var war ClanWar
	err := c.get(ctx, path, &war)
	return &war, err
}

// GetClan retrieves information about a clan.
func (c *Client) GetClan(ctx context.Context, tag string) (*Clan, error) {
	path := fmt.Sprintf("/clans/%s", FormatTag(tag))
	var clan Clan
	err := c.get(ctx, path, &clan)
	return &clan, err
}

// GetClanMembers lists clan members.
func (c *Client) GetClanMembers(ctx context.Context, clanTag string, opts *PagingOptions) (*List[ClanMember], error) {
	path := withQuery(fmt.Sprintf("/clans/%s/members", FormatTag(clanTag)), opts.query())
	var res List[ClanMember]
	err := c.get(ctx, path, &res)
	return &res, err
}

// GetClanCapitalRaidSeasons retrieves a clan's capital raid seasons.
func (c *Client) GetClanCapitalRaidSeasons(ctx context.Context, clanTag string, opts *PagingOptions) (*List[ClanCapitalRaidSeason], error) {
	path := withQuery(fmt.Sprintf("/clans/%s/capitalraidseasons", FormatTag(clanTag)), opts.query())
	var res List[ClanCapitalRaidSeason]
	err := c.get(ctx, path, &res)
	return &res, err
}
