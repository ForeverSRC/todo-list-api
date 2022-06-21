package detail

import (
	"context"
	"testing"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"github.com/ForeverSRC/todo-list-api/pkg/vo"
	"github.com/stretchr/testify/assert"
)

type mockRepo struct {
	t               *testing.T
	testMission     *model.Mission
	testItemDetails map[string]model.Item
}

func (m *mockRepo) GetMission(ctx context.Context, id string) (*model.Mission, error) {
	return m.testMission, nil
}

func (m *mockRepo) FetchItems(ctx context.Context, query *vo.ItemListQuery) (*model.ItemList, error) {
	return nil, nil
}

func (m *mockRepo) FetchItemsByIds(ctx context.Context, ids []string) (map[string]model.Item, error) {
	return m.testItemDetails, nil
}

func TestGetMissionDetail_success(t *testing.T) {
	repo := &mockRepo{
		t: t,
	}

	mv := model.MissionVo{
		Title:    "test-mission",
		Priority: model.P1,
		Items:    []string{"item-a", "item-b"},
	}

	detail := "test detail"
	mv.Detail = &detail

	testMission := &model.Mission{
		MissionVo: mv,
	}

	repo.testMission = testMission

	itemA := model.Item{
		State: model.ItemStateFinished.Pointer(),
	}

	itemB := model.Item{
		State: model.ItemStateUnFinished.Pointer(),
	}

	testItems := map[string]model.Item{"item-a": itemA, "item-b": itemB}
	repo.testItemDetails = testItems

	srv := NewService(repo)
	mission, err := srv.GetMission(context.TODO(), "test-mid")
	assert.NoError(t, err)
	assert.Equal(t, 2, mission.TotalItems)
	assert.Equal(t, 1, mission.FinishedItems)

}
