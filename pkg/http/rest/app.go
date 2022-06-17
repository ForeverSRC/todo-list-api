package rest

import (
	itemcreating "github.com/ForeverSRC/todo-list-api/pkg/service/item/creating"
	itemdeleting "github.com/ForeverSRC/todo-list-api/pkg/service/item/deleting"
	itemediting "github.com/ForeverSRC/todo-list-api/pkg/service/item/editing"
	itemlisting "github.com/ForeverSRC/todo-list-api/pkg/service/item/listing"
	itemmanaging "github.com/ForeverSRC/todo-list-api/pkg/service/item/managing"
	missioncreating "github.com/ForeverSRC/todo-list-api/pkg/service/mission/creating"
	missiondetail "github.com/ForeverSRC/todo-list-api/pkg/service/mission/detail"
	missionitemadd "github.com/ForeverSRC/todo-list-api/pkg/service/mission/itemadd"
	missionlisting "github.com/ForeverSRC/todo-list-api/pkg/service/mission/listing"
	missionupdate "github.com/ForeverSRC/todo-list-api/pkg/service/mission/update"
)

type App struct {
	ItemService
	MissionService
}

type ItemService struct {
	ItemLister  itemlisting.Service
	ItemCreator itemcreating.Service
	ItemEditor  itemediting.Service
	ItemManager itemmanaging.Service
	ItemDeleter itemdeleting.Service
}

type MissionService struct {
	MissionCreator   missioncreating.Service
	MissionLister    missionlisting.Service
	MissionDetail    missiondetail.Service
	MissionUpdater   missionupdate.Service
	MissionItemAdder missionitemadd.Service
}
