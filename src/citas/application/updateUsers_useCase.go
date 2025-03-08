package application

import (
	"fmt"
	"github.com/alejandroimen/API_Consumer/src/citas/domain/repository"
)

type UpdateUcitas struct {
	repo repository.ucitasRepository
}

func NewUpdateUcitas(repo repository.ucitasRepository) *UpdateUcitas {
	return &UpdateUcitas{repo: repo}
}

func (us *UpdateUcitas) Run(id int, name string, email string, password string) error {
	ucitas, err := us.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("user no encontrado: %w", err)
	}

	//actualizo los campos del user:
	ucitas.Name = name
	ucitas.Email = email
	ucitas.Password = password

	//guardo los cambios en el repositorio:
	if err := us.repo.Update(*ucitas); err != nil {
		return fmt.Errorf("error actualizando el user: %w", err)
	}

	return nil
}
