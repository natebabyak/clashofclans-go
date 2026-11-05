# clashofclans-go

A minimal, strongly typed Go client for the [Clash of Clans API](https://developer.clashofclans.com/#/documentation).

## Getting Started

### Prerequisites

- Go 1.25.6 or later.
- An API key from the [Clash of Clans developer portal](https://developer.clashofclans.com/).

### Installation

Install the package:

```sh
go get github.com/natebabyak/clashofclans-go
```

### Basic Usage

Pass your API key to `coc.NewClient`, then call `GetClan` with a clan tag.

```go
package main

import (
    "fmt"
    "log"

    coc "github.com/natebabyak/clashofclans-go"
)

func main() {
    client, err := coc.NewClient("YOUR_API_KEY")
    if err != nil {
        log.Fatal(err)
    }

    clan, err := client.GetClan("#2PP")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s has %d members.\n", clan.Name, clan.Members)
}
```

See [basic-usage/main.go](examples/basic-usage/main.go) for the standalone example.

### Advanced usage

**Paginate results.** Use `ClanMembers` to fetch pages as you iterate. `coc.Limit(50)` sets the page size. Break the loop when you have enough results.

```go
for member, err := range client.ClanMembers("#2PP", coc.Limit(50)) {
    if err != nil {
        return err
    }
    fmt.Println(member.Name)
}
```

For a single page, call `GetClanMembers`. To fetch the next page, pass `coc.After(page.Paging.Cursors.After)` when the cursor is nonempty.

**Search clans.** Use `SearchClans` with a `coc.ClanFilter`. For example, `coc.ClanFilter{Name: "Example", MinMembers: 20}` searches by name and minimum member count.

**Set a deadline.** Requests have a ten-second HTTP timeout by default. Pass `coc.WithContext(ctx)` to use a context with a deadline or cancellation.

**Handle API errors.** Use `errors.As` to check for a `*coc.APIError` and inspect its status code and reason.

---

[MIT License](LICENSE). Usage is subject to Supercell's Terms and Conditions. This material is unofficial and is not endorsed by Supercell. See the [Fan Content Policy](https://www.supercell.com/fan-content-policy).
