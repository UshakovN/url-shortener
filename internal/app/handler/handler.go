package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/UshakovN/url-shortener.git/internal/app/shortener"
	"github.com/UshakovN/url-shortener.git/internal/app/store"
	"github.com/UshakovN/url-shortener.git/internal/app/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	config *Config
	router *mux.Router
	store  store.Store
}

func NewHandler(config *Config) *Handler {
	return &Handler{
		config: config,
		router: mux.NewRouter(),
		store:  config.store,
	}
}

func (h *Handler) Start() error {
	h.configureRouter()
	log.Printf("server is started at port %s", h.config.port)
	return http.ListenAndServe(h.config.port, h.router)
}

func (h *Handler) configureRouter() {
	h.router.HandleFunc("/", h.handleStart())
	h.router.HandleFunc("/{url}", h.handleRedirect())
	subrouter := h.router.PathPrefix("/shortener").Subrouter()
	subrouter.HandleFunc("/short", h.handlePostShortUrl()).Methods(http.MethodPost)
	subrouter.HandleFunc("/source", h.handleGetSourceUrl()).Methods(http.MethodGet)
}

func (h *Handler) handleRedirect() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		shortUrl := mux.Vars(r)["url"]
		sourceUrl, ok := h.store.GetItem(shortUrl)
		if !ok {
			http.Error(rw, errors.New("url not found").Error(), http.StatusNotFound)
			return
		}
		http.Redirect(rw, r, sourceUrl, http.StatusMovedPermanently)
	}
}

func (h *Handler) handleStart() http.HandlerFunc {
	res := make(map[string]string)
	res["status"] = "active"
	jsonRes, _ := json.Marshal(res)
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Write(jsonRes)
	}
}

func (h *Handler) handleGetSourceUrl() http.HandlerFunc {
	type request struct {
		SourceUrl string `json:"source-url"`
	}
	return func(rw http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		shortUrl := r.FormValue("short-url")

		log.Printf("short-url: %s\n", shortUrl) // debug

		srcUrl, ok := h.store.GetItem(validator.TrimProtocol(shortUrl))
		if !ok {
			http.Error(rw, errors.New("url not found").Error(), http.StatusNotFound)
			return
		}
		jsonReq, err := json.Marshal(request{
			SourceUrl: srcUrl,
		})
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.Write(jsonReq)
	}
}

func (h *Handler) handlePostShortUrl() http.HandlerFunc {
	type request struct {
		ShortUrl string `json:"short-url"`
	}
	return func(rw http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		srcUrl := r.FormValue("source-url")

		log.Printf("source-url: %s\n", srcUrl) // debug

		if !validator.ValidateUrl(srcUrl) {
			http.Error(rw, errors.New("bad url").Error(), http.StatusBadRequest)
			return
		}

		shortUrl := shortener.GenerateShortUrl()
		h.store.PutItem(shortUrl, srcUrl)

		jsonRes, err := json.Marshal(request{
			ShortUrl: shortUrl,
		})
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.Write(jsonRes)
	}
}
