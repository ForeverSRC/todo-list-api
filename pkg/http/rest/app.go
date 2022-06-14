package rest

import (
	itemcreating "github.com/ForeverSRC/todo-list-api/pkg/service/item/creating"
	itemdeleting "github.com/ForeverSRC/todo-list-api/pkg/service/item/deleting"
	itemediting "github.com/ForeverSRC/todo-list-api/pkg/service/item/editing"
	itemlisting "github.com/ForeverSRC/todo-list-api/pkg/service/item/listing"
	itemmanaging "github.com/ForeverSRC/todo-list-api/pkg/service/item/managing"
)

type App struct {
	ItemLister  itemlisting.Service
	ItemCreator itemcreating.Service
	ItemEditor  itemediting.Service
	ItemManager itemmanaging.Service
	ItemDeleter itemdeleting.Service
}
