package repository

import (
	"github.com/alejandroimen/API_Consumer/src/citas/domain/entities"
)

type CitasRepository interface {
	Save(citas entities.Citas) error
	FindAll() ([]entities.Citas, error)
	FindByID(id int) (*entities.Citas, error)
	Update(citas entities.Citas) error
	Delete(id int) error
}
