package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/St-Ivanov/sprint6/internal/service"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка при попытке загрузить файл", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var buf bytes.Buffer
	_, err = buf.ReadFrom(file)
	if err != nil {
		http.Error(w, "Ошибка при попытке загрузить файл", http.StatusInternalServerError)
		return
	}

	formatFile := filepath.Ext(header.Filename)

	filePath := time.Now().UTC().Format("02012006_150405") + formatFile

	fileNew, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Ошибка при попытке загрузить файл", http.StatusInternalServerError)
		return
	}
	defer fileNew.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	dataParsed := service.DataConversion(buf.String())
	_, err = fileNew.WriteString(dataParsed)
	if err != nil {
		http.Error(w, "Ошибка при попытке загрузить файл", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, dataParsed)
}
