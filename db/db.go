package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func NewPostgresStorage(databse_uri string) (*pgx.Conn, error) {
	
	db, err := pgx.Connect(context.Background(), databse_uri)
	if err != nil {
		log.Println("Unable to connect to database: ",err)
		return nil,err
	}
	log.Println("Connected to database.")
	return db,err
}
