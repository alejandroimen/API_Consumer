package application

import (
	"fmt"
	"log"
	"github.com/alejandroimen/API_Consumer/src/citas/domain/repository"
	"github.com/alejandroimen/API_Consumer/src/citas/domain/entities"
	"github.com/alejandroimen/API_Consumer/src/citas/domain/services"
)

// Contiene un campo de repo de tipo repository.user... siendo esto una inyección de dependencias
type CreateCitas struct {
	repo repository.CitasRepository
	rabbit services.RabbitMQService
}

// constructor de createCitas, que recibe un repositorio como parametro y lo asigna al campo repo. siendo configurable
func NewCreateCita(repo repository.CitasRepository, rab services.RabbitMQService) *CreateCitas {
	return &CreateCitas{repo: repo, rabbit: rab}
}

func (cu *CreateCitas) Run(idUser int, fecha string, estado string) error {
	cita := entities.Citas{IdUser: idUser, Fecha: fecha, Estado: estado}
	err := cu.repo.Save(cita); 
	if err != nil {
		return fmt.Errorf("error al guardar el user: %w", err)
	}

	err = cu.rabbit.PublishCita(cita, "ctraeted")
	if err != nil {
		return fmt.Errorf("error al publicar evento en la cola: %w", err)
	}

	log.Printf("✅ Evento 'cita.created' publicado para el pedido %d", cita.ID)
	return nil
}
