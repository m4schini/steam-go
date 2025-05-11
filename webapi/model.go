package webapi

type PlayerSummariesResponse struct {
	Response struct {
		Players []Player `json:"players"`
	} `json:"response"`
}

type Player struct {
	SteamID      string `json:"steamid"`
	PersonaName  string `json:"personaname"`
	ProfileURL   string `json:"profileurl"`
	Avatar       string `json:"avatar"`
	AvatarMedium string `json:"avatarmedium"`
	AvatarFull   string `json:"avatarfull"`
}

type OwnedGamesResponse struct {
	GameCount int       `json:"game_count"`
	Games     []AppInfo `json:"games"`
}

type AppInfo struct {
	AppId                    string `json:"appid"`
	Name                     string `json:"name"`
	Playtime2Weeks           int    `json:"playtime_2weeks"`
	PlaytimeTotal            int    `json:"playtime_forever"`
	ImgIcon                  string `json:"img_icon_url"`
	ImgLogoUrl               string `json:"img_logo_url"`
	HasVisibleCommunityStats bool   `json:"has_community_visible_stats"`
}
