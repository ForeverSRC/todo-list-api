package managing

import (
	"context"
	"time"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"github.com/ForeverSRC/todo-list-api/pkg/repository"
)

type Service interface {
	ChangeItemState(ctx context.Context, req *Request) error
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

func (s *service) ChangeItemState(ctx context.Context, req *Request) error {
	item, err := s.repo.GetItem(ctx, req.Id)
	if err != nil {
		return err
	}

	item.State = req.State
	if req.State == model.ItemStateFinished {
		item.FinishTime = time.Now()
	}

	err = s.repo.UpdateItem(ctx, req.Id, *item)
	if err != nil {
		return err
	}

	return nil
}
