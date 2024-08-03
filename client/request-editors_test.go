package client_test

import (
	"context"
	"encoding/base64"
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
    if err.Error() != client.TokenUnsetError {
        t.Fatalf("expected error message %q, got %q", client.TokenUnsetError, err.Error())
    }
}

func TestAddBasicAuthorizationRequestEditor_UsernameUnsetPasswordUnset(t *testing.T) {
    // Given
    os.Unsetenv("ARTIFACTSMMO_USERNAME")
    os.Unsetenv("ARTIFACTSMMO_PASSWORD")

    req := &http.Request{
        Header: make(http.Header),
    }

    // When
    err := client.AddBasicAuthorizationRequestEditor(context.Background(), req)

    // Then
    if err == nil {
        t.Fatal("expected error, got nil")
    }
    if err.Error() != client.UsernameUnsetError {
        t.Fatalf("expected error message %q, got %q", client.UsernameUnsetError, err.Error())
    }
}

func TestAddBasicAuthorizationRequestEditor_UsernameSetPasswordUnset(t *testing.T) {
    // Given
    username := "mock_username_value"
    os.Setenv("ARTIFACTSMMO_USERNAME", username)
    os.Unsetenv("ARTIFACTSMMO_PASSWORD")
    defer os.Unsetenv("ARTIFACTSMMO_USERNAME")

    req := &http.Request{
        Header: make(http.Header),
    }

    // When
    err := client.AddBasicAuthorizationRequestEditor(context.Background(), req)

    // Then
    if err == nil {
        t.Fatal("expected error, got nil")
    }
    if err.Error() != client.PasswordUnsetError {
        t.Fatalf("expected error message %q, got %q", client.PasswordUnsetError, err.Error())
    }
}

func TestAddBasicAuthorizationRequestEditor_UsernameUnsetPasswordSet(t *testing.T) {
    // Given
    password := "mock_password_value"
    os.Setenv("ARTIFACTSMMO_PASSWORD", password)
    os.Unsetenv("ARTIFACTSMMO_USERNAME")
    defer os.Unsetenv("ARTIFACTSMMO_PASSWORD")

    req := &http.Request{
        Header: make(http.Header),
    }

    // When
    err := client.AddBasicAuthorizationRequestEditor(context.Background(), req)

    // Then
    if err == nil {
        t.Fatal("expected error, got nil")
    }
    if err.Error() != client.UsernameUnsetError {
        t.Fatalf("expected error message %q, got %q", client.UsernameUnsetError, err.Error())
    }
}

func TestAddBasicAuthorizationRequestEditor(t *testing.T) {
    // Given
    username := "mock_username_value"
    password := "mock_password_value"
    os.Setenv("ARTIFACTSMMO_USERNAME", username)
    os.Setenv("ARTIFACTSMMO_PASSWORD", password)
    defer os.Unsetenv("ARTIFACTSMMO_USERNAME")
    defer os.Unsetenv("ARTIFACTSMMO_PASSWORD")

    req := &http.Request{
        Header: make(http.Header),
    }

    // When
    err := client.AddBasicAuthorizationRequestEditor(context.Background(), req)

    // Then
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }


    base64Encode := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
    expectedHeader := fmt.Sprintf("Basic %s", base64Encode)
    authHeader := req.Header.Get("Authorization")
    if authHeader != expectedHeader {
        t.Fatalf("expected Authorization header %q, got %q", expectedHeader, authHeader)
    }
}
