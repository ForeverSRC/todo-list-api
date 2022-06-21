package listing

import (
	"context"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"github.com/ForeverSRC/todo-list-api/pkg/repository"
	"github.com/ForeverSRC/todo-list-api/pkg/vo"
)

type Service interface {
	ListMission(ctx context.Context, query vo.MissionListQuery) (*model.MissionList, error)
}

type service struct {
	repo repository.MissionLister
}

func NewService(r repository.MissionLister) Service {
	return &service{
		repo: r,
	}
}

func (s *service) ListMission(ctx context.Context, query vo.MissionListQuery) (*model.MissionList, error) {
	res, err := s.repo.GetMissions(ctx, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}
