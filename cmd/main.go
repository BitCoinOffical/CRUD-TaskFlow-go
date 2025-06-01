package main

import (
	"log"
	"net/http"

	"main.go/internal/handlers"
	"main.go/internal/storage"
)

func main() {
	storage.InitSQLite()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetTaskHandlers(w, r)
		case http.MethodPost:
			handlers.CreateTaskHandlers(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPatch:
			handlers.UpdateTaskByIDHandlers(w, r)
		case http.MethodDelete:
			handlers.DeleteTaskHandlers(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
