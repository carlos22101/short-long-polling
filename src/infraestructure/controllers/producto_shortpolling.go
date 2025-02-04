package controllers

import (
	"api-hexagonal-go/src/applications"
	"fmt"
	"time"
)

type ProductoPolling struct {
	UseCase *applications.ProductoUseCase
	lastLen int
}

func NewProductoPolling(useCase *applications.ProductoUseCase) *ProductoPolling {
	return &ProductoPolling{UseCase: useCase, lastLen: 0}
}

func (pp *ProductoPolling) StartPolling(){
	for {
		productos, err := pp.UseCase.GetAllProductos()
		if err != nil {
			fmt.Println("Error al obtener productos:", err)
		} else {
			if len(productos) != pp.lastLen {
				fmt.Println("🔄 Cambio detectado en productos. Total actual:", len(productos))
				pp.lastLen = len(productos)
			} else {
				fmt.Println("✅ No hay cambios en productos")
			}
		}
		time.Sleep(5 * time.Second) 
	}
}
