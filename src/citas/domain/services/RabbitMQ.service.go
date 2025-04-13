package services

import (
	"github.com/alejandroimen/API_Consumer/src/citas/domain/entities"
	"github.com/alejandroimen/API_Consumer/src/citas/domain/repository"
)

// Interfaz para el servicio de RabbitMQ
type RabbitMQService interface {
    PublishCita(cita entities.Citas, colaDestino string) error
    ConsumeCreatedUsers(repo repository.CitasRepository)
    Close() error
}
