package client_test

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"testing"

	"github.com/promiseofcake/artifactsmmo-go-client/client"
	"github.com/stretchr/testify/require"
)

func Test_NewBearerAuthorizationRequestFunc(t *testing.T) {
	token := "mock_token_value"

	req := &http.Request{
		Header: make(http.Header),
	}

	requestEditor := client.NewBearerAuthorizationRequestFunc(token)
	err := requestEditor(context.Background(), req)
	require.NoError(t, err)

	expectedHeader := fmt.Sprintf("Bearer %s", token)
	authHeader := req.Header.Get("Authorization")
	require.Equal(t, expectedHeader, authHeader)
}

func Test_NewBasicAuthorizationRequestFunc(t *testing.T) {
	username := "mock_username_value"
	password := "mock_password_value"

	req := &http.Request{
		Header: make(http.Header),
	}

	requestEditor := client.NewBasicAuthorizationRequestFunc(username, password)
	err := requestEditor(context.Background(), req)
	require.NoError(t, err)

	base64Encode := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	expectedHeader := fmt.Sprintf("Basic %s", base64Encode)
	authHeader := req.Header.Get("Authorization")
	require.Equal(t, expectedHeader, authHeader)
}
