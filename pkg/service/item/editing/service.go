package editing

import (
	"context"
	"time"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"github.com/ForeverSRC/todo-list-api/pkg/repository"
)

type Service interface {
	Edit(ctx context.Context, id string, value model.ItemVo) error
}

type Repository interface {
	repository.ItemUpdater
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) Edit(ctx context.Context, id string, value model.ItemVo) error {
	now := time.Now()
	i := model.Item{
		UpdateTime: &now,
		ItemVo:     value,
	}

	return s.repo.UpdateItem(ctx, id, i)
}
