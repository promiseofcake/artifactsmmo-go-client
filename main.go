package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/promiseofcake/artifactsmmo-cli/client"
	"github.com/promiseofcake/artifactsmmo-cli/internal/actions"
)

func main() {
	token := os.Getenv("MMO_TOKEN")
	if token == "" {
		log.Fatal("no token")
	}

	c, err := client.NewClientWithResponses("https://api.artifactsmmo.com", client.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	}))
	if err != nil {
		log.Fatal(err)
	}

	runner := &actions.Runner{
		Client: c,
	}

	ctx := context.Background()
	for {
		fmt.Println("about to gather")
		resp, gErr := runner.Gather(ctx)
		if gErr != nil {
			log.Fatal(err)
		}
		fmt.Printf("gather result status: %s (%d), cooldown: %d\n", resp.StatusText, resp.StatusCode, resp.Cooldown)
		time.Sleep(time.Duration(resp.Cooldown) * time.Second)
	}

}
