package creating

import (
	"context"
	"time"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"github.com/ForeverSRC/todo-list-api/pkg/repository"
)

type Service interface {
	CreateItem(ctx context.Context, item model.Item) error
}

type Repository interface {
	repository.ItemCreator
}

type service struct {
	repo Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) CreateItem(ctx context.Context, item model.Item) error {
	item.State = model.ItemStateUnFinished
	item.CreateTime = time.Now()
	return s.repo.InsertItem(ctx, item)
}
