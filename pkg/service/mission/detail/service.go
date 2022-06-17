package detail

import (
	"context"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"github.com/ForeverSRC/todo-list-api/pkg/repository"
)

type Service interface {
	GetMission(ctx context.Context, id string) (*model.Mission, error)
}

type Repository interface {
	repository.MissionGetter
	repository.ItemLister
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) GetMission(ctx context.Context, id string) (*model.Mission, error) {
	mission, err := s.repo.GetMission(ctx, id)
	if err != nil {
		return nil, err
	}

	items, err := s.repo.FetchItemsByIds(ctx, mission.Items)
	if err != nil {
		mission.ItemDetails = map[string]model.Item{}
		return mission, nil
	}

	mission.ItemDetails = items

	mission.TotalItems = len(items)
	finished := 0
	for _, v := range items {
		if *v.State == model.ItemStateFinished {
			finished++
		}
	}

	mission.FinishedItems = finished
	return mission, nil
}
