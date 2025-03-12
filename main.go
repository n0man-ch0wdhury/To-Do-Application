package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/noman/todo-application/controllers"
	"github.com/noman/todo-application/database"
	"github.com/noman/todo-application/middleware"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Initialize database
	database.InitDB()

	// Initialize controllers
	authController := controllers.NewAuthController()
	todoController := controllers.NewTodoController()

	// Initialize router
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/api/auth/register", authController.Register).Methods("POST")
	router.HandleFunc("/api/auth/login", authController.Login).Methods("POST")
	
	// Protected auth routes
	authRouter := router.PathPrefix("/api/auth").Subrouter()
	authRouter.Use(middleware.AuthMiddleware)
	authRouter.HandleFunc("/logout", authController.Logout).Methods("POST")

	// Protected routes
	todoRouter := router.PathPrefix("/api/todos").Subrouter()
	todoRouter.Use(middleware.AuthMiddleware)
	todoRouter.HandleFunc("", todoController.Create).Methods("POST")
	todoRouter.HandleFunc("", todoController.GetAll).Methods("GET")
	todoRouter.HandleFunc("/{id}", todoController.GetByID).Methods("GET")
	todoRouter.HandleFunc("/{id}", todoController.Update).Methods("PUT")
	todoRouter.HandleFunc("/{id}", todoController.Delete).Methods("DELETE")

	// Get server port from environment variable
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	// Initialize server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Start server
	log.Printf("Server is running on port %s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
