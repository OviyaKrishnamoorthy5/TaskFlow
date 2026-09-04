package main

import (
	"taskflow-backend/config"
	"taskflow-backend/routes"
)

func main() {
	config.ConnectDB()

	r := routes.SetupRouter()
	r.Run(":8080")
}
