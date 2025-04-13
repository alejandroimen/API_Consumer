package application

import (
	"fmt"
	"github.com/alejandroimen/API_Consumer/src/citas/domain/repository"
)

type UpdateCitas struct {
	repo repository.CitasRepository
}

func NewUpdateCitas(repo repository.CitasRepository) *UpdateCitas {
	return &UpdateCitas{repo: repo}
}

func (us *UpdateCitas) Run(id int, idUser int, fecha string, estado string) error {
	ucitas, err := us.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("user no encontrado: %w", err)
	}

	//actualizo la cita
	ucitas.IdUser = idUser
	ucitas.Fecha = fecha
	ucitas.Estado = estado

	//guardo los cambios en el repositorio:
	if err := us.repo.Update(*ucitas); err != nil {
		return fmt.Errorf("error actualizando el user: %w", err)
	}

	return nil
}
