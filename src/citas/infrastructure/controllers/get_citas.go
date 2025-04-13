package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/alejandroimen/API_Consumer/src/citas/application"
	"github.com/gin-gonic/gin"
)

type GetCitasController struct {
	getCitas *application.GetCitas
}

func NewGetCitasController(getCitas *application.GetCitas) *GetCitasController {
	return &GetCitasController{getCitas: getCitas}
}

func (gu *GetCitasController) Handle(ctx *gin.Context) {
	log.Println("Petición de listar todos los citas, recibido")

	citas, err := gu.getCitas.Run()
	if err != nil {
		log.Printf("Error buscando citas")
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Retornando %d citas", len(citas))
	ctx.JSON(200, citas)
}
func (c *GetCitasController) ShortPoll(ctx *gin.Context) {
	// Obtener los productos (esto simula si hay cambios o no)
	products, err := c.getCitas.Run()
	if err != nil {
		//ctx.JSON(http.StatusInternalserverError, gin.H{"error": err.Error()})
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
func (gu *GetCitasController) LongPoll(ctx *gin.Context) {
	timeout := time.After(30 * time.Second)
	select {
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
	case newData := <-waitForNewData():
		ctx.JSON(http.StatusOK, gin.H{"data": newData})
	}
}
