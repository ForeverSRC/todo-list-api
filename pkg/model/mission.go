package model

import (
	"fmt"
	"time"
)

type MissionPriority uint

const (
	P1 MissionPriority = 1
	P2 MissionPriority = 2
	P3 MissionPriority = 3

	P99 MissionPriority = 99
)

func (p MissionPriority) String() string {
	return fmt.Sprintf("P%d", p)
}

type MissionState uint

const (
	MissionStateInit MissionState = iota
	MissionStateProgressing
	MissionStateFinished
)

func (ms MissionState) Pointer() *MissionState {
	return &ms
}

type Mission struct {
	Id        string `json:"id" bson:"_id,omitempty"`
	MissionVo `json:",inline" bson:",inline"`

	State *MissionState `json:"state" bson:"state,omitempty"`

	CreateTime *time.Time `json:"createTime" bson:"create_time,omitempty"`
	UpdateTime *time.Time `json:"updateTime" bson:"update_time,omitempty"`

	ItemDetails   map[string]Item `json:"itemDetails" bson:"-"`
	TotalItems    int             `json:"totalItems" bson:"-"`
	FinishedItems int             `json:"finishedItems" bson:"-"`
}

type MissionVo struct {
	Title    string          `json:"title" bson:"title,omitempty"`
	Priority MissionPriority `json:"priority" bson:"priority,omitempty"`
	Detail   *string         `json:"detail" bson:"detail,omitempty"`
	Items    []string        `json:"items" bson:"items,omitempty"`
}

type MissionList struct {
	Missions []Mission `json:"missions"`
	NoMore   bool      `json:"noMore"`
}
