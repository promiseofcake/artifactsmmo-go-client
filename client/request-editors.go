package client

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
)

const (
	TokenUnsetError    = "ARTIFACTSMMO_TOKEN is unset"
	UsernameUnsetError = "ARTIFACTSMMO_USERNAME is unset"
	PasswordUnsetError = "ARTIFACTSMMO_PASSWORD is unset"
)

// Adds the header 'Authorization' with the value: Bearer $ARTIFACTSMMO_TOKEN to the request.
// Returns an error if the "ARTIFACTSMMO_TOKEN" environment variable is unset.
func AddAuthorizationTokenRequestEditor(ctx context.Context, req *http.Request) error {
	if token, isSet := os.LookupEnv("ARTIFACTSMMO_TOKEN"); isSet {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}
	return errors.New(TokenUnsetError)
}

// Adds the header 'Authorization' with the value: Basic base64Encode($ARTIFACTSMMO_USERNAME:$ARTIFACTSMMO_PASSWORD) to the request.
// Returns an error if either environment variable is unset.
func AddBasicAuthorizationRequestEditor(ctx context.Context, req *http.Request) error {
	user, isSet := os.LookupEnv("ARTIFACTSMMO_USERNAME")
	if !isSet {
		return errors.New(UsernameUnsetError)
	}
	password, isSet := os.LookupEnv("ARTIFACTSMMO_PASSWORD")
	if !isSet {
		return errors.New(PasswordUnsetError)
	}
	base64Encode := base64.StdEncoding.EncodeToString([]byte(user + ":" + password))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64Encode))
	return nil
}
