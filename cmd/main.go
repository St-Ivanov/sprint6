package main

import (
	"log"
	"os"

	"github.com/St-Ivanov/sprint6/internal/server"
)

func main() {
	file, err := os.OpenFile("log/info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	loger := log.New(file, "", 0)

	serv := server.NewHttpServer(loger)

	err = serv.Serv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
