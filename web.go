package steam

import (
	"context"
	"github.com/m4schini/steam-go/internal"
	"github.com/m4schini/steam-go/webapi"
)

type WebApi struct {
	APIKey string
	internal.ApiClient
}

func NewWebApiClient(apiKey string, opts ...ApiClientOption) *WebApi {
	apiClient := internal.DefaultWebApiClient()
	for _, opt := range opts {
		opt(&apiClient)
	}

	return &WebApi{
		apiKey,
		apiClient,
	}
}

func (w *WebApi) GetPlayerSummaries(ctx context.Context, steamIds ...string) (*webapi.PlayerSummariesResponse, error) {
	return webapi.GetPlayerSummaries(ctx, w.ApiClient.Client, w.BaseURL, w.APIKey, steamIds...)
}

func (w *WebApi) GetOwnedGames(ctx context.Context, steamID string, includeAppInfo bool) (*webapi.OwnedGamesResponse, error) {
	return webapi.GetOwnedGames(ctx, w.ApiClient.Client, w.BaseURL, w.APIKey, steamID, includeAppInfo)
}
