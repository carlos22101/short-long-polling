package controllers

import (
	"api-hexagonal-go/src/applications"
	"api-hexagonal-go/src/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsuarioController struct {
	UseCase *applications.UsuarioUseCase
}

func NewUsuarioController(useCase *applications.UsuarioUseCase) *UsuarioController {
	return &UsuarioController{UseCase: useCase}
}
func (uc *UsuarioController) GetAllUsuarios(c *gin.Context){
    usuarios, err := uc.UseCase.GetAllUsuarios()
    if err != nil{
        c.JSON(http.StatusInternalServerError , gin.H{"error" : err.Error()})
        return
    }
    c.JSON(http.StatusOK , usuarios)
}

func (uc *UsuarioController) GetUsuarioByID(c*gin.Context){
    id, _ :=strconv.Atoi(c.Param("id"))
    usuario , err := uc.UseCase.GetUsuarioByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound , gin.H{"error" : "usuario no encontrado"})
        return
    }
    c.JSON(http.StatusOK ,usuario)
}

func (uc *UsuarioController) CreateUsuario(c*gin.Context){
    var usuario entities.Usuario
    if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := uc.UseCase.CreateUsuario(&usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, usuario)
}

func (uc *UsuarioController) UpdateUsuario(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var usuario entities.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := uc.UseCase.UpdateUsuario(id, &usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, usuario)
}

func (uc *UsuarioController) DeleteUsuario(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := uc.UseCase.DeleteUsuario(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado correctamente"})
}