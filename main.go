package main

import (
	"api-hexagonal-go/src/applications"
	"api-hexagonal-go/src/infraestructure/controllers"
	"api-hexagonal-go/src/infraestructure/controllers/notifier"
	"api-hexagonal-go/src/infraestructure/database"
	"api-hexagonal-go/src/infraestructure/repositories"
	"api-hexagonal-go/src/infraestructure/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// 🔹 Servidor de Productos en el puerto 8000
func startServerProductos() {
	router := gin.Default()

	// Conectar BD
	database.InitDB()
	productoRepo := repositories.NewProductoRepository(database.DB)
	productoUseCase := applications.NewProductoUseCase(productoRepo)
	routes.ProductoRoutes(router, productoUseCase)

	// Short Polling
	productoPolling := controllers.NewProductoPolling(productoUseCase)
	go productoPolling.StartPolling()

	fmt.Println("🚀 Servidor de productos corriendo en http://localhost:8000")
	log.Fatal(router.Run(":8000"))
}

// 🔹 Servidor de Usuarios en el puerto 8001
func startServerUsuarios() {
	router := gin.Default()

	// Conectar BD
	database.InitDB()
	usuarioRepo := repositories.NewUsuarioRepository(database.DB)
	notifierInstance := notifier.NewNotifier() // ✅ Crear una única instancia de Notifier
	usuarioUseCase := applications.NewUsuarioUseCase(usuarioRepo, notifierInstance) // ✅ Pasar Notifier
	routes.UsuarioRoutes(router, usuarioUseCase)

	// Short Polling
	usuarioPolling := controllers.NewUsuarioPolling(usuarioUseCase)
	go usuarioPolling.StartPolling()

	// Long Polling
	usuarioLongPolling := controllers.NewUsuarioLongPolling(notifierInstance) // ✅ Usar el Notifier
	router.GET("/usuarios/longpolling", usuarioLongPolling.LongPollingHandler) // ✅ Ruta Long Polling

	fmt.Println("🚀 Servidor de usuarios corriendo en http://localhost:8001")
	log.Fatal(router.Run(":8001"))
}

func main() {
	go startServerProductos() // Servidor en puerto 8000
	go startServerUsuarios()  // Servidor en puerto 8001

	select {} // Mantiene el programa en ejecución
}
