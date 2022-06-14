package editing

import (
	"context"
	"time"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"github.com/ForeverSRC/todo-list-api/pkg/repository"
	"github.com/ForeverSRC/todo-list-api/pkg/vo"
)

type Service interface {
	Edit(ctx context.Context, id string, req vo.ItemEditRequest) error
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

func (s *service) Edit(ctx context.Context, id string, req vo.ItemEditRequest) error {
	i := model.Item{
		UpdateTime: time.Now(),
	}

	i.Description = req.Description

	return s.repo.UpdateItem(ctx, id, i)
}
