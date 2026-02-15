package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/St-Ivanov/sprint6/internal/server"
)

func main() {
	err := os.Chdir("..")
	if err != nil {
		log.Fatal(err)
	}

	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile(filepath.Join(curDir, "/log/info.log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
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
