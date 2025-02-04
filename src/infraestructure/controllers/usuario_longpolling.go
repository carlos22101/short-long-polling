package controllers

import (
	"api-hexagonal-go/src/applications"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type UsuarioLongPolling struct {
	UseCase  *applications.UsuarioUseCase
	mu       sync.Mutex
	waitList []chan []byte
	lastLen  int
}

func NewUsuarioLongPolling(useCase *applications.UsuarioUseCase) *UsuarioLongPolling {
	return &UsuarioLongPolling{UseCase: useCase, lastLen: 0}
}

func (ul *UsuarioLongPolling) StartLongPolling(c *gin.Context) {
	ch := make(chan []byte)
	ul.mu.Lock()
	ul.waitList = append(ul.waitList, ch)
	ul.mu.Unlock()

	select {
	case <-ch:
		c.JSON(http.StatusOK, gin.H{"message": "🔄 Se detectó un cambio en los usuarios"})
	case <-time.After(30 * time.Second): // Máximo 30s de espera
		c.JSON(http.StatusNoContent, gin.H{"message": "⏳ Sin cambios en usuarios"})
	}
}

func (ul *UsuarioLongPolling) NotifyChanges() {
	usuarios, err := ul.UseCase.GetAllUsuarios()
	if err != nil {
		fmt.Println("Error al obtener usuarios:", err)
		return
	}

	if len(usuarios) != ul.lastLen {
		fmt.Println("🚀 Notificando cambios a clientes Long Polling")
		ul.mu.Lock()
		for _, ch := range ul.waitList {
			close(ch)
		}
		ul.waitList = nil
		ul.mu.Unlock()
		ul.lastLen = len(usuarios)
	}
}
