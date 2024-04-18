package main

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	todo_list "todo-list"
	"todo-list/internal/handler"
	"todo-list/internal/repository"
)

//connect database
//create repository
//create service
//create handler
//init dependency injection

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error with config file: %s", err.Error())
	}

	_, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("error with connecting to database: %s", err.Error())
	}
	handlers := new(handler.Handler)
	srv := new(todo_list.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil { // Change port to match your configuration
		log.Fatalf("error with running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
