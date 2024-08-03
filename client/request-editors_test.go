package client_test

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"testing"

	"github.com/promiseofcake/artifactsmmo-go-client/client"
)

func TestAddAuthorizationTokenRequestEditor(t *testing.T) {
    // Given
    token := "mock_token_value"

    req := &http.Request{
        Header: make(http.Header),
    }

    // When
    requestEditor := client.NewAuthorizationBearerRequestEditor(token)
    err := requestEditor(context.Background(), req)

    // Then
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    expectedHeader := fmt.Sprintf("Bearer %s", token)
    authHeader := req.Header.Get("Authorization")
    if authHeader != expectedHeader {
        t.Fatalf("expected Authorization header %q, got %q", expectedHeader, authHeader)
    }
}


func TestAddBasicAuthorizationRequestEditor(t *testing.T) {
    // Given
    username := "mock_username_value"
    password := "mock_password_value"
    join := ':'

    req := &http.Request{
        Header: make(http.Header),
    }

    // When
    requestEditor := client.NewBasicAuthorizationRequestEditor(username, password, join, base64.StdEncoding)
    err := requestEditor(context.Background(), req)

    // Then
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }


    base64Encode := base64.StdEncoding.EncodeToString([]byte(username + string(join) + password))
    expectedHeader := fmt.Sprintf("Basic %s", base64Encode)
    authHeader := req.Header.Get("Authorization")
    if authHeader != expectedHeader {
        t.Fatalf("expected Authorization header %q, got %q", expectedHeader, authHeader)
    }
}
