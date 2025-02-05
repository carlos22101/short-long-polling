package controllers

import (
	"api-hexagonal-go/src/applications"
	"api-hexagonal-go/src/infraestructure/controllers/notifier"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type UsuarioLongPolling struct {
	UseCase  *applications.UsuarioUseCase
	mu       sync.Mutex
	waitList []chan struct{}
	lastLen  int
	Notifier *notifier.Notifier}

func NewUsuarioLongPolling(notifierInstance *notifier.Notifier) *UsuarioLongPolling {
	return &UsuarioLongPolling{Notifier: notifierInstance}
}


func (u *UsuarioLongPolling) LongPollingHandler(c *gin.Context) {
	timeout := time.After(30 * time.Second) // Tiempo máximo de espera
	updateChan := u.Notifier.Subscribe()

	select {
	case <-updateChan:
		c.JSON(http.StatusOK, gin.H{"message": "Datos actualizados"})
	case <-timeout:
		c.JSON(http.StatusNoContent, gin.H{"message": "Sin cambios"})
	}
}


func (ul *UsuarioLongPolling) NotifyChanges() {
	usuarios, err := ul.UseCase.GetAllUsuarios()
	if err != nil {
		fmt.Println("❌ Error al obtener usuarios:", err)
		return
	}

	fmt.Println("🟡 Usuarios actuales en la base de datos:", usuarios)

	ul.mu.Lock()
	fmt.Println("🔍 Revisando cambios... Última cantidad:", ul.lastLen, "Nueva cantidad:", len(usuarios))

	if len(usuarios) != ul.lastLen {
		fmt.Println("🚀 Notificando cambios a clientes Long Polling")

		// 🔥 IMPORTANTE: Cerrar los canales correctamente
		for _, ch := range ul.waitList {
			close(ch)  // 🔥 Esto debería despertar a los clientes
		}

		ul.waitList = nil      // Limpiar la lista de espera
		ul.lastLen = len(usuarios)  // Actualizar el último tamaño
	} else {
		fmt.Println("✅ No hay cambios en usuarios")
	}

	ul.mu.Unlock()
}

