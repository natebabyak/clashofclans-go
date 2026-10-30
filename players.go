package coc

import "fmt"

// GetPlayer retrieves information about a player.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetPlayer(playerTag string, opts ...RequestOption) (*Player, error) {
	path := fmt.Sprintf("/players/%s", FormatTag(playerTag))
	return get[Player](c, path, opts)
}

// GetPlayerBattleLog retrieves a player's recent battle log.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetPlayerBattleLog(playerTag string, opts ...RequestOption) (*List[BattleLogEntry], error) {
	path := fmt.Sprintf("/players/%s/battlelog", FormatTag(playerTag))
	return get[List[BattleLogEntry]](c, path, opts)
}

// VerifyPlayerToken verifies a player API token from the game settings.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) VerifyPlayerToken(playerTag, token string, opts ...RequestOption) (*VerifyTokenResponse, error) {
	path := fmt.Sprintf("/players/%s/verifytoken", FormatTag(playerTag))
	body := struct {
		Token string `json:"token"`
	}{Token: token}
	return post[VerifyTokenResponse](c, path, body, opts)
}

// GetPlayerLeagueHistory retrieves a player's ranked league history.
// Makes one API request. Use WithContext to supply cancellation or a deadline.
func (c *Client) GetPlayerLeagueHistory(playerTag string, opts ...RequestOption) (*List[LeagueSeasonResult], error) {
	path := fmt.Sprintf("/players/%s/leaguehistory", FormatTag(playerTag))
	return get[List[LeagueSeasonResult]](c, path, opts)
}
