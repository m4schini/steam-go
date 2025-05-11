package steam

import (
	"github.com/m4schini/steam-go/internal"
	"net/http"
)

type ApiClientOption func(client *internal.ApiClient)

func WithHttpClient(client *http.Client) ApiClientOption {
	return func(apiClient *internal.ApiClient) {
		apiClient.Client = client
	}
}

func WithBaseUrl(baseUrl string) ApiClientOption {
	return func(apiClient *internal.ApiClient) {
		apiClient.BaseURL = baseUrl
	}
}
