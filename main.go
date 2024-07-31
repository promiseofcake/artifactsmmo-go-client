package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/promiseofcake/artifactsmmo-cli/internal/client"
)

type AuthTransport struct {
	Token string
}

func (b AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	clonedReq := req.Clone(req.Context())
	clonedReq.Header.Set("Authorization", "Bearer "+b.Token)
	return http.DefaultTransport.RoundTrip(clonedReq)
}
func main() {
	token := os.Getenv("MMO_TOKEN")
	if token == "" {
		log.Fatal("no token")
	}

	transport := AuthTransport{
		Token: token,
	}

	cfg := client.NewConfiguration()
	cfg.Host = "api.artifactsmmo.com"
	cfg.Scheme = "https"
	cfg.HTTPClient = &http.Client{
		Transport: transport,
	}

	c := client.NewAPIClient(cfg)
	s, resp, err := c.MyCharactersAPI.GetCharacterLogsMyNameLogsGet(context.Background(), "Ragnor").Execute()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s, resp)
}
