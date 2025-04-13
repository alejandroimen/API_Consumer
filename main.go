package main

import (
	"log"

	"github.com/alejandroimen/API_Consumer/src/core"
	citasInfra "github.com/alejandroimen/API_Consumer/src/citas/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conexión a MySQL
	db, err := core.NewMySQLConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Configuración del enrutador de Gin
	r := gin.Default()
	r.Use(core.SetupCORS())

	citasInfra.InitCitasDependencies(r, db)


	// Iniciar servidor
	log.Println("server started at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
