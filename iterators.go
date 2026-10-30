package coc

import "iter"

// ClanWarLog lazily iterates items from GetClanWarLog, fetching pages as needed.
// Limit controls page size. Use WithContext for cancellation; breaking the loop
// stops further requests. For one explicit page, use GetClanWarLog.
func (c *Client) ClanWarLog(clanTag string, opts ...ListOption) iter.Seq2[ClanWarLogEntry, error] {
	return paginate(func(pageOpts ...ListOption) (*List[ClanWarLogEntry], error) {
		return c.GetClanWarLog(clanTag, pageOpts...)
	}, opts)
}

// Clans lazily iterates items from SearchClans, fetching pages as needed.
// Limit controls page size. Use WithContext for cancellation; breaking the loop
// stops further requests. For one explicit page, use SearchClans.
func (c *Client) Clans(filter ClanFilter, opts ...ListOption) iter.Seq2[Clan, error] {
	filter.LabelIDs = append([]int(nil), filter.LabelIDs...)
	return paginate(func(pageOpts ...ListOption) (*List[Clan], error) { return c.SearchClans(filter, pageOpts...) }, opts)
}

// ClanMembers lazily iterates items from GetClanMembers, fetching pages as needed.
// Limit controls page size. Use WithContext for cancellation; breaking the loop
// stops further requests. For one explicit page, use GetClanMembers.
func (c *Client) ClanMembers(clanTag string, opts ...ListOption) iter.Seq2[ClanMember, error] {
	return paginate(func(pageOpts ...ListOption) (*List[ClanMember], error) { return c.GetClanMembers(clanTag, pageOpts...) }, opts)
}

// ClanCapitalRaidSeasons lazily iterates items from GetClanCapitalRaidSeasons, fetching pages as needed.
// Limit controls page size. Use WithContext for cancellation; breaking the loop
// stops further requests. For one explicit page, use GetClanCapitalRaidSeasons.
func (c *Client) ClanCapitalRaidSeasons(clanTag string, opts ...ListOption) iter.Seq2[ClanCapitalRaidSeason, error] {
	return paginate(func(pageOpts ...ListOption) (*List[ClanCapitalRaidSeason], error) {
		return c.GetClanCapitalRaidSeasons(clanTag, pageOpts...)
	}, opts)
}

// CapitalLeagues lazily iterates items from GetCapitalLeagues, fetching pages as needed.
// Limit controls page size. Use WithContext for cancellation; breaking the loop
// stops further requests. For one explicit page, use GetCapitalLeagues.
func (c *Client) CapitalLeagues(opts ...ListOption) iter.Seq2[CapitalLeague, error] {
	return paginate(c.GetCapitalLeagues, opts)
}

// LeagueTiers lazily iterates items from GetLeagueTiers, fetching pages as needed.
// Limit controls page size. Use WithContext for cancellation; breaking the loop
// stops further requests. For one explicit page, use GetLeagueTiers.
func (c *Client) LeagueTiers(opts ...ListOption) iter.Seq2[LeagueTier, error] {
	return paginate(c.GetLeagueTiers, opts)
}

// Leagues lazily iterates items from GetLeagues, fetching pages as needed.
// Limit controls page size. Use WithContext for cancellation; breaking the loop
// stops further requests. For one explicit page, use GetLeagues.
func (c *Client) Leagues(opts ...ListOption) iter.Seq2[League, error] {
	return paginate(c.GetLeagues, opts)
}

// LeagueSeasonRankings lazily iterates items from GetLeagueSeasonRankings, fetching pages as needed.
// Limit controls page size. Use WithContext for cancellation; breaking the loop
// stops further requests. For one explicit page, use GetLeagueSeasonRankings.
func (c *Client) LeagueSeasonRankings(leagueID int, seasonID string, opts ...ListOption) iter.Seq2[PlayerRanking, error] {
	return paginate(func(pageOpts ...ListOption) (*List[PlayerRanking], error) {
		return c.GetLeagueSeasonRankings(leagueID, seasonID, pageOpts...)
	}, opts)
}

// BuilderBaseLeagues lazily iterates items from GetBuilderBaseLeagues, fetching pages as needed.
// Limit controls page size. Use WithContext for cancellation; breaking the loop
// stops further requests. For one explicit page, use GetBuilderBaseLeagues.
func (c *Client) BuilderBaseLeagues(opts ...ListOption) iter.Seq2[BuilderBaseLeague, error] {
	return paginate(c.GetBuilderBaseLeagues, opts)
}

