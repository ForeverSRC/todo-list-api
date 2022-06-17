package vo

import (
	"github.com/ForeverSRC/todo-list-api/pkg/model"
	mathutil "github.com/ForeverSRC/todo-list-api/pkg/utils/math"
)

type MissionListQuery struct {
	State      *model.MissionState `form:"state"`
	Descending bool                `form:"descending"`
	Page       int64               `form:"page"`
	PageSize   int64               `form:"pageSize"`
}

func (q *MissionListQuery) CheckAndFix() error {
	q.Page = mathutil.Max(q.Page, defaultPage)
	if q.PageSize < 1 {
		q.PageSize = defaultPageSize
	}

	return nil
}

type ItemAddType int

const (
	AddExistItemType ItemAddType = 1
	AddNewItemType   ItemAddType = 2
)

func (ia ItemAddType) Valid() bool {
	switch ia {
	case AddExistItemType, AddNewItemType:
		return true
	default:
		return false
	}
}

type MissionAddItemReq struct {
	AddType  ItemAddType   `json:"addType"`
	ItemId   string        `json:"itemId"`
	ItemInfo *model.ItemVo `json:"itemInfo,omitempty"`
}
