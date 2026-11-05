package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	coc "github.com/natebabyak/clashofclans-go"
)

func main() {
	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatal(err)
	}

	client, err := coc.NewClient(os.Getenv("COC_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	clan, err := client.GetClan("#2PPJRCR8G")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s has %d members!\n", clan.Name, clan.Members)
}
