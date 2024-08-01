package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/promiseofcake/artifactsmmo-cli/internal/actions"
)

func main() {
	token := os.Getenv("MMO_TOKEN")
	if token == "" {
		log.Fatal("no token provided")
	}
	runner, err := actions.NewDefaultRunner(token)
	if err != nil {
		log.Fatal(err)
	}

	// setup character
	character := os.Getenv("MMO_CHARACTER")
	if character == "" {
		log.Fatal("no character")
	}

	// gather loop
	ctx := context.Background()
	for {
		fmt.Println("about to gather")
		resp, err := runner.Gather(ctx, character)
		if err != nil {
			log.Fatal(err)
		}

		sec := resp.GetRemainingCooldown()
		fmt.Printf("gather result status: (%d), xp: %d, items: %v, cooldown: %d \n",
			resp.StatusCode,
			resp.SkillInfo.Xp,
			resp.SkillInfo.Items,
			sec,
		)
		time.Sleep(time.Duration(sec) * time.Second)
	}
}
