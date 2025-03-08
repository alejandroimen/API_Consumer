package application

import (
	"fmt"
	"github.com/alejandroimen/API_Consumer/src/citas/domain/repository"
)

type DeleteCitas struct {
	repo repository.CitasRepository
}

func NewDeleteCitas(repo repository.CitasRepository) *DeleteCitas {
	return &DeleteCitas{repo: repo}
}

func (du *DeleteCitas) Run(id int) error {
	_, err := du.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("ucitas no encontrado: %w", err)
	}

	if err := du.repo.Delete(id); err != nil {
		return fmt.Errorf("error eliminando el usuairo: %w", err)
	}
	return nil
}
