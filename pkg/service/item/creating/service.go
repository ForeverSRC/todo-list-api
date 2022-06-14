package creating

import (
	"context"
	"time"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"github.com/ForeverSRC/todo-list-api/pkg/repository"
)

type Service interface {
	CreateItem(ctx context.Context, item model.ItemVo) error
}

type Repository interface {
	repository.ItemCreator
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) CreateItem(ctx context.Context, item model.ItemVo) error {
	i := model.Item{
		ItemVo: item,
	}
	i.State = model.ItemStateUnFinished.Pointer()

	now := time.Now()
	i.CreateTime = now
	i.UpdateTime = now

	return s.repo.InsertItem(ctx, i)
}
