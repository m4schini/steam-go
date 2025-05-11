package steam

import (
	"context"
	"steam-go/curator"
	"steam-go/internal"
)

type CuratorApi struct {
	internal.ApiClient
}

func NewCuratorApiClient(opts ...ApiClientOption) *CuratorApi {
	apiClient := internal.DefaultCuratorApiClient()
	for _, opt := range opts {
		opt(&apiClient)
	}

	return &CuratorApi{apiClient}
}

func (c *CuratorApi) GetReviews(ctx context.Context, curatorId string, start, count, limit int) (reviews <-chan curator.Review, totalCount int, err error) {
	ctx, cancel := context.WithCancel(ctx)
	ch, total, err := curator.GetReviews(ctx, c.Client, c.BaseURL, curatorId, start, count)
	if err != nil {
		cancel()
		return nil, 0, err
	}

	out := make(chan curator.Review, count)
	go func() {
		defer cancel()
		defer close(out)
		var i int
		for review := range ch {
			i++
			if i > limit {
				return
			}
			out <- review
		}
	}()

	return out, total - (start - 1) - limit, nil
}

func (c *CuratorApi) GetAllReviews(ctx context.Context, curatorId string) (reviews <-chan curator.Review, totalCount int, err error) {
	return curator.GetReviews(ctx, c.Client, c.BaseURL, curatorId, 0, 100)
}

func Filter(predicates ...func(review curator.Review) bool) func(in <-chan curator.Review, totalCount int, err error) (<-chan curator.Review, int, error) {
	return func(in <-chan curator.Review, totalCount int, err error) (<-chan curator.Review, int, error) {
		if err != nil {
			return nil, totalCount, err
		}
		out := make(chan curator.Review, totalCount)

		go func() {
			defer close(out)
			for review := range in {
				skip := false
				for _, p := range predicates {
					if !p(review) {
						skip = true
						break
					}
				}
				if skip {
					continue
				}
				out <- review
			}

		}()
		return out, totalCount, nil
	}
}
