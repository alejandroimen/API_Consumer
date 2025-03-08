package main

import (
	"log"

	"github.com/alejandroimen/API_Consumer/src/core"
	citasApp "github.com/alejandroimen/API_Consumer/src/citas/application"
	citasController "github.com/alejandroimen/API_Consumer/src/citas/infrastructure/controllers"
	citasRepo "github.com/alejandroimen/API_Consumer/src/citas/infrastructure/repository"
	citasRoutes "github.com/alejandroimen/API_Consumer/src/citas/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conexión a MySQL
	db, err := core.NewMySQLConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Repositorios
	citasRepository := citasRepo.NewCreatecitasRepoMySQL(db)

	// Casos de uso para citas
	createcitas := citasApp.NewCreateCitas(citasRepository)
	getcitass := citasApp.NewGetcitas(citasRepository)
	deletecitass := citasApp.NewDeletecitas(citasRepository)
	updatecitass := citasApp.NewUpdatecitas(citasRepository)

	// Controladores para citas
	createcitasController := citasController.NewCreatecitasController(createcitas)
	getcitasController := citasController.NewcitassController(getcitass)
	deletecitasController := citasController.NewDeletecitasController(deletecitass)
	updatecitasController := citasController.NewUpdatecitasController(updatecitass)

	// Configuración del enrutador de Gin
	r := gin.Default()

	// Configurar rutas de citas
	citasRoutes.SetupcitasRoutes(r, createcitasController, getcitasController, deletecitasController, updatecitasController)

	// Iniciar servidor
	log.Println("server started at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
