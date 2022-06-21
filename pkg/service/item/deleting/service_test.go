package deleting

import (
	"context"
	"testing"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"github.com/stretchr/testify/assert"
)

type mockRepo struct {
	t           *testing.T
	deletedItem model.Item
}

func (m *mockRepo) DeleteItem(ctx context.Context, id string) (model.Item, error) {
	return m.deletedItem, nil
}

func (m *mockRepo) DeleteItemFromMission(ctx context.Context, mid string, itemId string) error {
	assert.Equal(m.t, m.deletedItem.RelatedMission, mid)
	return nil
}

func TestDeleteItem_success(t *testing.T) {
	tid := "test-item-id"
	repo := &mockRepo{
		t: t,
		deletedItem: model.Item{
			Id: tid,
			ItemVo: model.ItemVo{
				RelatedMission: "related-mission-id",
			},
		},
	}

	s := NewService(repo)
	err := s.DeleteItem(context.TODO(), tid)
	assert.NoError(t, err)
}

func TestDeleteItemWithNoRelatedMission_success(t *testing.T) {
	tid := "test-item-id"
	repo := &mockRepo{
		t: t,
		deletedItem: model.Item{
			Id:     tid,
			ItemVo: model.ItemVo{},
		},
	}

	s := NewService(repo)
	err := s.DeleteItem(context.TODO(), tid)
	assert.NoError(t, err)
}
