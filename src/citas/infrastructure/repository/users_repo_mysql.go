package repository

import (
	"database/sql"
	"fmt"
)

type ucitasRepoMySQL struct {
	db *sql.DB
}

func NewCreateucitasRepoMySQL(db *sql.DB) *ucitasRepoMySQL {
	return &ucitasRepoMySQL{db: db}
}

func (r *ucitasRepoMySQL) Save(ucitas entities.ucitas) error {
	query := "INSERT INTO ucitas (name, email, password) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, ucitas.Name, ucitas.Email, ucitas.Password)
	if err != nil {
		return fmt.Errorf("error insertando ucitas: %w", err)
	}
	return nil
}

func (r *ucitasRepoMySQL) FindByID(id int) (*entities.ucitas, error) {
	query := "SELECT id, name, email FROM ucitas WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var ucitas entities.ucitas
	if err := row.Scan(&ucitas.ID, &ucitas.Name, &ucitas.Email); err != nil {
		return nil, fmt.Errorf("error buscando el ucitas: %w", err)
	}
	return &ucitas, nil
}

func (r *ucitasRepoMySQL) FindAll() ([]entities.ucitas, error) {
	query := "SELECT id, name, email, password FROM ucitas"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error buscando los ucitas: %w", err)
	}
	defer rows.Close()

	var ucitas []entities.ucitas
	for rows.Next() {
		var ucitas entities.ucitas
		if err := rows.Scan(&ucitas.ID, &ucitas.Name, &ucitas.Email, &ucitas.Password); err != nil {
			return nil, err
		}
		ucitas = append(ucitas, ucitas)
	}
	return ucitas, nil
}

func (r *ucitasRepoMySQL) Update(ucitas entities.ucitas) error {
	query := "UPDATE ucitas SET name = ?, email = ?, password = ? WHERE id = ?"
	_, err := r.db.Exec(query, ucitas.Name, ucitas.Email, ucitas.Password, ucitas.ID)
	if err != nil {
		return fmt.Errorf("error actualizando ucitas: %w", err)
	}
	return nil
}

func (r *ucitasRepoMySQL) Delete(id int) error {
	query := "DELETE FROM ucitas WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error eliminando ucitas: %w", err)
	}
	return nil
}
