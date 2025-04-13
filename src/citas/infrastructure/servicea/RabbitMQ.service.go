package infrastructure

import (
    "log"

    "github.com/alejandroimen/API_Consumer/src/citas/domain/services"
    "github.com/alejandroimen/API_Consumer/src/citas/infrastructure/adapters"
)

func InitRabbitMQService(connectionString string) services.RabbitMQService {
    rabbitMQService, err := adapters.NewRabbitMQAdapter(connectionString)
    if err != nil {
        log.Fatalf("Error inicializando RabbitMQ: %s", err)
    }
    return rabbitMQService
}