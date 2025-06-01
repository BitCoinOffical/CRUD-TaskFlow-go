package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"main.go/internal/models"
	"main.go/internal/storage"
)

func CreateTaskHandlers(w http.ResponseWriter, r *http.Request) {
	var task models.Tasks
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "not corect json", http.StatusBadRequest)
		return
	}
	err = storage.CreateTask(&task)
	if err != nil {
		log.Println("Ошибка при создании задачи:", err)
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetTaskHandlers(w http.ResponseWriter, r *http.Request) {
	tasks, err := storage.GetTask()
	if err != nil {
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}
	if tasks == nil {
		tasks = []models.Tasks{}
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func UpdateTaskByIDHandlers(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var updateTask struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&updateTask); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	err = storage.UpdateTaskByID(id, updateTask.Status)
	if err != nil {
		http.Error(w, "failed to update", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteTaskHandlers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	IdTask := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, _ := strconv.Atoi(IdTask)
	storage.DeleteTask(id)
}
