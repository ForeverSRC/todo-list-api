package listing

import (
	"context"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"github.com/ForeverSRC/todo-list-api/pkg/repository"
	"github.com/ForeverSRC/todo-list-api/pkg/vo"
)

type Service interface {
	ListItems(ctx context.Context, query *vo.ItemListQuery) (model.ItemList, error)
}

type Repository interface {
	repository.ItemLister
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) ListItems(ctx context.Context, query *vo.ItemListQuery) (model.ItemList, error) {
	if err := query.CheckAndFix(); err != nil {
		return nil, err
	}

	l, err := s.repo.FetchItems(ctx, query)
	if err != nil {
		return nil, err
	}

	return l, nil

}
