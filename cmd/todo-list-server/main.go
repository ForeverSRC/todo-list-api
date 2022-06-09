package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ForeverSRC/todo-list-api/pkg/config"
	"github.com/ForeverSRC/todo-list-api/pkg/http/rest"
	itemcreating "github.com/ForeverSRC/todo-list-api/pkg/item/creating"
	itemlisting "github.com/ForeverSRC/todo-list-api/pkg/item/listing"
	itemmanaging "github.com/ForeverSRC/todo-list-api/pkg/item/managing"
	"github.com/ForeverSRC/todo-list-api/pkg/storage/mongodb"
)

func init() {
	config.Init()
}

func main() {

	var itemStore *mongodb.Storage
	var itemCreator itemcreating.Service
	var itemLister itemlisting.Service
	var itemStateManager itemmanaging.Service

	itemStore = mongodb.NewStorage(config.Config.GetString("mongo.url"), config.Config.GetString("mongo.db"))
	defer itemStore.Close()

	itemCreator = itemcreating.NewService(itemStore)
	itemLister = itemlisting.NewService(itemStore)
	itemStateManager = itemmanaging.NewService(itemStore)

	handler := rest.Handler(itemCreator, itemLister, itemStateManager)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.Config.Get("port")),
		Handler: handler,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("listen error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown error:%v", err)
	}
}
