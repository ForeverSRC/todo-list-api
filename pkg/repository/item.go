package repository

import (
	"context"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"github.com/ForeverSRC/todo-list-api/pkg/vo"
)

type ItemGetter interface {
	GetItem(ctx context.Context, id string) (*model.Item, error)
}

type ItemCreator interface {
	InsertItem(ctx context.Context, item model.Item) (string, error)
}

type ItemLister interface {
	FetchItems(ctx context.Context, query *vo.ItemListQuery) (*model.ItemList, error)
	FetchItemsByIds(ctx context.Context, ids []string) (map[string]model.Item, error)
}

type ItemUpdater interface {
	UpdateItem(ctx context.Context, id string, item model.Item) error
}

type ItemDeleter interface {
	DeleteItem(ctx context.Context, id string) (model.Item, error)
}
