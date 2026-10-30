package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
	"strings"

	coc "github.com/natebabyak/clashofclans-go"
)

type comboCount struct {
	combo string
	count int
}

func main() {
	token := flag.String("token", "", "Clash of Clans API token (required)")
	count := flag.Int("players", 100, "number of top global players to inspect")
	flag.Parse()
	client, err := coc.NewClient(*token)
	if err != nil {
		log.Fatal(err)
	}

	rankings, err := client.GetPlayerRankings(coc.Global, coc.Limit(*count))
	if err != nil {
		log.Fatal(err)
	}

	combosByHero := make(map[string]map[string]int)
	for _, ranking := range rankings.Items {
		player, err := client.GetPlayer(ranking.Tag)
		if err != nil {
			log.Printf("get player %s: %v", ranking.Tag, err)
			continue
		}
		for _, hero := range player.Heroes {
			if len(hero.Equipment) < 2 {
				continue
			}
			if combosByHero[hero.Name] == nil {
				combosByHero[hero.Name] = make(map[string]int)
			}
			combosByHero[hero.Name][equipmentCombo(hero)]++
		}
	}

	heroes := make([]string, 0, len(combosByHero))
	for hero := range combosByHero {
		heroes = append(heroes, hero)
	}
	sort.Strings(heroes)

	for _, hero := range heroes {
		fmt.Printf("%s\n", hero)
		for i, combo := range sortedCombos(combosByHero[hero]) {
			if i == 5 {
				break
			}
			fmt.Printf("  %d. %s (%d players)\n", i+1, combo.combo, combo.count)
		}
	}
}

func equipmentCombo(hero coc.PlayerItemLevel) string {
	equipment := make([]string, len(hero.Equipment))
	for i, item := range hero.Equipment {
		equipment[i] = item.Name
	}
	sort.Strings(equipment)
	return strings.Join(equipment, " + ")
}

func sortedCombos(counts map[string]int) []comboCount {
	combos := make([]comboCount, 0, len(counts))
	for combo, count := range counts {
		combos = append(combos, comboCount{combo: combo, count: count})
	}
	sort.Slice(combos, func(i, j int) bool {
		if combos[i].count != combos[j].count {
			return combos[i].count > combos[j].count
		}
		return combos[i].combo < combos[j].combo
	})
	return combos
}
