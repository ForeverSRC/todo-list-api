package update

import (
	"context"
	"time"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"github.com/ForeverSRC/todo-list-api/pkg/repository"
)

type Service interface {
	UpdateMission(ctx context.Context, id string, value model.MissionVo) error
}

type service struct {
	repo repository.MissionUpdater
}

func NewService(r repository.MissionUpdater) Service {
	return &service{
		repo: r,
	}
}

func (s *service) UpdateMission(ctx context.Context, id string, value model.MissionVo) error {
	m := model.Mission{
		MissionVo: value,
	}

	now := time.Now()
	m.UpdateTime = &now

	return s.repo.UpdateMission(ctx, id, m)
}
