# clash-of-clans-go

A Go client for the Clash of Clans API.

## Features

- Full API coverage
- Fully typed
- Near-zero performance overhead

## Getting Started

1. [Register for access to the Clash of Clans API](https://developer.clashofclans.com/sign-in#/register)

2. [Create a Clash of Clans API key](https://developer.clashofclans.com/account#/new-key)

```bash
go get github.com/natebabyak/clashofclans-go
```

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    coc "github.com/natebabyak/clash-of-clans-go"
)

func main() {
    client := coc.NewClient(os.Getenv("COC_API_KEY"))

    ctx := context.Background()

    clan, err := client.GetClan(ctx, "#2PP")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s has %d members!\n", clan.Name, clan.Members)
}
```

## Examples

See [`examples/equipment-cli`](examples/equipment-cli) for a small CLI that fetches top global player rankings and prints common hero equipment combos. Copy `.env.example` to `.env` and set `COC_API_KEY`.

## Legal

Usage is subject to Supercell's Terms and Conditions.

This material is unofficial and is not endorsed by Supercell. For more information see Supercell's Fan Content Policy: [www.supercell.com/fan-content-policy](https://www.supercell.com/fan-content-policy).

## License

MIT
