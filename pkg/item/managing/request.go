package managing

import "github.com/ForeverSRC/todo-list-api/pkg/model"

type Request struct {
	Id    string          `json:"id"`
	State model.ItemState `json:"state"`
}
