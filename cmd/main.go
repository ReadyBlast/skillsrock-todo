package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"skillsrock-todo/internal/config"
	"skillsrock-todo/internal/database"
	"skillsrock-todo/internal/handler"
	"skillsrock-todo/internal/repository"
	"skillsrock-todo/internal/service"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Get()
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.DbSslmode)

	dbPool, err := database.Connect(dbURL)
	if err != nil {
		log.Fatalf("[ERROR] Unable to connect to database: %v", err)
	}
	if err := dbPool.Ping(context.Background()); err != nil {
		log.Fatalf("[ERROR]: Unable to ping database %v", err)
	}
	defer dbPool.Close()

	log.Println("[INFO]: Connected to database")

	app := fiber.New()

	taskRepo := repository.NewTaskRepository(dbPool)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	api := app.Group("/tasks")
	api.Post("/", taskHandler.CreateTask)
	api.Get("/", taskHandler.GetTasks)
	api.Put("/:id", taskHandler.UpdateTask)
	api.Delete("/:id", taskHandler.DeleteTask)

	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("[ERROR]: Unable to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("[INFO]: Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Shutdown(); err != nil {
		log.Fatalf("[ERROR]: Unable to gracefully shutdown server: %v", err)
	}

	<-ctx.Done()
	log.Println("[INFO]: Server shutdown successfully")
}
