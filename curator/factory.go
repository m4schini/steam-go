package curator

import (
	"context"
	"fmt"
	"net/http"
)

func NewGetCuratorReviewsRequest(ctx context.Context, baseUrl, curatorId string, start, count int) (*http.Request, error) {
	if curatorId == "" {
		return nil, fmt.Errorf("curatorId is empty")
	}
	url := fmt.Sprintf("%v/curator/%v/ajaxgetfilteredrecommendations/?query&start=%v&count=%v&dynamic_data=&tagids=&sort=recent&app_types=&curations=&reset=false", baseUrl, curatorId, start, count)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:136.0) Gecko/20100101 Firefox/136.0")
	header.Set("Accept", "text/javascript, text/html, application/xml, text/xml, */*")
	header.Set("Accept-Language", "de,en-US;q=0.7,en;q=0.3")
	header.Set("X-Requested-With", "XMLHttpRequest")
	header.Set("X-Prototype-Version", "1.7")
	header.Set("Sec-Fetch-Dest", "empty")
	header.Set("Sec-Fetch-Mode", "cors")
	header.Set("Sec-Fetch-Site", "same-origin")
	header.Set("Sec-GPC", "1")
	header.Set("Pragma", "no-cache")
	header.Set("Cache-Control", "no-cache")
	header.Set("Referer", fmt.Sprintf("%v/curator/%v/", baseUrl, curatorId))
	req.Header = header
	return req, nil
}
