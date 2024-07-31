package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/promiseofcake/artifactsmmo-cli/client"
)

func main() {
	token := os.Getenv("MMO_TOKEN")
	if token == "" {
		log.Fatal("no token")
	}

	ctx := context.Background()

	c, err := client.NewClientWithResponses("https://api.artifactsmmo.com", client.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	}))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.ActionMoveMyNameActionMovePostWithResponse(ctx, "Rangor", client.ActionMoveMyNameActionMovePostJSONRequestBody{
		X: 5,
		Y: 5,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resp.Body))
}
