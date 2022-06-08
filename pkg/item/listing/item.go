package listing

import (
	"fmt"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	mathutil "github.com/ForeverSRC/todo-list-api/pkg/utils/math"
)

const (
	defaultPage     = 1
	defaultPageSize = 5
)

type ItemList []model.Item

type ItemListQuery struct {
	Uid      string          `form:"uid"`
	State    model.ItemState `form:"state"`
	Page     int64           `form:"page"`
	PageSize int64           `form:"pageSize"`
}

func (q *ItemListQuery) CheckAndFix() error {
	if !q.State.ValidState() {
		return fmt.Errorf("invalid item state: %d", q.State)
	}

	q.Page = mathutil.Max(q.Page, defaultPage)
	if q.PageSize < 1 {
		q.PageSize = defaultPageSize
	}

	return nil
}
