package creating

import (
	"context"
	"fmt"
	"testing"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"github.com/stretchr/testify/assert"
)

type mockItemCreator struct {
	t      *testing.T
	itemId string
	err    error
}

func (m *mockItemCreator) InsertItem(ctx context.Context, item model.Item) (string, error) {
	if !m.verify(item) {
		return "", fmt.Errorf("error")
	}
	return m.itemId, m.err
}

func (m mockItemCreator) verify(item model.Item) bool {
	b := assert.Equal(m.t, model.ItemStateUnFinished, *item.State) &&
		assert.Equal(m.t, *item.CreateTime, *item.UpdateTime)

	return b
}

func TestCreateItem_success(t *testing.T) {
	repo := &mockItemCreator{
		t:      t,
		itemId: "test-id",
		err:    nil,
	}

	s := NewService(repo)
	item := model.ItemVo{}
	id, err := s.CreateItem(context.TODO(), item)

	assert.Nil(t, err)
	assert.Equal(t, id, repo.itemId)
}
