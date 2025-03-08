package repository

import ( 	
	"github.com/alejandroimen/API_Consumer/src/citas/domain/entities"
)

type citasRepository interface {
	Save(citas entities.citas) error
	FindByID(id int) (*entities.citas, error)
	FindAll() ([]entities.citas, error)
	Update(citas entities.citas) error
	Delete(id int) error
}
