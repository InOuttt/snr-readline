package service

import (
	"context"
	"encoding/json"

	"github.com/InOuttt/snr/insert/model"
	"github.com/InOuttt/snr/insert/repository"
)

type FeedService interface {
	CreateFeed(ctx context.Context, data string) error
}

type feedService struct {
	fr repository.FeedRepository
}

func NewFeedService(fr repository.FeedRepository) FeedService {
	return feedService{
		fr: fr,
	}
}

func (fs feedService) CreateFeed(ctx context.Context, data string) error {

	var model model.Feed
	if err := json.Unmarshal([]byte(data), &model); err != nil {
		return err
	}

	if err := fs.fr.CreateFeed(ctx, model); err != nil {
		return err
	}

	return nil
}
