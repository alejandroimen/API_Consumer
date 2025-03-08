package application

//"fmt"
import (
	
	"github.com/alejandroimen/API_Consumer/src/citas/domain/repository"
	"github.com/alejandroimen/API_Consumer/src/citas/domain/entities"
)

type GetCitas struct {
	repo repository.CitasRepository
}

func NewGetCitas(repo repository.CitasRepository) *GetCitas {
	return &GetCitas{repo: repo}
}

func (gu *GetCitas) Run() ([]entities.Citas, error) {
	ucitass, err := gu.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return ucitass, nil
}
