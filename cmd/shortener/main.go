package main

import "github.com/UshakovN/url-shortener.git/internal/app/handler"

func main() {
	cfg := handler.NewConfig()
	hnd := handler.NewHandler(cfg)
	hnd.Start()
}
