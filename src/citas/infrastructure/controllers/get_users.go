package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/alejandroimen/API_Consumer/src/ucitas/application"
	"github.com/gin-gonic/gin"
)

type GetucitassController struct {
	getucitass *application.Getucitass
}

func NewucitassController(getucitass *application.Getucitass) *GetucitassController {
	return &GetucitassController{getucitass: getucitass}
}

func (gu *GetucitassController) Handle(ctx *gin.Context) {
	log.Println("Petición de listar todos los ucitas, recibido")

	ucitas, err := gu.getucitass.Run()
	if err != nil {
		log.Printf("Error buscando ucitas")
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Retornando %d ucitas", len(ucitas))
	ctx.JSON(200, ucitas)
}
func (c *GetucitassController) ShortPoll(ctx *gin.Context) {
	// Obtener los productos (esto simula si hay cambios o no)
	products, err := c.getucitass.Run()
	if err != nil {
		ctx.JSON(http.StatusInternalserverError, gin.H{"error": err.Error()})
		return
	}

	if len(products) == 0 {
		// No hay productos (o cambios)
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
		return
	}

	// Devolver productos (o cambios detectados)
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Datos actualizados",
		"products": products,
	})
}

// Controlador para Long Polling
func (gu *GetucitassController) LongPoll(ctx *gin.Context) {
	timeout := time.After(30 * time.Second)
	select {
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
	case newData := <-waitForNewData():
		ctx.JSON(http.StatusOK, gin.H{"data": newData})
	}
}
