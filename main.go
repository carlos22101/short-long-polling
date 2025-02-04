package main

import (
	"api-hexagonal-go/src/infraestructure/controllers"
	"api-hexagonal-go/src/infraestructure/database"
	"api-hexagonal-go/src/infraestructure/repositories"
	"api-hexagonal-go/src/applications"
	"api-hexagonal-go/src/infraestructure/routes"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func startServerProductos() {
	router := gin.Default()

	// Conectar BD
	database.InitDB()
	productoRepo := repositories.NewProductoRepository(database.DB)
	productoUseCase := applications.NewProductoUseCase(productoRepo)
	routes.ProductoRoutes(router, productoUseCase)

	
	productoPolling := controllers.NewProductoPolling(productoUseCase)
	go productoPolling.StartPolling()

	fmt.Println("🚀 Servidor de productos corriendo en http://localhost:8000")
	log.Fatal(router.Run(":8000"))
}

func startServerUsuarios() {
	router := gin.Default()

	// Conectar BD
	database.InitDB()
	usuarioRepo := repositories.NewUsuarioRepository(database.DB)
	usuarioUseCase := applications.NewUsuariouseCase(usuarioRepo)
	routes.UsuarioRoutes(router, usuarioUseCase)

	// Iniciar Short Polling
	usuarioPolling := controllers.NewUsuarioPolling(usuarioUseCase)
	go usuarioPolling.StartPolling()

	// Iniciar Long Polling
	usuarioLongPolling := controllers.NewUsuarioLongPolling(usuarioUseCase)
	go func() {
		for {
			time.Sleep(5 * time.Second)
			usuarioLongPolling.NotifyChanges()
		}
	}()

	fmt.Println("🚀 Servidor de usuarios corriendo en http://localhost:8001")
	log.Fatal(router.Run(":8001"))
}

func main() {
	go startServerProductos() // Servidor en puerto 8000
	go startServerUsuarios()  // Servidor en puerto 8001

	select {} // Mantiene el programa en ejecución
}
