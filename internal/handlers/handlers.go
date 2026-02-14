package handlers

import (
	"bufio"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/St-Ivanov/sprint6/internal/service"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	curDir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file, err := os.ReadFile(filepath.Join(curDir, "/index.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.MultipartForm.RemoveAll()

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	formatFile := filepath.Ext(header.Filename)

	curDir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filePath := filepath.Join(curDir, "/data/", time.Now().UTC().Format("02.01.2006_15.04.05"), formatFile)

	fileNew, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer fileNew.Close()

	for scanner.Scan() {
		data := scanner.Text()
		dataParsed, err := service.DataConversion(data)
		if err != nil {
			continue
		}
		_, err = fileNew.WriteString(dataParsed + "\n")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(dataParsed + "\n"))
	}
}
