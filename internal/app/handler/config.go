package handler

import (
	"github.com/UshakovN/url-shortener.git/internal/app/db"
	"github.com/UshakovN/url-shortener.git/internal/app/store"
)

type Config struct {
	port  string
	store store.Store
}

func NewConfig() *Config {
	return &Config{
		port:  ":8080",
		store: db.NewStore(db.NewConfig()),
	}
}
