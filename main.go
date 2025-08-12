package main

import (
	"main/config"
	"main/database"
	"main/routes"
)

func main() {
	config.Loadenv()
	database.ConnectDB()
	routes.SetUpRoutes()
}
