package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alejandroimen/API_Consumer/src/citas/domain/entities"

)

type CitasRepository struct {
	db *sql.DB
}

func NewCitasRepository(db *sql.DB) *CitasRepository {
	return &CitasRepository{db: db}
}

func (r *CitasRepository) Save(cita entities.Citas) error {
	query := "INSERT INTO citas (idUser, fecha, estado) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, cita.IdUser, cita.Fecha, cita.Estado)
	if err != nil {
		return fmt.Errorf("error insertando citas: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener el ID del pedido: %w", err)
	}
	cita.ID = int(id)

	log.Printf("✅ Pedido guardado en la BD: %+v", cita)


	return nil
}

func (r *CitasRepository) FindAll() ([]entities.Citas, error) {
	query := "SELECT id, name, email FROM citas"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("error buscando los Users: %w", err)
	}
	defer rows.Close()

	var citas []entities.Citas
	for rows.Next() {
		var cita entities.Citas
		if err := rows.Scan(&cita.ID, &cita.IdUser, &cita.Fecha, &cita.Estado); err != nil {
			return nil, err
		}
		citas = append(citas, cita)
	}
	return citas, nil
}

func (r *CitasRepository) FindByID(id int) (*entities.Citas, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var cita entities.Citas
	if err := row.Scan(&cita.ID, cita.IdUser, cita.Fecha, cita.Estado); err != nil {
		return nil, fmt.Errorf("error buscando el User: %w", err)
	}
	return &cita, nil
}

func (r *CitasRepository) Update(citas entities.Citas) error {
	query := "UPDATE citas SET name = ?, email = ?, password = ? WHERE id = ?"
	_, err := r.db.Exec(query, citas.ID, citas.IdUser, citas.Fecha, citas.Estado)
	if err != nil {
		return fmt.Errorf("error actualizando citas: %w", err)
	}
	return nil
}

func (r *CitasRepository) Delete(id int) error {
	query := "DELETE FROM citas WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error eliminando citas: %w", err)
	}
	return nil
}