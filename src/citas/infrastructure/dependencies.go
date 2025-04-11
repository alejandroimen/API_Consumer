package infraestructure

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	citasApp "github.com/alejandroimen/API_Consumer/src/citas/application"
	citasController "github.com/alejandroimen/API_Consumer/src/citas/infraestructure/controllers"
    citasRepo "github.com/alejandroimen/API_Consumer/src/citas/infraestructure/repository"
    citasRoutes "github.com/alejandroimen/API_Consumer/src/citas/infraestructure/routes"
	"github.com/alejandroimen/API_Consumer/src/citas/infraestructure/adapters"
)

func InitCitasDependencies(Engine *gin.Engine, db *sql.DB){
	adapters.InitRabbitMQ()

	citasRepository := citasRepo.NewNotificationRepositoryMySQL(db)

	go adapters.ConsumeCreatedOrders(citasRepository)
	createCita := citasApp.NewCreateNotification(citasRepository)
	getByUserNoti := citasApp.NewGetNotificationsByUser(citasRepository)

	createCitasController := citasController.NewCreateNotificationController(createCita)
	getByUserCitasController := citasController.NewGetNotificationsByUserController(getByUserNoti)

	citasRoutes.NotificationRoutes(Engine, createCitasController, getByUserCitasController)

}