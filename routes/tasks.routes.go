package routes

import (
	"encoding/json"
	"main/db"
	"main/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetTasksByIdHandler(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	var tasks []models.Task
	db.DB.Where("user_id = ?", userId).Find(&tasks)

	response := map[string]interface{}{
		"tasks": tasks,
	}

	json.NewEncoder(w).Encode(response)
}
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	createdTask := db.DB.Create(&task)
	err := createdTask.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(task)
}
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	taskId := mux.Vars(r)["id"]

	db.DB.First(&task, taskId)

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	json.NewEncoder(w).Encode(task)
}
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	taskId := mux.Vars(r)["id"]

	db.DB.First(&task, taskId)

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	} else {
		db.DB.Delete(&task)
	}
	w.WriteHeader(http.StatusOK)
}
