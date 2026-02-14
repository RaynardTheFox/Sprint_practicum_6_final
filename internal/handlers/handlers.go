package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// IndexHandler обрабатывает запрос к корневому эндпоинту и возвращает HTML из файла index.html.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	htmlContent, err := os.ReadFile("index.html")
	if err != nil {
		log.Printf("Error reading index.html: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(htmlContent); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

// UploadHandler обрабатывает загрузку файла, конвертирует его содержимое и возвращает результат.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Printf("Error parsing multipart form: %v", err)
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		log.Printf("Error getting file from form: %v", err)
		http.Error(w, "Error getting file from form", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	input := string(fileData)
	converted, err := service.Convert(input)
	if err != nil {
		log.Printf("Error converting: %v", err)
		http.Error(w, "Error converting data", http.StatusInternalServerError)
		return
	}

	fileName := time.Now().UTC().String()
	ext := filepath.Ext(header.Filename)
	if ext == "" {
		ext = ".txt"
	}
	fileName = sanitizeFileName(fileName) + ext

	outputFile, err := os.Create(fileName)
	if err != nil {
		log.Printf("Error creating output file: %v", err)
		http.Error(w, "Error creating output file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(converted)
	if err != nil {
		log.Printf("Error writing to output file: %v", err)
		http.Error(w, "Error writing to output file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprint(w, converted); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func sanitizeFileName(name string) string {
	var result []rune
	for _, r := range name {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') ||
			r == '.' || r == '-' || r == '_' {
			result = append(result, r)
		} else {
			result = append(result, '_')
		}
	}
	return string(result)
}
