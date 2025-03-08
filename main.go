package main

import (
	"log"

	"github.com/alejandroimen/API_Consumer/src/core"
	productApp "github.com/alejandroimen/API_Consumer/src/products/application"
	productController "github.com/alejandroimen/API_Consumer/src/products/infrastructure/controllers"
	productRepo "github.com/alejandroimen/API_Consumer/src/products/infrastructure/repository"
	productRoutes "github.com/alejandroimen/API_Consumer/src/products/infrastructure/routes"
	ucitasApp "github.com/alejandroimen/API_Consumer/src/ucitas/application"
	ucitasController "github.com/alejandroimen/API_Consumer/src/ucitas/infrastructure/controllers"
	ucitasRepo "github.com/alejandroimen/API_Consumer/src/ucitas/infrastructure/repository"
	ucitasRoutes "github.com/alejandroimen/API_Consumer/src/ucitas/infrastructure/routes"
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
	ucitasRepository := ucitasRepo.NewCreateucitasRepoMySQL(db)

	// Casos de uso para productos
	createProduct := productApp.NewCreateProduct(productRepository)
	getProducts := productApp.NewGetProducts(productRepository)
	updateProduct := productApp.NewUpdateProduct(productRepository)
	deleteProduct := productApp.NewDeleteProduct(productRepository)

	// Casos de uso para ucitas
	createucitas := ucitasApp.NewCreateucitas(ucitasRepository)
	getucitass := ucitasApp.NewGetucitass(ucitasRepository)
	deleteucitass := ucitasApp.NewDeleteucitas(ucitasRepository)
	updateucitass := ucitasApp.NewUpdateucitas(ucitasRepository)

	// Controladores para productos
	createProductController := productController.NewCreateProductController(createProduct)
	getProductsController := productController.NewGetProductsController(getProducts)
	updateProductController := productController.NewUpdateProductController(updateProduct)
	deleteProductController := productController.NewDeleteProductController(deleteProduct)

	// Controladores para ucitas
	createucitasController := ucitasController.NewCreateucitasController(createucitas)
	getucitasController := ucitasController.NewucitassController(getucitass)
	deleteucitasController := ucitasController.NewDeleteucitasController(deleteucitass)
	updateucitasController := ucitasController.NewUpdateucitasController(updateucitass)

	// Configuración del enrutador de Gin
	r := gin.Default()

	// Configurar rutas de productos
	productRoutes.SetupProductRoutes(r, createProductController, getProductsController, updateProductController, deleteProductController)

	// Configurar rutas de ucitas
	ucitasRoutes.SetupucitasRoutes(r, createucitasController, getucitasController, deleteucitasController, updateucitasController)

	// Iniciar servidor
	log.Println("server started at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
