package coc

import (
	"net/url"
	"strconv"
	"strings"
)

// RankingLocation identifies a leaderboard region. Use Global for worldwide
// rankings or LocationID for a numeric ID returned by GetLocations.
type RankingLocation string

// Global selects worldwide rankings instead of a numeric location.
const Global RankingLocation = "global"

// LocationID converts an ID from GetLocations to a ranking location argument.
// For example: client.GetPlayerRankings(coc.LocationID(location.Id)).
func LocationID(id int) RankingLocation { return RankingLocation(strconv.Itoa(id)) }

func normalizeTag(tag string) string {
	return "#" + strings.TrimPrefix(strings.ToUpper(strings.TrimSpace(tag)), "#")
}

// FormatTag normalizes a tag and escapes it for use in a URL path.
// Route methods do this automatically; pass raw tags with or without '#'.
func FormatTag(tag string) string { return url.PathEscape(normalizeTag(tag)) }

func withQuery(path string, q url.Values) string {
	if len(q) == 0 {
		return path
	}
	return path + "?" + q.Encode()
}
