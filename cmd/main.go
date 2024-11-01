// cmd/main.go
package main

import (
	"fmt"
	"net/http"
	"splitwise-backend/auth"
	"splitwise-backend/config"
	"splitwise-backend/db"
	"splitwise-backend/handlers"
	"splitwise-backend/middleware"
	"splitwise-backend/repositories"
	"splitwise-backend/services"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	db.ConnectDB(cfg)
	defer db.DB.Close()

	// Initialize the router
	router := mux.NewRouter()

	// Initialize repositories and services
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	handlers.RegisterUserRoutes(router, userService)

	// Register Auth routes
	auth.RegisterAuthRoutes(router, userService)

	// Initialize group repository and service
	groupRepo := repositories.NewGroupRepository()
	groupService := services.NewGroupService(groupRepo)
	handlers.RegisterGroupRoutes(router, groupService)

	// Initialize expense repository and service
	expenseRepo := repositories.NewExpenseRepository()
	expenseService := services.NewExpenseService(expenseRepo)
	handlers.RegisterExpenseRoutes(router, expenseService)

	// Use the auth middleware to protect certain routes
	protectedRoutes := router.PathPrefix("/groups").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)
	handlers.RegisterGroupRoutes(protectedRoutes, groupService)

	protectedRoutes = router.PathPrefix("/expenses").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)
	handlers.RegisterExpenseRoutes(protectedRoutes, expenseService)

	// Log all registered routes
	err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		methods, err := route.GetMethods()
		if err != nil {
			return err
		}
		path, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Printf("Route: %s, Methods: %v\n", path, methods)
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking routes: %v\n", err)
	}

	// Start the server
	fmt.Println("Server is running on port 9000...")
	if err := http.ListenAndServe(":9000", router); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
