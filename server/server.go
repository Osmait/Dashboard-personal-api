package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/osmait/admin-finanzas/database"
	"github.com/osmait/admin-finanzas/repository"
	"github.com/rs/cors"
)

type Config struct {
	Port        string
	JWTSecret   string
	DataBaseUrl string
}

type Server interface {
	Config() *Config
}
type Broke struct {
	config *Config
	router *mux.Router
}

func (b *Broke) Config() *Config {
	return b.config
}
func NewServer(ctx context.Context, config *Config) (*Broke, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("secret is required")
	}
	if config.DataBaseUrl == "" {
		return nil, errors.New("url is required")
	}
	broker := &Broke{
		config: config,
		router: mux.NewRouter(),
	}
	return broker, nil
}

func (b *Broke) Strat(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()

	binder(b, b.router)
	handler := cors.AllowAll().Handler(b.router)
	repo, err := database.NewPostgresRepository(b.config.DataBaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)
	log.Println("Starting server on port", b.Config().Port)
	if err := http.ListenAndServe(b.config.Port, handler); err != nil {
		log.Fatal("ListenAndServer:", err)
	}

}
