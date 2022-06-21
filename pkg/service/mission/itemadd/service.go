package itemadd

import (
	"context"

	"github.com/ForeverSRC/todo-list-api/pkg/repository"
)

type Service interface {
	AddItem(ctx context.Context, mid string, itemId string) error
}

type service struct {
	repo repository.MissionItemAdder
}

func NewService(r repository.MissionItemAdder) Service {
	return &service{repo: r}
}

func (s *service) AddItem(ctx context.Context, mid string, itemId string) error {
	return s.repo.AddItemToMission(ctx, mid, itemId)
}
