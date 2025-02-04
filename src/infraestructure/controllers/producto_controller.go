package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"api-hexagonal-go/src/applications"
	"api-hexagonal-go/src/domain/entities"
)

type ProductoController struct {
	UseCase *applications.ProductoUseCase
}

func NewProductoController(useCase *applications.ProductoUseCase) *ProductoController {
	return &ProductoController{UseCase: useCase}
}

func (pc *ProductoController) GetAllProductos(c *gin.Context) {
	productos, err := pc.UseCase.GetAllProductos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, productos)
}

func (pc *ProductoController) GetProductoByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	producto, err := pc.UseCase.GetProductoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	c.JSON(http.StatusOK, producto)
}

func (pc *ProductoController) CreateProducto(c *gin.Context) {
	var producto entities.Producto
	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := pc.UseCase.CreateProducto(&producto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, producto)
}

func (pc *ProductoController) UpdateProducto(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var producto entities.Producto
	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := pc.UseCase.UpdateProducto(id, &producto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, producto)
}

func (pc *ProductoController) DeleteProducto(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := pc.UseCase.DeleteProducto(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado correctamente"})
}
