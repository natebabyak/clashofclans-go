package coc

import (
	"context"
	"fmt"
)

// GetPlayer retrieves information about a player.
func (c *Client) GetPlayer(ctx context.Context, playerTag string) (*Player, error) {
	path := fmt.Sprintf("/players/%s", FormatTag(playerTag))
	var player Player
	err := c.get(ctx, path, &player)
	return &player, err
}

// GetPlayerBattleLog retrieves a player's recent battle log.
func (c *Client) GetPlayerBattleLog(ctx context.Context, playerTag string) (*List[BattleLogEntry], error) {
	path := fmt.Sprintf("/players/%s/battlelog", FormatTag(playerTag))
	var res List[BattleLogEntry]
	err := c.get(ctx, path, &res)
	return &res, err
}

// VerifyPlayerToken verifies a player API token from the game settings.
func (c *Client) VerifyPlayerToken(ctx context.Context, playerTag, token string) (*VerifyTokenResponse, error) {
	path := fmt.Sprintf("/players/%s/verifytoken", FormatTag(playerTag))
	body := struct {
		Token string `json:"token"`
	}{Token: token}
	var res VerifyTokenResponse
	err := c.post(ctx, path, body, &res)
	return &res, err
}

// GetPlayerLeagueHistory retrieves a player's ranked league history.
func (c *Client) GetPlayerLeagueHistory(ctx context.Context, playerTag string) (*List[LeagueSeasonResult], error) {
	path := fmt.Sprintf("/players/%s/leaguehistory", FormatTag(playerTag))
	var res List[LeagueSeasonResult]
	err := c.get(ctx, path, &res)
	return &res, err
}
