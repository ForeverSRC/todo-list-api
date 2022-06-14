package vo

import (
	"fmt"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	mathutil "github.com/ForeverSRC/todo-list-api/pkg/utils/math"
)

const (
	defaultPage     = 1
	defaultPageSize = 5
)

type ItemListQuery struct {
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

type ItemManageRequest struct {
	Id    string          `json:"id"`
	State model.ItemState `json:"state"`
}

type ItemEditRequest struct {
	Description string `json:"description"`
}
