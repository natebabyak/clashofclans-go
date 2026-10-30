// Package coc provides typed, per-route access to the Clash of Clans API.
//
// Construct a client with an explicit token using NewClient. Get and Search
// methods make one request and return typed responses. Context is optional:
// pass WithContext when cancellation or a deadline is needed. List routes
// accept Limit, After, and Before, with API defaults when omitted.
//
// Iterator methods, such as ClanMembers and PlayerRankings, lazily follow
// cursors. Limit sets their page size; breaking a range loop stops fetching.
// All requests use the client's HTTP timeout, which defaults to ten seconds.
// The package never reads environment variables or .env files.
//
// API reference: https://developer.clashofclans.com/#/documentation
package coc
