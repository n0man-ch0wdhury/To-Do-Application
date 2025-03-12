package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/noman/todo-application/database"
	"github.com/noman/todo-application/models"
)

// TodoRepository handles database operations for todos
type TodoRepository struct {
	db *sql.DB
}

// NewTodoRepository creates a new TodoRepository
func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		db: database.DB,
	}
}

// Create creates a new todo in the database
func (r *TodoRepository) Create(todo *models.Todo) error {
	// Set the ID and timestamps
	todo.ID = uuid.New()
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	// Insert the todo into the database
	query := `
	INSERT INTO todos (id, title, description, completed, user_id, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.Exec(query, todo.ID, todo.Title, todo.Description, todo.Completed, todo.UserID, todo.CreatedAt, todo.UpdatedAt)
	return err
}

// GetByID gets a todo by ID
func (r *TodoRepository) GetByID(id uuid.UUID) (*models.Todo, error) {
	query := `
	SELECT id, title, description, completed, user_id, created_at, updated_at
	FROM todos
	WHERE id = $1
	`

	row := r.db.QueryRow(query, id)

	todo := &models.Todo{}
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.UserID, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("todo not found")
		}
		return nil, err
	}

	return todo, nil
}

// GetAllByUserID gets all todos for a user
func (r *TodoRepository) GetAllByUserID(userID uuid.UUID) ([]*models.Todo, error) {
	query := `
	SELECT id, title, description, completed, user_id, created_at, updated_at
	FROM todos
	WHERE user_id = $1
	ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []*models.Todo{}
	for rows.Next() {
		todo := &models.Todo{}
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.UserID, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

// Update updates a todo in the database
func (r *TodoRepository) Update(todo *models.Todo) error {
	// Update the timestamp
	todo.UpdatedAt = time.Now()

	// Update the todo in the database
	query := `
	UPDATE todos
	SET title = $1, description = $2, completed = $3, updated_at = $4
	WHERE id = $5 AND user_id = $6
	`

	result, err := r.db.Exec(query, todo.Title, todo.Description, todo.Completed, todo.UpdatedAt, todo.ID, todo.UserID)
	if err != nil {
		return err
	}

	// Check if the todo was found
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("todo not found or not owned by user")
	}

	return nil
}

// Delete deletes a todo from the database
func (r *TodoRepository) Delete(id, userID uuid.UUID) error {
	query := `
	DELETE FROM todos
	WHERE id = $1 AND user_id = $2
	`

	result, err := r.db.Exec(query, id, userID)
	if err != nil {
		return err
	}

	// Check if the todo was found
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("todo not found or not owned by user")
	}

	return nil
}