package routes

import (
	"github.com/alejandroimen/API_Consumer/src/citas/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupucitasRoutes(r *gin.Engine, createCitasController *controllers.CreateCitasController, getcitasController *controllers.GetCitasController, deletecitasController *controllers.DeletecitasController, updateucitasController *controllers.UpdateCitasController) {
	// Rutas CRUD
	r.POST("/citas", createCitasController.Handle)
	r.GET("/citas", getcitasController.Handle)
	r.DELETE("/citas/:id", deletecitasController.Handle)
	r.PUT("/citas/:id", updateucitasController.Handle)
}
