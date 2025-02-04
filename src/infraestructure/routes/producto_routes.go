package routes

import (
	"api-hexagonal-go/src/applications"
	"api-hexagonal-go/src/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func ProductoRoutes(router *gin.Engine, productoUseCase *applications.ProductoUseCase) {
	productoController := controllers.ProductoController{UseCase: productoUseCase}

	productoGroup := router.Group("/productos")
	{
		productoGroup.GET("/", productoController.GetAllProductos)
		productoGroup.GET("/:id", productoController.GetProductoByID)
		productoGroup.POST("/", productoController.CreateProducto)
		productoGroup.PUT("/:id", productoController.UpdateProducto)
		productoGroup.DELETE("/:id", productoController.DeleteProducto)
	}
}
