package router

import (
	"github.com/gin-gonic/gin"
	"api-go-test/internal/controller"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("../internal/templates/*")
	r.Static("/assets", "../internal/assets")
	r.GET("/alunos", controller.ExibeTodosAlunos)
	r.POST("/alunos", controller.CriaNovoAluno)
	r.GET("/aluno/:id", controller.BuscaAluno)
	r.DELETE("/aluno/:id", controller.DeletaAluno)
	r.PATCH("/aluno/:id", controller.EditaAluno)
	r.GET("/aluno/cpf/:cpf", controller.BuscaAlunoPorCpf)
	r.GET("/index", controller.ExibePaginaIndex)
	r.NoRoute(controller.PaginaNaoEncontrada)
	r.Run()
}