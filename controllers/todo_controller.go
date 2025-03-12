package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/noman/todo-application/middleware"
	"github.com/noman/todo-application/models"
	"github.com/noman/todo-application/repository"
)

// TodoController handles todo requests
type TodoController struct {
	todoRepo *repository.TodoRepository
}

// NewTodoController creates a new TodoController
func NewTodoController() *TodoController {
	return &TodoController{
		todoRepo: repository.NewTodoRepository(),
	}
}

// Create handles creating a new todo
func (c *TodoController) Create(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the context
	userID, err := middleware.GetUserIDFromContext(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Parse the request body
	var req models.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate the request
	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	// Create the todo
	todo := &models.Todo{
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
		UserID:      userID,
	}

	if err := c.todoRepo.Create(todo); err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
		return
	}

	// Return the created todo
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo.ToResponse())
}

// GetAll handles getting all todos for a user
func (c *TodoController) GetAll(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the context
	userID, err := middleware.GetUserIDFromContext(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get all todos for the user
	todos, err := c.todoRepo.GetAllByUserID(userID)
	if err != nil {
		http.Error(w, "Failed to get todos", http.StatusInternalServerError)
		return
	}

	// Convert todos to responses
	responses := make([]models.TodoResponse, len(todos))
	for i, todo := range todos {
		responses[i] = todo.ToResponse()
	}

	// Return the todos
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

// GetByID handles getting a todo by ID
func (c *TodoController) GetByID(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the context
	userID, err := middleware.GetUserIDFromContext(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get the todo ID from the URL
	vars := mux.Vars(r)
	todoID, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	// Get the todo
	todo, err := c.todoRepo.GetByID(todoID)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Check if the todo belongs to the user
	if todo.UserID != userID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Return the todo
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo.ToResponse())
}

// Update handles updating a todo
func (c *TodoController) Update(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the context
	userID, err := middleware.GetUserIDFromContext(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get the todo ID from the URL
	vars := mux.Vars(r)
	todoID, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	// Get the existing todo
	todo, err := c.todoRepo.GetByID(todoID)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Check if the todo belongs to the user
	if todo.UserID != userID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Parse the request body
	var req models.UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Update the todo fields if provided
	if req.Title != "" {
		todo.Title = req.Title
	}
	if req.Description != "" {
		todo.Description = req.Description
	}
	if req.Completed != nil {
		todo.Completed = *req.Completed
	}

	// Update the todo in the database
	if err := c.todoRepo.Update(todo); err != nil {
		http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}

	// Return the updated todo
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo.ToResponse())
}

// Delete handles deleting a todo
func (c *TodoController) Delete(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the context
	userID, err := middleware.GetUserIDFromContext(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get the todo ID from the URL
	vars := mux.Vars(r)
	todoID, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	// Delete the todo
	if err := c.todoRepo.Delete(todoID, userID); err != nil {
		http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
		return
	}

	// Return success
	w.WriteHeader(http.StatusNoContent)
}