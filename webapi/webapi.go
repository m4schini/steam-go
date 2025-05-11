package webapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	neturl "net/url"
	"strconv"
	"strings"
)

func GetPlayerSummaries(ctx context.Context, client *http.Client, baseUrl, apiKey string, steamIds ...string) (*PlayerSummariesResponse, error) {
	endpoint := fmt.Sprintf("%s/ISteamUser/GetPlayerSummaries/v0002/", baseUrl)

	params := neturl.Values{}
	params.Set("key", apiKey)
	params.Set("steamids", neturl.QueryEscape(strings.Join(steamIds, ",")))
	url := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result PlayerSummariesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func GetOwnedGames(ctx context.Context, client *http.Client, baseUrl, apiKey, steamID string, includeAppInfo bool) (*OwnedGamesResponse, error) {
	endpoint := fmt.Sprintf("%s/IPlayerService/GetOwnedGames/v0001/", baseUrl)

	params := neturl.Values{}
	params.Set("key", apiKey)
	params.Set("steamid", steamID)
	params.Set("include_appinfo", strconv.FormatBool(includeAppInfo))
	url := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result OwnedGamesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
