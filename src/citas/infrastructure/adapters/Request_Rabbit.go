package adapters

import (
    "encoding/json"
    "fmt"
    "log"
    "math/rand"
    "time"

    "github.com/alejandroimen/API_Consumer/src/citas/domain/entities"
    "github.com/alejandroimen/API_Consumer/src/citas/domain/repository"
    "github.com/alejandroimen/API_Consumer/src/citas/domain/services"
    amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
    conn    *amqp.Connection
    channel *amqp.Channel
}

func NewRabbitMQAdapter(connectionString string) (services.RabbitMQService, error) {
    conn, err := amqp.Dial(connectionString)
    if err != nil {
        return nil, fmt.Errorf("error al conectar con RabbitMQ: %w", err)
    }

    channel, err := conn.Channel()
    if err != nil {
        return nil, fmt.Errorf("error al abrir un canal en RabbitMQ: %w", err)
    }

    return &RabbitMQAdapter{conn: conn, channel: channel}, nil
}

func (r *RabbitMQAdapter) PublishCita(cita entities.Citas, colaDestino string) error {
    orderJSON, err := json.Marshal(cita)
    if err != nil {
        return fmt.Errorf("error al convertir la cita a JSON: %w", err)
    }

    err = r.channel.Publish(
        "citas",       // exchange
        colaDestino,   // routing key
        false,
        false,
        amqp.Publishing{
            ContentType: "application/json",
            Body:        orderJSON,
        },
    )
    if err != nil {
        return fmt.Errorf("error al publicar el mensaje en RabbitMQ: %w", err)
    }

    return nil
}

func (r *RabbitMQAdapter) ConsumeCreatedUsers(repo repository.CitasRepository) {
    go func() {
        for {
            log.Println("🔄 Iniciando consumidor de 'created.user'...")

            // 1. Reintento de conexión si el canal está cerrado
            if r.channel == nil || r.channel.IsClosed() {
                log.Println("⚠️ Canal cerrado. Intentando reconectar...")
                var err error
                r.conn, err = amqp.Dial("amqp://rabbit:rabbit@35.170.173.77:5672/vh")
                if err != nil {
                    log.Printf("❌ Error al reconectar con RabbitMQ: %s", err)
                    time.Sleep(5 * time.Second)
                    continue
                }

                r.channel, err = r.conn.Channel()
                if err != nil {
                    log.Printf("❌ Error al abrir nuevo canal: %s", err)
                    time.Sleep(5 * time.Second)
                    continue
                }
            }

            // 2. Declarar cola y hacer binding
            queue, err := r.channel.QueueDeclare(
                "created.user", true, false, false, false, nil,
            )
            if err != nil {
                log.Printf("❌ Error al declarar la cola: %s", err)
                time.Sleep(5 * time.Second)
                continue
            }

            err = r.channel.QueueBind(
                "created.user", "created.user", "citas", false, nil,
            )
            if err != nil {
                log.Printf("❌ Error al hacer binding: %s", err)
                time.Sleep(5 * time.Second)
                continue
            }

            // 3. Crear canal para detectar cierre
            closeChan := make(chan *amqp.Error)
            r.channel.NotifyClose(closeChan)

            // 4. Consumir mensajes
            msgs, err := r.channel.Consume(
                queue.Name, "", true, false, false, false, nil,
            )
            if err != nil {
                log.Printf("❌ Error al consumir mensajes: %s", err)
                time.Sleep(5 * time.Second)
                continue
            }

            // 5. Procesar mensajes hasta que el canal se cierre
            for {
                select {
                case d, ok := <-msgs:
                    if !ok {
                        log.Println("⚠️ Canal de mensajes cerrado.")
                        time.Sleep(5 * time.Second)
                        break
                    }

                    var userId int
                    if err := json.Unmarshal(d.Body, &userId); err != nil {
                        log.Printf("❌ Error al decodificar el mensaje: %s", err)
                        continue
                    }

                    rand.Seed(time.Now().UnixNano())
                    estado := []string{"Pendiente", "Confirmada"}[rand.Intn(2)]

                    cita := entities.Citas{
                        IdUser: userId,
                        Fecha:  time.Now().Format("2006-01-02"),
                        Estado: estado,
                    }

                    if err := repo.Save(cita); err != nil {
                        log.Printf("❌ Error al guardar la cita: %s", err)
                        continue
                    }

                    var colaDestino string
                    if estado == "Confirmada" {
                        colaDestino = "citas.confirmadas"
                    } else {
                        colaDestino = "citas.pendientes"
                    }

                    if err := r.PublishCita(cita, colaDestino); err != nil {
                        log.Printf("❌ Error al publicar la cita en '%s': %s", colaDestino, err)
                    }

                    log.Printf("✅ Cita creada para el usuario %d con estado %s", userId, estado)

                case err := <-closeChan:
                    log.Printf("❌ Canal cerrado con error: %v", err)
                    time.Sleep(5 * time.Second)
                    break
                }
            }

            // Esperamos un momento antes de intentar reconectar
            time.Sleep(3 * time.Second)
        }
    }()
}

func (r *RabbitMQAdapter) Close() error {
    if r.channel != nil {
        r.channel.Close()
    }
    if r.conn != nil {
        r.conn.Close()
    }
    return nil
}
