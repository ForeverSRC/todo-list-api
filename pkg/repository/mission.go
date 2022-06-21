package repository

import (
	"context"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"github.com/ForeverSRC/todo-list-api/pkg/vo"
)

type MissionCreator interface {
	CreateMission(ctx context.Context, mission model.Mission) (string, error)
}

type MissionLister interface {
	GetMissions(ctx context.Context, query vo.MissionListQuery) (*model.MissionList, error)
}

type MissionGetter interface {
	GetMission(ctx context.Context, id string) (*model.Mission, error)
}

type MissionUpdater interface {
	UpdateMission(ctx context.Context, id string, mission model.Mission) error
}

type MissionItemAdder interface {
	AddItemToMission(ctx context.Context, mid string, itemId string) error
}

type MissionItemDeleter interface {
	DeleteItemFromMission(ctx context.Context, mid string, itemId string) error
}
