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
	file, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Ошибка при попытке получить HTML страницу", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10)
	if err != nil {
		http.Error(w, "Ошибка при попытке загрузить файл 1", http.StatusInternalServerError)
		return
	}
	defer r.MultipartForm.RemoveAll()

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка при попытке загрузить файл 2", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	formatFile := filepath.Ext(header.Filename)

	filePath := "data/" + time.Now().UTC().Format("02012006_150405") + formatFile

	fileNew, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Ошибка при попытке загрузить файл 3", http.StatusInternalServerError)
		return
	}
	defer fileNew.Close()

	w.Header().Set("Content-Type", "text/html")

	for scanner.Scan() {
		data := scanner.Text()
		dataParsed, err := service.DataConversion(data)
		if err != nil {
			continue
		}
		_, err = fileNew.WriteString(dataParsed + "\n")
		if err != nil {
			http.Error(w, "Ошибка при попытке загрузить файл 4", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(dataParsed + "\n"))
	}
	w.WriteHeader(http.StatusOK)
}
