package model

import (
	"time"
)

type ItemState uint

const (
	ItemStateUnFinished ItemState = 1
	ItemStateFinished   ItemState = 2
)

var stateMap = map[ItemState]string{
	ItemStateFinished:   "finished",
	ItemStateUnFinished: "unfinished",
}

func (s ItemState) String() string {
	return stateMap[s]
}

func (s ItemState) ValidState() bool {
	switch s {
	case ItemStateFinished, ItemStateUnFinished:
		return true
	default:
		return false
	}
}

type Item struct {
	Id string `json:"id" bson:"_id"`

	ItemVo

	State ItemState `json:"state" bson:"state"`

	CreateTime time.Time `json:"createTime" bson:"create_time"`
	FinishTime time.Time `json:"finishTime" bson:"finish_time"`
}

type ItemList []Item

type ItemVo struct {
	Title       string   `json:"title" bson:"title" binding:"required"`
	Tags        []string `json:"tags" bson:"tags"`
	Score       uint     `json:"score" bson:"score"`
	Description string   `json:"description" bson:"description"`
}