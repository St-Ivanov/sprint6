package server

import (
	"log"
	"net/http"

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
			Addr: "8080", Handler: r, ErrorLog: loger, ReadTimeout: 5, WriteTimeout: 10, IdleTimeout: 15,
		},
	}
	return &serv
}
