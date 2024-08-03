package client_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/promiseofcake/artifactsmmo-go-client/client"
)

func TestAddAuthorizationTokenRequestEditor_TokenSet(t *testing.T) {
    // Given
    token := "mock_token_value"
    os.Setenv("ARTIFACTSMMO_TOKEN", token)
    defer os.Unsetenv("ARTIFACTSMMO_TOKEN")

    req := &http.Request{
        Header: make(http.Header),
    }

    // When
    err := client.AddAuthorizationTokenRequestEditor(context.Background(), req)

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

func TestAddAuthorizationTokenRequestEditor_TokenUnset(t *testing.T) {
    // Given
    os.Unsetenv("ARTIFACTSMMO_TOKEN")

    req := &http.Request{
        Header: make(http.Header),
    }

    // When
    err := client.AddAuthorizationTokenRequestEditor(context.Background(), req)

    // Then
    if err == nil {
        t.Fatal("expected error, got nil")
    }
    if err.Error() != "ARTIFACTSMMO_TOKEN is unset" {
        t.Fatalf("expected error message %q, got %q", "ARTIFACTSMMO_TOKEN is unset", err.Error())
    }
}
