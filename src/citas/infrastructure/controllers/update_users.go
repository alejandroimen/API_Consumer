package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/alejandroimen/API_Consumer/src/ucitas/application"
	"github.com/gin-gonic/gin"
)

type UpdateucitasController struct {
	updateucitas *application.Updateucitas
}

func NewUpdateucitasController(updateucitas *application.Updateucitas) *UpdateucitasController {
	return &UpdateucitasController{updateucitas: updateucitas}
}

func (update *UpdateucitasController) Handle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID de user inválido"})
		return
	}

	var request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "petición del body inválida"})
		return
	}

	if err := update.updateucitas.Run(id, request.Email, request.Name, request.Password); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "user actualizado correctamente"})
}

// Controlador para Short Polling
func (update *UpdateucitasController) ShortPoll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
}

// Controlador para Long Polling
func (update *UpdateucitasController) LongPoll(ctx *gin.Context) {
	timeout := time.After(30 * time.Second)
	select {
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
	case newData := <-waitForNewData():
		ctx.JSON(http.StatusOK, gin.H{"data": newData})
	}
}
