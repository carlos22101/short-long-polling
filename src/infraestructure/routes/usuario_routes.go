package routes

import (
	"api-hexagonal-go/src/applications"
	"api-hexagonal-go/src/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func UsuarioRoutes(router *gin.Engine, usuarioUseCase *applications.UsuarioUseCase) {
	usuarioController := controllers.UsuarioController{UseCase: usuarioUseCase}
	usuarioLongPolling := controllers.NewUsuarioLongPolling(usuarioUseCase)


	usuarioGroup := router.Group("/usuarios")
	{
		usuarioGroup.GET("/", usuarioController.GetAllUsuarios)
		usuarioGroup.GET("/:id", usuarioController.GetUsuarioByID)
		usuarioGroup.POST("/", usuarioController.CreateUsuario)
		usuarioGroup.PUT("/:id", usuarioController.UpdateUsuario)
		usuarioGroup.DELETE("/:id", usuarioController.DeleteUsuario)
		usuarioGroup.GET("/longpolling", usuarioLongPolling.StartLongPolling)
	}
}
