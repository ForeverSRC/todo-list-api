package deleting

import (
	"context"

	"github.com/ForeverSRC/todo-list-api/pkg/repository"
)

type Service interface {
	DeleteItem(ctx context.Context, id string) error
}

type Repository interface {
	repository.ItemDeleter
	repository.MissionItemDeleter
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) DeleteItem(ctx context.Context, id string) error {
	item, err := s.repo.DeleteItem(ctx, id)
	if err != nil {
		return err
	}

	if len(item.RelatedMission) > 0 {
		return s.repo.DeleteItemFromMission(ctx, item.RelatedMission, item.Id)
	}

	return nil
}
