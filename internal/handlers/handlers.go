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

type Handler struct {
	storage *storage.DataBase
}

func NewHandler(db *storage.DataBase) *Handler {
	return &Handler{storage: db}
}

func (h *Handler) CreateTaskHandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var task models.Tasks
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "not corect json", http.StatusBadRequest)
		return
	}
	err = h.storage.CreateTask(&task)
	if err != nil {
		log.Println("Ошибка при создании задачи:", err)
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *Handler) GetTaskHandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	tasks, err := h.storage.GetTask()
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

func (h *Handler) UpdateTaskByIDHandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
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

	err = h.storage.UpdateTaskByID(id, updateTask.Status)
	if err != nil {
		http.Error(w, "failed to update", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteTaskHandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	w.Header().Set("Content-Type", "application/json")
	IdTask := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, _ := strconv.Atoi(IdTask)
	h.storage.DeleteTask(id)
}
