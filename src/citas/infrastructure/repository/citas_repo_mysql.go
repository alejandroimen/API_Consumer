package repository

import (
	"database/sql"
	"fmt"
	"github.com/alejandroimen/API_Consumer/src/citas/domain/entities"
)

type citasRepoMySQL struct {
	db *sql.DB
}

func NewCreatecitasRepoMySQL(db *sql.DB) *citasRepoMySQL {
	return &citasRepoMySQL{db: db}
}

func (r *citasRepoMySQL) Save(citas entities.Citas) error {
	query := "INSERT INTO citas (name, email, password) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, citas.Name, citas.Email, citas.Password)
	if err != nil {
		return fmt.Errorf("error insertando citas: %w", err)
	}
	return nil
}

func (r *citasRepoMySQL) FindByID(id int) (*entities.Citas, error) {
	query := "SELECT id, name, email FROM citas WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var citas entities.Citas
	if err := row.Scan(&citas.ID, &citas.Name, &citas.Email); err != nil {
		return nil, fmt.Errorf("error buscando el citas: %w", err)
	}
	return &citas, nil
}

func (r *citasRepoMySQL) FindAll() ([]entities.Citas, error) {
	query := "SELECT id, name, email, password FROM citas"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error buscando los citas: %w", err)
	}
	defer rows.Close()

	var citas []entities.Citas
	for rows.Next() {
		var cita entities.Citas
		if err := rows.Scan(&cita.ID, &cita.Name, &cita.Email, &cita.Password); err != nil {
			return nil, err
		}
		citas = append(citas, cita)
	}
	return citas, nil
}

func (r *citasRepoMySQL) Update(citas entities.Citas) error {
	query := "UPDATE citas SET name = ?, email = ?, password = ? WHERE id = ?"
	_, err := r.db.Exec(query, citas.Name, citas.Email, citas.Password, citas.ID)
	if err != nil {
		return fmt.Errorf("error actualizando citas: %w", err)
	}
	return nil
}

func (r *citasRepoMySQL) Delete(id int) error {
	query := "DELETE FROM citas WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error eliminando citas: %w", err)
	}
	return nil
}
