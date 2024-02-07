package api

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func Match(ctx context.Context, client *resty.Client, apiKey, matchID string) (MatchResponse, error) {
	var results MatchResponse

	response, err := client.R().
		SetResult(&results).
		SetContext(ctx).
		SetPathParam("matchID", matchID).
		SetQueryParam("api_token", apiKey).
		Get("https://mwomercs.com/api/v1/matches/{matchID}")
	if err != nil {
		return results, err
	}

	if response.IsError() {
		return results, fmt.Errorf("%v", string(response.Body()))
	}

	return results, nil
}
