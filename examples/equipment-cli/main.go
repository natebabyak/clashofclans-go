package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"sync"

	coc "github.com/natebabyak/clashofclans-go"
)

const (
	topPlayers = 100
	workers    = 2
)

type result struct {
	hero  string
	combo string
}

type comboCount struct {
	combo string
	count int
}

func main() {
	token, err := apiToken(".env")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	client := coc.NewClient(token)
	limit := topPlayers
	rankings, err := client.GetPlayerRankings(ctx, "global", &coc.PagingOptions{Limit: &limit})
	if err != nil {
		log.Fatalf("get global player rankings: %v", err)
	}

	jobs := make(chan string)
	results := make(chan result)
	var wg sync.WaitGroup
	for range workers {
		wg.Go(func() {
			for tag := range jobs {
				player, err := client.GetPlayer(ctx, tag)
				if err != nil {
					log.Printf("get player %s: %v", tag, err)
					continue
				}
				for _, hero := range player.Heroes {
					if len(hero.Equipment) < 2 {
						continue
					}
					results <- result{hero: hero.Name, combo: equipmentCombo(hero)}
				}
			}
		})
	}

	go func() {
		for _, player := range rankings.Items {
			jobs <- player.Tag
		}
		close(jobs)
		wg.Wait()
		close(results)
	}()

	combosByHero := make(map[string]map[string]int)
	for item := range results {
		if combosByHero[item.hero] == nil {
			combosByHero[item.hero] = make(map[string]int)
		}
		combosByHero[item.hero][item.combo]++
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

func apiToken(path string) (string, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read %s: %w", path, err)
	}

	for line := range strings.SplitSeq(string(contents), "\n") {
		line = strings.TrimSpace(line)
		line = strings.TrimPrefix(line, "export ")
		key, value, found := strings.Cut(line, "=")
		if found && strings.TrimSpace(key) == "COC_API_KEY" {
			value = strings.Trim(strings.TrimSpace(value), "\"'")
			if value != "" {
				return value, nil
			}
		}
	}

	return "", fmt.Errorf("COC_API_KEY is required in %s", path)
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
