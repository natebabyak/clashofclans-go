package coc_test

import (
	"context"
	"fmt"
	"log"

	coc "github.com/natebabyak/clashofclans-go"
)

func ExampleNewClient() {
	client, err := coc.NewClient("YOUR_API_TOKEN")
	if err != nil {
		log.Fatal(err)
	}
	clan, err := client.GetClan("#2PP")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(clan.Name)
}

func ExampleClient_GetPlayerRankings() {
	client, err := coc.NewClient("YOUR_API_TOKEN")
	if err != nil {
		log.Fatal(err)
	}
	page, err := client.GetPlayerRankings(coc.Global, coc.Limit(100))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(page.Items)
	if cursor := page.Paging.Cursors.After; cursor != "" {
		next, err := client.GetPlayerRankings(coc.Global, coc.Limit(100), coc.After(cursor))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(next.Items)
	}
}

func ExampleClient_ClanMembers() {
	client, err := coc.NewClient("YOUR_API_TOKEN")
	if err != nil {
		log.Fatal(err)
	}
	for member, err := range client.ClanMembers("#2PP", coc.Limit(50)) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(member.Name)
	}
}

func ExampleWithContext() {
	client, err := coc.NewClient("YOUR_API_TOKEN")
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	player, err := client.GetPlayer("#ABC", coc.WithContext(ctx))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(player.Name)
}

func ExampleClient_SearchClans() {
	client, err := coc.NewClient("YOUR_API_TOKEN")
	if err != nil {
		log.Fatal(err)
	}
	clans, err := client.SearchClans(coc.ClanFilter{
		Name:       "Example",
		MinMembers: 20,
		LabelIDs:   []int{1, 2},
	}, coc.Limit(50))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(clans.Items)
}
