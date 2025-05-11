package curator

import (
	"context"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/m4schini/steam-go/internal"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func GetReviews(ctx context.Context, client *http.Client, baseUrl, curatorId string, start, pageCount int) (<-chan Review, int, error) {
	log := internal.Logger("curator").With(zap.String("curatorId", curatorId))
	out := make(chan Review, pageCount)

	log.Debug("initial pagination request")
	reviews, totalCount, err := PaginateReviews(ctx, client, baseUrl, curatorId, start, pageCount)
	if err != nil {
		return nil, 0, err
	}

	go func() {
		defer close(out)

		for start < totalCount {
			start += pageCount
			for _, review := range reviews {
				out <- review
			}

			log.Debug("paginate to next reviews page", zap.Int("start", start))
			reviews, totalCount, err = PaginateReviews(ctx, client, baseUrl, curatorId, start, pageCount)
			if err != nil {
				log.Error("GetReviews pagination failed", zap.Error(err))
				return
			}
		}
	}()

	return out, totalCount, nil
}

func PaginateReviews(ctx context.Context, client *http.Client, baseUrl, curatorId string, start, count int) ([]Review, int, error) {
	log := internal.Logger("curator").With(zap.String("curatorId", curatorId), zap.Int("start", start), zap.Int("count", count))
	req, err := NewGetCuratorReviewsRequest(ctx, baseUrl, curatorId, start, count)
	if err != nil {
		return nil, 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	var response GetReviewsResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, 0, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response.ResultsHtml))
	if err != nil {
		return nil, response.TotalCount, err
	}

	reviews := make([]Review, 0, count)
	doc.Find(`div.recommendation`).Each(func(i int, div *goquery.Selection) {
		var (
			log    = log
			review Review
		)
		log.Debug("discovered review")

		div.Find(`a[data-ds-appid]`).First().Each(func(i int, appDiv *goquery.Selection) {
			review.AppId, _ = appDiv.Attr("data-ds-appid")
			log = log.With(zap.String("appId", review.AppId))
			log.Debug("extracted app id")
		})
		div.Find(`div.recommendation_desc`).First().Each(func(i int, div *goquery.Selection) {
			review.ReviewContent = strings.TrimSpace(div.Text())
			log.Debug("extracted review content", zap.Int("size", len(review.ReviewContent)))
		})
		div.Find(`div.recommendation_readmore a[target]`).First().Each(func(i int, div *goquery.Selection) {
			review.FullReviewUrl, _ = div.Attr("href")
			log.Debug("extracted full review url", zap.String("url", review.FullReviewUrl))
		})
		div.Find(`div.recommendation_type_ctn > span`).First().Each(func(i int, span *goquery.Selection) {
			switch {
			case span.HasClass("color_recommended"):
				review.Recommendation = Recommended
			case span.HasClass("color_informational"):
				review.Recommendation = Informative
			case span.HasClass("color_not_recommended"):
				review.Recommendation = NotRecommended
			default:
				review.Recommendation = Unknown
			}
			log.Debug("extracted recommendation", zap.String("class", span.AttrOr("class", "")), zap.Any("recommendation", review.Recommendation))
		})

		log.Debug("review extracted")
		reviews = append(reviews, review)
	})

	return reviews, response.TotalCount, nil
}
