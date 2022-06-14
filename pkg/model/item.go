package model

import (
	"time"
)

type ItemState uint

const (
	ItemStateInit ItemState = iota
	ItemStateUnFinished
	ItemStateFinished
)

func (s ItemState) ValidState() bool {
	switch s {
	case ItemStateInit, ItemStateFinished, ItemStateUnFinished:
		return true
	default:
		return false
	}
}

func (s ItemState) Pointer() *ItemState {
	return &s
}

type Item struct {
	Id string `json:"id" bson:"_id,omitempty"`

	ItemVo `json:",inline" bson:",inline"`

	State *ItemState `json:"state" bson:"state,omitempty"`

	CreateTime time.Time `json:"createTime" bson:"create_time,omitempty"`
	UpdateTime time.Time `json:"updateTime" bson:"update_time,omitempty"`
	FinishTime time.Time `json:"finishTime" bson:"finish_time,omitempty"`
}

type ItemList []Item

type ItemVo struct {
	Title       string   `json:"title" bson:"title,omitempty" binding:"required"`
	Tags        []string `json:"tags" bson:"tags,omitempty"`
	Score       uint     `json:"score" bson:"score,omitempty"`
	Description string   `json:"description" bson:"description,omitempty"`
}
