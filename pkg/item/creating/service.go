package creating

import (
	"time"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
)

type Service interface {
	CreateItem(item model.Item) error
}

type Repository interface {
	InsertItem(item model.Item) error
}

type service struct {
	repo Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) CreateItem(item model.Item) error {
	item.State = model.ItemStateUnFinished
	item.CreateTime = time.Now()
	return s.repo.InsertItem(item)
}
