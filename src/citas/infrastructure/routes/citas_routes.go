package routes

import (
	"github.com/alejandroimen/API_Consumer/src/citas/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupucitasRoutes(r *gin.Engine, createucitasController *controllers.CreateucitasController, getucitasController *controllers.GetucitassController, deleteucitasController *controllers.DeleteucitasController, updateucitasController *controllers.UpdateucitasController) {
	// Rutas CRUD
	r.POST("/ucitass", createucitasController.Handle)
	r.GET("/ucitass", getucitasController.Handle)
	r.DELETE("/ucitass/:id", deleteucitasController.Handle)
	r.PUT("/ucitass/:id", updateucitasController.Handle)

	// Nuevas rutas para polling en POST
	r.POST("/ucitass/poll/short", createucitasController.ShortPoll)
	r.POST("/ucitass/poll/long", createucitasController.LongPoll)

	// Nuevas rutas para polling en DELETE
	r.DELETE("/ucitass/poll/short", deleteucitasController.ShortPoll)
	r.DELETE("/ucitass/poll/long", deleteucitasController.LongPoll)

	// Nuevas rutas para polling en PUT
	r.PUT("/ucitass/poll/short", updateucitasController.ShortPoll)
	r.PUT("/ucitass/poll/long", updateucitasController.LongPoll)

	// Nuevas rutas para polling en GET
	r.GET("/ucitass/poll/short", getucitasController.ShortPoll)
	r.GET("/ucitass/poll/long", getucitasController.LongPoll)
}
