package server

import (
	"log"
	"net/http"
	"time"

	"github.com/St-Ivanov/sprint6/internal/handlers"
	"github.com/go-chi/chi/v5"
)

type HttpServer struct {
	Loger *log.Logger
	Serv  *http.Server
}

func NewHttpServer(loger *log.Logger) *HttpServer {
	r := chi.NewRouter()

	r.Get("/", handlers.MainHandler)
	r.Post("/upload", handlers.UploadHandler)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		loger.Fatal(err)
	}
	serv := HttpServer{
		Loger: loger, Serv: &http.Server{
			Addr: ":8070", Handler: r, ErrorLog: loger, ReadTimeout: time.Duration(5) * time.Second, WriteTimeout: time.Duration(10) * time.Second, IdleTimeout: time.Duration(15) * time.Second,
		},
	}
	return &serv
}
