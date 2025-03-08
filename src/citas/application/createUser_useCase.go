package application

import (
	"fmt"

	"github.com/alejandroimen/API_Consumer/src/citas/domain/repository"
	"github.com/alejandroimen/API_Consumer/src/citas/domain/entities"
)

// Contiene un campo de repo de tipo repository.user... siendo esto una inyección de dependencias
type CreateCitas struct {
	repo repository.citasRepository
}

// constructor de createCitas, que recibe un repositorio como parametro y lo asigna al campo repo. siendo configurable
func NewCreateUser(repo repository.citasRepository) *CreateCitas {
	return &CreateCitas{repo: repo}
}

func (cu *CreateCitas) Run(name string, email string, password string) error {
	user := entities.User{Name: name, Email: email, Password: password}
	if err := cu.repo.Save(user); err != nil {
		return fmt.Errorf("error al guardar el user: %w", err)
	}
	return nil
}
