package main

import (
	"main/config"
	"main/database"
	_ "main/docs"
	"main/routes"
	"main/storage"
)

// @title           Streaming Backend API
// @version         1.0
// @description     API documentation for the streaming live chat backend.
// @contact.name    Jackson API Support
// @contact.email   Jacksontpa7@gmail.com
// @license.name    MIT
// @BasePath        /api/v1
func main() {
	config.Loadenv()
	storage.Init()
	database.ConnectDB()
	routes.SetUpRoutes()
}
