# clashofclans-go

A small, strongly typed Go client for the [Clash of Clans API](https://developer.clashofclans.com/#/documentation). Each route has a dedicated method. No external dependencies.

## Getting started

[Create an API key](https://developer.clashofclans.com/) and install the module:

```sh
go get github.com/natebabyak/clashofclans-go
```

```go
package main

import (
    "fmt"
    "log"

    coc "github.com/natebabyak/clashofclans-go"
)

func main() {
    client, err := coc.NewClient("YOUR_API_TOKEN")
    if err != nil {
        log.Fatal(err)
    }

    clan, err := client.GetClan("#2PP")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s has %d members!\n", clan.Name, clan.Members)
}
```

Tokens are always passed explicitly. The library never reads environment variables or `.env` files. Tags can include or omit `#`.

## Lists and paging

Options are optional, typed values. Omitting them uses the API defaults. `Get` and `Search` methods fetch exactly one response page:

```go
page, err := client.GetPlayerRankings(coc.Global, coc.Limit(100))
if err != nil {
    return err
}

// Fetch another page only when the API supplies a cursor.
if cursor := page.Paging.Cursors.After; cursor != "" {
    next, err := client.GetPlayerRankings(coc.Global, coc.Limit(100), coc.After(cursor))
    if err != nil {
        return err
    }
    fmt.Println(next.Items)
}
```

Use `coc.LocationID(id)` for regional rankings. `After` and `Before` accept opaque API cursors and cannot both be nonempty. Paging options are accepted only by routes that support paging.

Iterator methods follow cursors lazily. `Limit` controls the page size, not the total number of results. Breaking the loop stops further requests:

```go
for member, err := range client.ClanMembers("#2PP", coc.Limit(50)) {
    if err != nil {
        return err
    }
    fmt.Println(member.Name)
}
```

Every paged route has an iterator: `GetClanMembers` → `ClanMembers`, `GetPlayerRankings` → `PlayerRankings`, and `SearchClans` → `Clans`. Iterators follow `after` cursors by default, or `before` cursors when started with `Before`. A failing iteration yields an error with a zero-valued item and then stops.

## Clan search

Filters use ordinary values. Zero-valued fields are omitted; label IDs are encoded automatically:

```go
clans, err := client.SearchClans(coc.ClanFilter{
    Name:       "Example",
    MinMembers: 20,
    LabelIDs:   []int{1, 2},
}, coc.Limit(50))
```

## Cancellation and configuration

Requests default to a ten-second HTTP timeout. Supply a context when cancellation or a deadline is needed:

```go
player, err := client.GetPlayer("#ABC", coc.WithContext(ctx))
```

The same option works on list methods and iterators. Reuse the client across goroutines. To customize its transport, timeout, or API root:

```go
client, err := coc.NewClient(token,
    coc.WithHTTPClient(&http.Client{Timeout: 30 * time.Second}),
    coc.WithBaseURL("https://api.clashofclans.com/v1"),
)
```

A custom HTTP client replaces the default, including its timeout.

## API errors

Unsuccessful HTTP responses return `*coc.APIError`. Transport and context errors remain inspectable through `errors.Is` and `errors.As`:

```go
var apiErr *coc.APIError
if errors.As(err, &apiErr) {
    fmt.Println(apiErr.StatusCode, apiErr.Reason, apiErr.Message)
}
```

## Migrating from the previous API

- `NewClient(token)` now returns `(*Client, error)`.
- Remove the leading context argument; pass `coc.WithContext(ctx)` when needed.
- Remove trailing `nil` options. Replace `*PagingOptions` with `coc.Limit(n)`, `coc.After(cursor)`, or `coc.Before(cursor)`.
- Replace `*SearchClansOptions` with `ClanFilter` values and pass paging options separately.
- For rankings, use `coc.Global` or `coc.LocationID(id)`.
- Route methods return a nil response on failure.

## Example CLI

[`examples/equipment-cli`](examples/equipment-cli) fetches a page of top global players, retrieves their profiles, and prints common hero equipment combinations:

```sh
go run ./examples/equipment-cli -token YOUR_API_TOKEN -players 100
```

The example processes profiles sequentially to keep the API usage clear.

## Legal

Usage is subject to Supercell's Terms and Conditions.

This material is unofficial and is not endorsed by Supercell. For more information see Supercell's [Fan Content Policy](https://www.supercell.com/fan-content-policy).

## License

MIT
