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
	return s.repo.DeleteItem(ctx, id)
}
