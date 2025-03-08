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

	// Nuevas rutas para polling en POST
	r.POST("/citas/poll/short", createCitasController.ShortPoll)
	r.POST("/citas/poll/long", createCitasController.LongPoll)

	// Nuevas rutas para polling en DELETE
	r.DELETE("/citas/poll/short", deletecitasController.ShortPoll)
	r.DELETE("/citas/poll/long", deletecitasController.LongPoll)

	// Nuevas rutas para polling en PUT
	r.PUT("/citas/poll/short", updateucitasController.ShortPoll)
	r.PUT("/citas/poll/long", updateucitasController.LongPoll)

	// Nuevas rutas para polling en GET
	r.GET("/citas/poll/short", getcitasController.ShortPoll)
	r.GET("/citas/poll/long", getcitasController.LongPoll)
}
