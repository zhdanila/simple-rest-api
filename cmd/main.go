package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	todo_list "todo-list"
	"todo-list/internal/handler"
	"todo-list/internal/repository"
	"todo-list/internal/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error with config file: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("error with connecting to database: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)
	srv := new(todo_list.Server)

	go func() {
		if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
			log.Fatalf("error with running server: %s", err.Error())
		}
	}()

	fmt.Println("TodoApp started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	fmt.Println("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error with shutting down server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Fatalf("error with db down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error with .env file, %s", err.Error())
	}

	return viper.ReadInConfig()
}
