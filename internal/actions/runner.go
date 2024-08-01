package actions

import (
	"context"
	"fmt"

	"github.com/promiseofcake/artifactsmmo-cli/client"
)

const (
	username = "Rangor"
)

type Runner struct {
	Client *client.ClientWithResponses
}

type Response struct {
	Cooldown   int64
	StatusCode int
	StatusText string
}

func (r *Runner) Move(ctx context.Context, x, y int) (*Response, error) {
	resp, err := r.Client.ActionMoveMyNameActionMovePostWithResponse(ctx, username, client.ActionMoveMyNameActionMovePostJSONRequestBody{
		X: x,
		Y: y,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to move: %w", err)
	}

	return &Response{
		Cooldown:   int64(resp.JSON200.Data.Cooldown.RemainingSeconds),
		StatusCode: resp.StatusCode(),
		StatusText: resp.Status(),
	}, nil
}

func (r *Runner) Gather(ctx context.Context) (*Response, error) {
	resp, err := r.Client.ActionGatheringMyNameActionGatheringPostWithResponse(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to gather: %w", err)
	}

	return &Response{
		Cooldown:   int64(resp.JSON200.Data.Cooldown.RemainingSeconds),
		StatusCode: resp.StatusCode(),
		StatusText: resp.Status(),
	}, nil
}
