package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
)

// Adds the header: <code>Authorization: Bearer $ARTIFACTSMMO_TOKEN</code> to the request. Returns
// an error if the "ARTIFACTSMMO_TOKEN" environment variable is unset.
func AddAuthorizationTokenRequestEditor(ctx context.Context, req *http.Request) error {
  token, isSet := os.LookupEnv("ARTIFACTSMMO_TOKEN")
  if !isSet {
    return errors.New("ARTIFACTSMMO_TOKEN is unset")
  }
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
  return nil
}
