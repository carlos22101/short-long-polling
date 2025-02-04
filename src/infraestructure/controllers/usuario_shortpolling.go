package controllers

import (
	"api-hexagonal-go/src/applications"
	"fmt"
	"time"
)

type UsuarioPolling struct {
	UseCase *applications.UsuarioUseCase
	lastLen int
}

func NewUsuarioPolling(useCase *applications.UsuarioUseCase) *UsuarioPolling {
	return &UsuarioPolling{UseCase: useCase, lastLen: 0}
}

func (up *UsuarioPolling) StartPolling() {
	for {
		usuarios, err := up.UseCase.GetAllUsuarios()
		if err != nil {
			fmt.Println("Error al obtener usuarios:", err)
		} else {
			if len(usuarios) != up.lastLen {
				fmt.Println("🔄 Cambio detectado en usuarios. Total actual:", len(usuarios))
				up.lastLen = len(usuarios)
			} else {
				fmt.Println("✅ No hay cambios en usuarios")
			}
		}
		time.Sleep(5 * time.Second) // Short Polling cada 5 segundos
	}
}
