package creating

import (
	"context"
	"time"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"github.com/ForeverSRC/todo-list-api/pkg/repository"
)

type Service interface {
	CreateMission(ctx context.Context, mission model.MissionVo) (string, error)
}

type service struct {
	repo repository.MissionCreator
}

func NewService(r repository.MissionCreator) Service {
	return &service{
		repo: r,
	}
}

func (s *service) CreateMission(ctx context.Context, mission model.MissionVo) (string, error) {
	m := &model.Mission{
		MissionVo: mission,
		State:     model.MissionStateInit.Pointer(),
	}

	now := time.Now()
	m.CreateTime = &now
	m.UpdateTime = &now

	id, err := s.repo.CreateMission(ctx, *m)
	if err != nil {
		return "", err
	}

	return id, nil
}
