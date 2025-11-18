package database

import (
	"database/sql"

	"github.com/Harshitttttttt/go-todo/internal/models"
)

type TaskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

// Create Task
func (r *TaskRepository) CreateTask(title string) (*models.Task, error) {
	// Insert row
	_, err := r.DB.Exec(`
        INSERT INTO todos (title) VALUES (?);
    `, title)
	if err != nil {
		return nil, err
	}

	// Fetch last inserted row
	row := r.DB.QueryRow(`
        SELECT id, title, done, created_at 
        FROM todos 
        ORDER BY created_at DESC 
        LIMIT 1;
    `)

	task := &models.Task{}
	err = row.Scan(&task.Id, &task.Title, &task.Done, &task.CreatedAt)
	if err != nil {
		return nil, err
	}

	return task, err
}

// Get all tasks
func (r *TaskRepository) GetAllTasks() ([]*models.Task, error) {
	// Get all rows
	query := `
		SELECT id, title, done, created_at
		FROM todos;
	`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		var t models.Task
		err := rows.Scan(&t.Id, &t.Title, &t.Done, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// Get a task by it's id
func (r *TaskRepository) GetTaskById(id string) (*models.Task, error) {
	query := `
		SELECT id, title, done, created_at FROM todos WHERE id = (?);
	`

	row := r.DB.QueryRow(query, id)

	var task models.Task
	err := row.Scan(&task.Id, &task.Title, &task.Done, &task.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &task, nil
}

// Delete a Task
func (r *TaskRepository) DeleteTaskById(id string) (bool, error) {
	query := `
		DELETE FROM todos WHERE id = (?);
	`

	res, err := r.DB.Exec(query, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

// Update a task
func (r *TaskRepository) UpdateTaskById(id string, title string, done bool) (bool, error) {
	query := `
		UPDATE todos 
		SET title = ?, done = ?
		WHERE id = ?;
	`

	res, err := r.DB.Exec(query, title, done, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
