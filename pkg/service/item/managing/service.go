package managing

import (
	"context"
	"time"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"github.com/ForeverSRC/todo-list-api/pkg/repository"
	"github.com/ForeverSRC/todo-list-api/pkg/vo"
)

type Service interface {
	ChangeItemState(ctx context.Context, req *vo.ItemManageRequest) error
}

type Repository interface {
	repository.ItemGetter
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

func (s *service) ChangeItemState(ctx context.Context, req *vo.ItemManageRequest) error {
	item := model.Item{
		State: req.State.Pointer(),
	}

	now := time.Now()
	item.UpdateTime = now
	if req.State == model.ItemStateFinished {
		item.FinishTime = now
	}

	err := s.repo.UpdateItem(ctx, req.Id, item)
	if err != nil {
		return err
	}

	return nil
}
