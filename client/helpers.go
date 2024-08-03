package client

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
)

// NewBearerAuthorizationRequestFunc returns a RequestEditorFn that adds the authorization bearer token
// to the request
func NewBearerAuthorizationRequestFunc(token string) RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}
}

// NewBasicAuthorizationRequestEditor returns a RequestEditorFn that adds the basic authorization headers
// to the request
func NewBasicAuthorizationRequestFunc(username string, password string) RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		encodedAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		req.Header.Add("Authorization", "Basic "+encodedAuth)
		return nil
	}
}
