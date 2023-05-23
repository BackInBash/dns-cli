package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/BackInBash/dns-cli/internal/api"
	"github.com/golang-jwt/jwt/v5"
)

func createClient() (*api.Client, error) {
	apiClient, err := api.NewClient(dnsApiUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create api client: %v", err)
	}
	apiClient.RequestEditors = []api.RequestEditorFn{
		func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", authenticationToken))
			return nil
		},
	}
	return apiClient, nil
}

func parseClaimFromJWT(tokenString string, claim string) (string, error) {
	parser := jwt.NewParser()
	token, _, err := parser.ParseUnverified(tokenString, &jwt.MapClaims{})
	if err != nil {
		return "", err
	}
	claims := *token.Claims.(*jwt.MapClaims)
	if id, ok := claims[claim]; ok {
		return id.(string), nil
	}
	return "", fmt.Errorf("claim %q not found in the token", claim)
}
