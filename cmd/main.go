package main

import (
	"log"

	"github.com/AnirudhBathala/ecom-api/cmd/api"
	"github.com/AnirudhBathala/ecom-api/config"
	"github.com/AnirudhBathala/ecom-api/db"
)

func main() {
	

	db, err := db.NewPostgresStorage(config.Configs.Database_URI)
	if err!=nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := api.NewAPIServer(":8080",db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
