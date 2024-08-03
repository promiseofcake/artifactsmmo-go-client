package client

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
)

const (
	TokenUnsetError    = "ARTIFACTSMMO_TOKEN is unset"
	UsernameUnsetError = "ARTIFACTSMMO_USERNAME is unset"
	PasswordUnsetError = "ARTIFACTSMMO_PASSWORD is unset"
)

// Returns a request editor that adds the authorization bearer token and does not return an error.
func NewAuthorizationBearerRequestEditor(token string) RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}
}

// Returns a request editor that adds the basic authorization header and does not return an error.
func NewBasicAuthorizationRequestEditor(username string, password string, joiner rune, encoding *base64.Encoding) RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		base64Encode := encoding.EncodeToString([]byte(username + string(joiner) + password))
		req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64Encode))
		return nil
	}
}
