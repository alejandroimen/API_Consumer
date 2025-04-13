package infrastructure

import (
    "database/sql"
    "log"

    "github.com/alejandroimen/API_Consumer/src/citas/application"
    "github.com/alejandroimen/API_Consumer/src/citas/infrastructure/adapters"
    "github.com/alejandroimen/API_Consumer/src/citas/infrastructure/controllers"
    "github.com/alejandroimen/API_Consumer/src/citas/infrastructure/repository"
    "github.com/alejandroimen/API_Consumer/src/citas/infrastructure/routes"
    "github.com/gin-gonic/gin"
)

func InitCitasDependencies(Engine *gin.Engine, db *sql.DB) {
    // Inicializar RabbitMQ
    rabbitMQService, err := adapters.NewRabbitMQAdapter("amqp://rabbit:rabbit@35.170.173.77:5672/vh")
    if err != nil {
        log.Fatalf("Error inicializando RabbitMQ: %s", err)
    }

    // Inicializar el repositorio de citas
    citasRepository := repository.NewCitasRepository(db)

    // Configurar el consumo de mensajes de RabbitMQ
    go rabbitMQService.ConsumeCreatedUsers(citasRepository)

    // Inicializar los casos de uso
    createCita := application.NewCreateCita(citasRepository, rabbitMQService)               // Caso de uso para crear citas
    getCita := application.NewGetCitas(citasRepository)                   // Caso de uso para obtener citas
    updateCita := application.NewUpdateCitas(citasRepository)             // Caso de uso para actualizar citas
    deleteCita := application.NewDeleteCitas(citasRepository)             // Caso de uso para eliminar citas

    // Inicializar los controladores
    createCitasController := controllers.NewCreateCitasController(createCita)
    getCitaController := controllers.NewGetCitasController(getCita)
    updateCitaController := controllers.NewUpdateCitasController(updateCita)
    deleteCitaController := controllers.NewDeleteCitasController(deleteCita)

    // Configurar las rutas en Gin
    routes.SetupucitasRoutes(Engine, createCitasController, getCitaController, deleteCitaController, updateCitaController)

}
