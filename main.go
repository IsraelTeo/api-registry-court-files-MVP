package main

import (
	"fmt"
	"log"

	"github.com/IsraelTeo/api-registry-court-files-MVP/configuration"
	"github.com/IsraelTeo/api-registry-court-files-MVP/database"
	"github.com/IsraelTeo/api-registry-court-files-MVP/route"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar las variables de entorno
	godotenv.Load()

	// Inicializar la configuración cargando las variables de entorno
	cfg := configuration.InitConfig()

	// multiplexer para rutas

	// Conectar a la base de datos utilizando la configuración cargada
	err := database.Connection(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Database connection established successfully!")

	// Migración de entidades
	err = database.MigrateDB()
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	fmt.Println("Database migration successful")

	r := route.InitRoutes()

	// Wrap con CORS middleware
	handler := configuration.CORS(r)

	//Inicia servidor en el puerto: 8080
	err = configuration.StartServer(":8080", handler)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
}
