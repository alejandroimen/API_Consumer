package repository

import (
	"database/sql"
	"fmt"
	"log"
	"encoding/json"

	"github.com/alejandroimen/API_Consumer/src/citas/domain/entities"
	"github.com/streadway/amqp"

)

type citasRepoMySQL struct {
	db *sql.DB
	channel *amqp.Channel
}

func NewCreatecitasRepoMySQL(db *sql.DB) *citasRepoMySQL {
	return &citasRepoMySQL{db: db}
}

func (r *citasRepoMySQL) Save(cita entities.Citas) error {
	query := "INSERT INTO citas (name, email, password) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, cita.Name, cita.Email, cita.Password)
	if err != nil {
		return fmt.Errorf("error insertando citas: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener el ID del pedido: %w", err)
	}
	cita.ID = int(id)

	log.Printf("✅ Pedido guardado en la BD: %+v", cita)

	err = r.PublishOrderCreated(cita)
	if err != nil {
		return fmt.Errorf("error al publicar evento en la cola: %w", err)
	}

	log.Printf("✅ Evento 'cita.created' publicado para el pedido %d", cita.ID)
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

func (r *citasRepoMySQL) PublishOrderCreated(cita entities.Citas) error {
	citaJSON, err := json.Marshal(cita)
	if err != nil {
		return fmt.Errorf("error al convertir la orden a JSON: %w", err)
	}

	err = r.channel.Publish(
		"",               // exchange
		"Citas",  // queue name
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        citaJSON,
		},
	)
	if err != nil {
		return fmt.Errorf("error al publicar el mensaje en RabbitMQ: %w", err)
	}

	log.Printf("✅ Evento 'cita.created' publicado para el pedido %d", cita.ID)
	return nil
}