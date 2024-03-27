package controller

import (
	"api-go-test/internal/database"
	"api-go-test/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)


func ExibeTodosAlunos(c *gin.Context){
	var alunos []model.Aluno
    if err := database.DB.Find(&alunos).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, alunos)
}

func CriaNovoAluno(c *gin.Context) {
	var aluno model.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := aluno.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAluno(c *gin.Context){
	id := c.Params.ByName("id")
	var aluno model.Aluno
    database.DB.First(&aluno, id)

	if aluno.Nome == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context){
	id := c.Params.ByName("id")
	var aluno model.Aluno
	database.DB.First(&aluno, id)

	if aluno.Nome == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	database.DB.Delete(&aluno)
	c.JSON(http.StatusOK, gin.H{"message": "Aluno deletado com sucesso!"})
}

func EditaAluno(c *gin.Context){
	id := c.Params.ByName("id")
	var aluno model.Aluno
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCpf(c *gin.Context){
	cpf := c.Params.ByName("cpf")
	var aluno model.Aluno
	database.DB.Where("cpf = ?", cpf).First(&aluno)

	if aluno.Nome == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func ExibePaginaIndex(c *gin.Context){
	var alunos []model.Aluno
	database.DB.Find(&alunos)

	c.HTML(http.StatusOK, "index.html", gin.H{"alunos": alunos})
}

func PaginaNaoEncontrada(c *gin.Context){
	c.HTML(http.StatusFound, "404.html",nil)
}