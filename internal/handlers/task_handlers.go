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
