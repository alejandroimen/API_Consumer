package controllers

import (
	"net/http"
	"strconv"
	"time"
	"log"

	"github.com/alejandroimen/API_Consumer/src/citas/application"
	"github.com/gin-gonic/gin"
)

type UpdateCitasController struct {
	updateCitas *application.UpdateCitas
}

func NewUpdateCitasController(updateCitas *application.UpdateCitas) *UpdateCitasController {
	return &UpdateCitasController{updateCitas: updateCitas}
}

func (update *UpdateCitasController) Handle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID de user inválido"})
		return
	}

	var request struct {
        IdUser string `json:"idUser"`
        Fecha  string `json:"fecha"`
        Estado string `json:"estado"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "petición del body inválida"})
		return
	}

	// Convertir IdUser de string a int
    idUser, err := strconv.Atoi(request.IdUser)
    if err != nil {
        log.Printf("Error al convertir idUser a entero: %v", err)
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "idUser debe ser un número válido"})
        return
    }

	if err := update.updateCitas.Run(id, idUser, request.Fecha, request.Estado); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "user actualizado correctamente"})
}

// Controlador para Short Polling
func (update *UpdateCitasController) ShortPoll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
}

// Controlador para Long Polling
func (update *UpdateCitasController) LongPoll(ctx *gin.Context) {
	timeout := time.After(30 * time.Second)
	select {
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
	case newData := <-waitForNewData():
		ctx.JSON(http.StatusOK, gin.H{"data": newData})
	}
}
