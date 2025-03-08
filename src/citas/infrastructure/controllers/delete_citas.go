package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/alejandroimen/API_Consumer/src/citas/application"
	"github.com/gin-gonic/gin"
)

type DeletecitasController struct {
	deletecitas *application.Deletecitas
}

func NewDeletecitasController(deletecitas *application.Deletecitas) *DeletecitasController {
	return &DeletecitasController{deletecitas: deletecitas}
}

func (du *DeletecitasController) Handle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID de user inválido"})
		return
	}

	if err := du.deletecitas.Run(id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "user eliminado correctamente"})
}

// Controlador para Short Polling
func (du *DeletecitasController) ShortPoll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
}

// Controlador para Long Polling
func (du *DeletecitasController) LongPoll(ctx *gin.Context) {
	timeout := time.After(30 * time.Second)
	select {
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
	case newData := <-waitForNewData():
		ctx.JSON(http.StatusOK, gin.H{"data": newData})
	}
}
