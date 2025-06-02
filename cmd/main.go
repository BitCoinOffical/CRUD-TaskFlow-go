package main

import (
	"log"
	"net/http"

	"main.go/internal/handlers"
	"main.go/internal/router"
	"main.go/internal/storage"
)

func main() {
	db := storage.InitSQLite("tasks.db")
	h := handlers.NewHandler(db)
	mux := http.NewServeMux()

	mux.Handle("/tasks", router.MethodRouter{
		"GET":  h.GetTaskHandlers,
		"POST": h.CreateTaskHandlers,
	})

	mux.Handle("/tasks/", router.MethodRouter{
		"PATCH":  h.UpdateTaskByIDHandlers,
		"DELETE": h.DeleteTaskHandlers,
	})

	mux.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
