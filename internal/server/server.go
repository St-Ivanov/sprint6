package server

import (
	"log"
	"net/http"
	"time"

	"github.com/St-Ivanov/sprint6/internal/handlers"
)

type HttpServer struct {
	Loger *log.Logger
	Serv  *http.Server
}

func NewHttpServer(loger *log.Logger) *HttpServer {
	r := http.NewServeMux()

	r.HandleFunc("/", handlers.MainHandler)
	r.HandleFunc("/upload", handlers.UploadHandler)

	serv := HttpServer{
		Loger: loger, Serv: &http.Server{
			Addr:         ":8080",
			Handler:      r,
			ErrorLog:     loger,
			ReadTimeout:  time.Duration(5) * time.Second,
			WriteTimeout: time.Duration(10) * time.Second,
			IdleTimeout:  time.Duration(15) * time.Second,
		},
	}
	return &serv
}
