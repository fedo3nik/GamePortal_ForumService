package main

import (
	"context"
	"log"
	"net/http"

	"github.com/fedo3nik/GamePortal_ForumService/internal/application/service"
	"github.com/fedo3nik/GamePortal_ForumService/internal/config"
	"github.com/fedo3nik/GamePortal_ForumService/internal/interface/controller"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Panicf("Config create error: %v", err)
	}

	pool, err := pgxpool.Connect(context.Background(), c.PostgresURL)
	if err != nil {
		log.Panicf("Connect error: %v", err)
	}

	defer pool.Close()

	handler := mux.NewRouter()

	forumService := service.NewForumService(pool, "aKey", "rKey")
	addForumHandler := controller.NewHTTPAddForumHandler(forumService)
	getForumHandler := controller.NewHTTPGetForumHandler(forumService)

	handler.Handle("/forum/new-forum", addForumHandler).Methods("POST")
	handler.Handle("/forum/get-forum/{id}", getForumHandler).Methods("GET")

	err = http.ListenAndServe(c.Host+c.Port, handler)
	if err != nil {
		log.Panicf("Listen & Serve serror: %v", err)
	}
}
