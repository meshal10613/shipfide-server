package main

import (
	"fmt"

	"server/internal/config"
	"server/internal/database"
	"server/internal/server"
)

func main() {
	//? Load environment variables
	cfg, err := config.LoadEnv()
	if err != nil {
		panic(fmt.Sprintf("failed to load environment variables: %v", err))
	}

	//? Connect to the database
	db := database.ConnectDB(cfg)

	//? Start the server
	server.StartServer(db, cfg)
}
