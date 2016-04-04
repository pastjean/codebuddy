package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/pastjean/codebuddy/models"
	"github.com/pastjean/codebuddy/pairs"
)

func main() {
	teams := models.TeamsMap{}
	err := json.NewDecoder(os.Stdin).Decode(&teams)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Pairs are:\n\n")

	for k, v := range teams {
		p := pairs.GeneratePairs(pairs.Shuffle(v, time.Now().UnixNano()))
		fmt.Printf("For %v team:\n", k)

		for _, tuple := range p {
			fmt.Printf("- %v\n", strings.Join(tuple, ", "))
		}
		fmt.Print("\n")
	}
}
