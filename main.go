package main

import (
	"log"

	"github.com/alejandroimen/API_Consumer/src/core"
	citasApp "github.com/alejandroimen/API_Consumer/src/citas/application"
	citasController "github.com/alejandroimen/API_Consumer/src/citas/infrastructure/controllers"
	citasRepo "github.com/alejandroimen/API_Consumer/src/citas/infrastructure/repository"
	citasRoutes "github.com/alejandroimen/API_Consumer/src/citas/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conexión a MySQL
	db, err := core.NewMySQLConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Repositorios
	productRepository := productRepo.NewProductRepoMySQL(db)
	citasRepository := citasRepo.NewCreatecitasRepoMySQL(db)

	// Casos de uso para productos
	createProduct := productApp.NewCreateProduct(productRepository)
	getProducts := productApp.NewGetProducts(productRepository)
	updateProduct := productApp.NewUpdateProduct(productRepository)
	deleteProduct := productApp.NewDeleteProduct(productRepository)

	// Casos de uso para citas
	createcitas := citasApp.NewCreatecitas(citasRepository)
	getcitass := citasApp.NewGetcitass(citasRepository)
	deletecitass := citasApp.NewDeletecitas(citasRepository)
	updatecitass := citasApp.NewUpdatecitas(citasRepository)

	// Controladores para productos
	createProductController := productController.NewCreateProductController(createProduct)
	getProductsController := productController.NewGetProductsController(getProducts)
	updateProductController := productController.NewUpdateProductController(updateProduct)
	deleteProductController := productController.NewDeleteProductController(deleteProduct)

	// Controladores para citas
	createcitasController := citasController.NewCreatecitasController(createcitas)
	getcitasController := citasController.NewcitassController(getcitass)
	deletecitasController := citasController.NewDeletecitasController(deletecitass)
	updatecitasController := citasController.NewUpdatecitasController(updatecitass)

	// Configuración del enrutador de Gin
	r := gin.Default()

	// Configurar rutas de productos
	productRoutes.SetupProductRoutes(r, createProductController, getProductsController, updateProductController, deleteProductController)

	// Configurar rutas de citas
	citasRoutes.SetupcitasRoutes(r, createcitasController, getcitasController, deletecitasController, updatecitasController)

	// Iniciar servidor
	log.Println("server started at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
