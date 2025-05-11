package internal

import (
	"github.com/m4schini/steam-go/model"
	"net/http"
)

type ApiClient struct {
	Client  *http.Client
	BaseURL string
}

func DefaultCuratorApiClient() ApiClient {
	return ApiClient{
		Client:  http.DefaultClient,
		BaseURL: model.SteamPoweredBaseUrl,
	}
}

func DefaultWebApiClient() ApiClient {
	return ApiClient{
		Client:  http.DefaultClient,
		BaseURL: model.SteamPoweredWebBaseUrl,
	}
}