// LeagueSeasons lazily iterates items from GetLeagueSeasons, fetching pages as needed.
// Limit controls page size. Use WithContext for cancellation; breaking the loop
// stops further requests. For one explicit page, use GetLeagueSeasons.
func (c *Client) LeagueSeasons(leagueID int, opts ...ListOption) iter.Seq2[LeagueSeason, error] {
	return paginate(func(pageOpts ...ListOption) (*List[LeagueSeason], error) {
		return c.GetLeagueSeasons(leagueID, pageOpts...)
	}, opts)
}

// WarLeagues lazily iterates items from GetWarLeagues, fetching pages as needed.
// Limit controls page size. Use WithContext for cancellation; breaking the loop
// stops further requests. For one explicit page, use GetWarLeagues.
func (c *Client) WarLeagues(opts ...ListOption) iter.Seq2[WarLeague, error] {
	return paginate(c.GetWarLeagues, opts)
}

// ClanRankings lazily iterates items from GetClanRankings, fetching pages as needed.
// Limit controls page size. Use WithContext for cancellation; breaking the loop
// stops further requests. For one explicit page, use GetClanRankings.
func (c *Client) ClanRankings(locationID RankingLocation, opts ...ListOption) iter.Seq2[ClanRanking, error] {
	return paginate(func(pageOpts ...ListOption) (*List[ClanRanking], error) {
		return c.GetClanRankings(locationID, pageOpts...)
	}, opts)
}

// PlayerRankings lazily iterates items from GetPlayerRankings, fetching pages as needed.
// Limit controls page size. Use WithContext for cancellation; breaking the loop
// stops further requests. For one explicit page, use GetPlayerRankings.
func (c *Client) PlayerRankings(locationID RankingLocation, opts ...ListOption) iter.Seq2[PlayerRanking, error] {
	return paginate(func(pageOpts ...ListOption) (*List[PlayerRanking], error) {
		return c.GetPlayerRankings(locationID, pageOpts...)
	}, opts)
}

// PlayerBuilderBaseRankings lazily iterates items from GetPlayerBuilderBaseRankings, fetching pages as needed.
// Limit controls page size. Use WithContext for cancellation; breaking the loop
// stops further requests. For one explicit page, use GetPlayerBuilderBaseRankings.
func (c *Client) PlayerBuilderBaseRankings(locationID RankingLocation, opts ...ListOption) iter.Seq2[PlayerBuilderBaseRanking, error] {
	return paginate(func(pageOpts ...ListOption) (*List[PlayerBuilderBaseRanking], error) {
		return c.GetPlayerBuilderBaseRankings(locationID, pageOpts...)
	}, opts)
}

// ClanBuilderBaseRankings lazily iterates items from GetClanBuilderBaseRankings, fetching pages as needed.
// Limit controls page size. Use WithContext for cancellation; breaking the loop
// stops further requests. For one explicit page, use GetClanBuilderBaseRankings.
func (c *Client) ClanBuilderBaseRankings(locationID RankingLocation, opts ...ListOption) iter.Seq2[ClanBuilderBaseRanking, error] {
	return paginate(func(pageOpts ...ListOption) (*List[ClanBuilderBaseRanking], error) {
		return c.GetClanBuilderBaseRankings(locationID, pageOpts...)
	}, opts)
}

// Locations lazily iterates items from GetLocations, fetching pages as needed.
// Limit controls page size. Use WithContext for cancellation; breaking the loop
// stops further requests. For one explicit page, use GetLocations.
func (c *Client) Locations(opts ...ListOption) iter.Seq2[Location, error] {
	return paginate(c.GetLocations, opts)
}

// ClanCapitalRankings lazily iterates items from GetClanCapitalRankings, fetching pages as needed.
// Limit controls page size. Use WithContext for cancellation; breaking the loop
// stops further requests. For one explicit page, use GetClanCapitalRankings.
func (c *Client) ClanCapitalRankings(locationID RankingLocation, opts ...ListOption) iter.Seq2[ClanCapitalRanking, error] {
	return paginate(func(pageOpts ...ListOption) (*List[ClanCapitalRanking], error) {
		return c.GetClanCapitalRankings(locationID, pageOpts...)
	}, opts)
}

// PlayerLabels lazily iterates items from GetPlayerLabels, fetching pages as needed.
// Limit controls page size. Use WithContext for cancellation; breaking the loop
// stops further requests. For one explicit page, use GetPlayerLabels.
func (c *Client) PlayerLabels(opts ...ListOption) iter.Seq2[Label, error] {
	return paginate(c.GetPlayerLabels, opts)
}

// ClanLabels lazily iterates items from GetClanLabels, fetching pages as needed.
// Limit controls page size. Use WithContext for cancellation; breaking the loop
// stops further requests. For one explicit page, use GetClanLabels.
func (c *Client) ClanLabels(opts ...ListOption) iter.Seq2[Label, error] {
	return paginate(c.GetClanLabels, opts)
}
