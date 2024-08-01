package actions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/promiseofcake/artifactsmmo-cli/client"
)

// Runner is an executor for Actions
type Runner struct {
	Client *client.ClientWithResponses
}

// NewDefaultRunner returns a new Actions command runner with a default client
func NewDefaultRunner(token string) (*Runner, error) {
	c, err := client.NewClientWithResponses("https://api.artifactsmmo.com", client.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to init new client: %w", err)
	}
	return &Runner{
		Client: c,
	}, nil
}

// NewRunnerWithClient returns a new Actions command runner with a pre-configured client
func NewRunnerWithClient(client *client.ClientWithResponses) *Runner {
	return &Runner{
		Client: client,
	}
}

// GetMyCharacterInfo returns current info and status about your own specific character
func (r *Runner) GetMyCharacterInfo(ctx context.Context, character string) (*MyCharacterResponse, error) {
	resp, err := r.Client.GetMyCharactersMyCharactersGetWithResponse(ctx, &client.GetMyCharactersMyCharactersGetParams{
		Page: toIntPointer(5),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch characters: %w", err)
	}

	for _, c := range resp.JSON200.Data {
		if c.Name == character {
			return &MyCharacterResponse{
				MyCharacterSchema: c,
			}, nil
		}
	}

	return nil, fmt.Errorf("failed to find character: %s", character)
}

// Fight attacks the mob at the current position for the given character
func (r *Runner) Fight(ctx context.Context, character string) (*FightResponse, error) {
	resp, err := r.Client.ActionFightMyNameActionFightPostWithResponse(ctx, character)
	if err != nil {
		return nil, fmt.Errorf("failed to attack: %w", err)
	}

	return &FightResponse{
		FightResponse: resp.JSON200.Data.Fight,
		Response: Response{
			CharacterResponse: CharacterResponse{resp.JSON200.Data.Character},
			CooldownSchema:    resp.JSON200.Data.Cooldown,
			StatusCode:        resp.StatusCode(),
			StatusText:        resp.Status(),
		},
	}, nil

}

// Gather performs resource gathering at the present position for the given character
func (r *Runner) Gather(ctx context.Context, character string) (*GatherResponse, error) {
	resp, err := r.Client.ActionGatheringMyNameActionGatheringPostWithResponse(ctx, character)
	if err != nil {
		return nil, fmt.Errorf("failed to gather: %w", err)
	}

	return &GatherResponse{
		SkillInfo: resp.JSON200.Data.Details,
		Response: Response{
			CharacterResponse: CharacterResponse{resp.JSON200.Data.Character},
			CooldownSchema:    resp.JSON200.Data.Cooldown,
			StatusCode:        resp.StatusCode(),
			StatusText:        resp.Status(),
		},
	}, nil
}

// Move changes the x, y position the given character
func (r *Runner) Move(ctx context.Context, character string, x, y int) (*Response, error) {
	resp, err := r.Client.ActionMoveMyNameActionMovePostWithResponse(ctx, character, client.ActionMoveMyNameActionMovePostJSONRequestBody{
		X: x,
		Y: y,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to move: %w", err)
	}

	return &Response{
		CharacterResponse: CharacterResponse{resp.JSON200.Data.Character},
		CooldownSchema:    resp.JSON200.Data.Cooldown,
		StatusCode:        resp.StatusCode(),
		StatusText:        resp.Status(),
	}, nil
}
