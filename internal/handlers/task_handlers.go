package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Harshitttttttt/go-todo/internal/database"
	"github.com/Harshitttttttt/go-todo/internal/util"
)

func CreateTask(repo *database.TaskRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Title string `json:"title"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			util.RespondWithError(w, http.StatusBadRequest, "Invalid Request Body")
			return
		}

		if req.Title == "" {
			util.RespondWithError(w, http.StatusBadRequest, "Title is required")
			return
		}

		task, err := repo.CreateTask(req.Title)
		if err != nil {
			log.Println(err)
			util.RespondWithError(w, http.StatusInternalServerError, "Failed to create task")
			return
		}

		util.RespondWithJSON(w, http.StatusCreated, task)
	}
}

func GetAllTasks(repo *database.TaskRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := repo.GetAllTasks()
		if err != nil {
			log.Println("Error getting all tasks: ", err)
			util.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch tasks")
			return
		}

		util.RespondWithJSON(w, http.StatusOK, tasks)
	}
}

func GetASingleTask(repo *database.TaskRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		task, err := repo.GetTaskById(id)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, "Error fetching task from DB")
		}

		if task == nil {
			util.RespondWithError(w, http.StatusNotFound, "Task not found")
		}

		util.RespondWithJSON(w, http.StatusOK, task)
	}
}

func DeleteATask(repo *database.TaskRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		deleted, err := repo.DeleteTaskById(id)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, "DB error while deleting task")
			return
		}

		if !deleted {
			util.RespondWithError(w, http.StatusNotFound, "Task not found")
			return
		}

		util.RespondWithJSON(w, http.StatusOK, map[string]string{
			"message": "Task deleted successfully",
		})
	}
}

func UpdateATask(repo *database.TaskRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		var req struct {
			Title string `json:"title"`
			Done  bool   `json:"done"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			util.RespondWithError(w, http.StatusBadRequest, "Invalid Request Body")
			return
		}

		updated, err := repo.UpdateTaskById(id, req.Title, req.Done)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, "Failed to update task")
			return
		}

		if !updated {
			util.RespondWithError(w, http.StatusNotFound, "Task Not Found")
			return
		}

		util.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Task Updated Successfully"})
	}
}
