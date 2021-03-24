package main

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/grpc"

	"github.com/fedo3nik/GamePortal_ForumService/internal/application/service"
	"github.com/fedo3nik/GamePortal_ForumService/internal/config"
	grpcInfra "github.com/fedo3nik/GamePortal_ForumService/internal/infrastructure/grpc"
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

	grpcConn, err := grpc.Dial(c.Grpc, grpc.WithInsecure())
	if err != nil {
		log.Panicf("Grpc connection error: %v", err)
	}

	grpcClient := grpcInfra.NewSenderClient(grpcConn)

	emp := grpcInfra.Empty{}

	grpcResp, err := grpcClient.Send(context.Background(), &emp)
	if err != nil {
		log.Panicf("Grpc received error: %v", err)
	}

	handler := mux.NewRouter()

	forumService := service.NewForumService(pool, grpcResp.AccessPublicKey, grpcResp.RefreshPublicKey)
	addForumHandler := controller.NewHTTPAddForumHandler(forumService)
	getForumHandler := controller.NewHTTPGetForumHandler(forumService)

	handler.Handle("/forum/new-forum", addForumHandler).Methods("POST")
	handler.Handle("/forum/get-forum/{id}", getForumHandler).Methods("GET")

	go func() {
		err = http.ListenAndServe(c.Host+c.Port, handler)
		if err != nil {
			log.Panicf("Listen & Serve serror: %v", err)
		}
	}()

	select {}
}
