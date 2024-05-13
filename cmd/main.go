package main

import (
	"context"
	"log"

	"github.com/AnirudhBathala/ecom-api/cmd/api"
	"github.com/AnirudhBathala/ecom-api/config"
	"github.com/AnirudhBathala/ecom-api/db"
)

func main() {
	
	config:=config.InitConfig()

	db, err := db.NewPostgresStorage(config.Database_URI)
	if err!=nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	server := api.NewAPIServer(":8080", db,config)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
