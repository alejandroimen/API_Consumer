package controllers

import (
	"log"
	"net/http"
	"time"
	"strconv"

	"github.com/alejandroimen/API_Consumer/src/citas/application"
	"github.com/gin-gonic/gin"
)

type CreateCitasController struct {
	CreateCitas *application.CreateCitas
}

func NewCreateCitasController(CreateCitas *application.CreateCitas) *CreateCitasController {
	return &CreateCitasController{CreateCitas: CreateCitas}
}

func (c *CreateCitasController) Handle(ctx *gin.Context) {
    log.Println("Petición de crear un user, recibido")

    var request struct {
        IdUser string `json:"idUser"`
        Fecha  string `json:"fecha"`
        Estado string `json:"estado"`
    }

    if err := ctx.ShouldBindJSON(&request); err != nil {
        log.Printf("Error decodificando la petición del body: %v", err)
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "petición del body inválida"})
        return
    }

    // Convertir IdUser de string a int
    idUser, err := strconv.Atoi(request.IdUser)
    if err != nil {
        log.Printf("Error al convertir idUser a entero: %v", err)
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "idUser debe ser un número válido"})
        return
    }

    log.Printf("Creando user: user=%d, fecha=%s", idUser, request.Fecha)

    if err := c.CreateCitas.Run(idUser, request.Fecha, request.Estado); err != nil {
        log.Printf("Error creando el user: %v", err)
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    log.Printf("User creado exitosamente")
    ctx.JSON(http.StatusCreated, gin.H{"message": "user creado exitosamente"})
}

// Controlador para Short Polling
func (c *CreateCitasController) ShortPoll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
}

// Controlador para Long Polling
func (c *CreateCitasController) LongPoll(ctx *gin.Context) {
	timeout := time.After(30 * time.Second)
	select {
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
	case newData := <-waitForNewData():
		ctx.JSON(http.StatusOK, gin.H{"data": newData})
	}
}

func waitForNewData() <-chan string {
	newDataChannel := make(chan string)
	go func() {
		time.Sleep(10 * time.Second) // Simula el tiempo hasta que haya nuevos datos
		newDataChannel <- "Datos nuevos disponibles"
	}()
	return newDataChannel
}
