package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

var pgInstance *Postgres
	
func NewPostgresStorage(databse_uri string) (*Postgres, error) {

	db,err:=pgxpool.New(context.Background(),databse_uri)
	if err!=nil {
		return nil,fmt.Errorf("unable to create connection pool: %w", err)
		
	}
	pgInstance = &Postgres{db}
	fmt.Println("Connected to DB...")
	return pgInstance,nil
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.Pool.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.Pool.Close()
}
