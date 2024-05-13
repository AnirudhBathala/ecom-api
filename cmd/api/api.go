package api

import (
	"log"
	"net/http"

	"github.com/AnirudhBathala/ecom-api/config"
	"github.com/AnirudhBathala/ecom-api/services/user"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type APIServer struct {
	addr string
	db   *pgx.Conn
	config config.Config
}

func (s *APIServer) Run() error {
	router := chi.NewRouter()

	router.Get("/test",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from server"))
	})

	router.Route("/api/v1/",func(r chi.Router) {
		userStore:=user.NewStore(s.db)
		userHandler:=user.NewHandler(userStore)
		userHandler.RigesterRoutes(r)
	})
	

	log.Println("Listening on PORT:",s.addr)
	return http.ListenAndServe(s.addr, router)
}

func NewAPIServer(addr string, db *pgx.Conn,config config.Config) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
		config: config,
	}
}
