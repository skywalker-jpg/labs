package main

import (
	"Labs2/internal/config"
	"Labs2/internal/database"
	"Labs2/internal/logger"
	"Labs2/internal/server"
	"Labs2/internal/storage"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}

func run() error {
	config, err := config.NewConfig("config.yaml")
	if err != nil {
		return err
	}

	logger, err := logger.New(config.Logger)
	if err != nil {
		return err
	}

	db, err := database.Connection(config.DB)
	if err != nil {
		return err
	}
	defer db.Close()

	//err = database.MigrateUp(config.DB)
	//if err != nil {
	//	logger.Error("Error creating migrations", err.Error())
	//	return err
	//}

	st := storage.NewStorage(db)

	server, err := server.New(config.Server, logger, st)
	if err != nil {
		return err
	}

	return server.Serve()
}
