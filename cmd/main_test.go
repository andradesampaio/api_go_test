package main

import (
	"api-go-test/internal/controller"
	"api-go-test/internal/database"
	"api-go-test/internal/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"bytes"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routers := gin.Default()
	routers.GET("/alunos", controller.ExibeTodosAlunos)
	routers.GET("/aluno/:id", controller.BuscaAluno)
	routers.GET("/aluno/cpf/:cpf", controller.BuscaAlunoPorCpf)
	routers.DELETE("/aluno/:id", controller.DeletaAluno)
	routers.PATCH("/aluno/:id", controller.EditaAluno)
	return routers
}


func CriaAlunoMock() {
	aluno := model.Aluno{Nome: "João", CPF: "12345678901", RG: "123456789", Email: "joao@gmail.com"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno model.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestListandoTodosOsAlunos(t *testing.T) {
    database.Connect()

	r := SetupDasRotasDeTeste()
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Deveria retornar status 200")
}

func TestBuscaAlunoPorCPF(t *testing.T){
	database.Connect()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	

	r := SetupDasRotasDeTeste()
	req, _ := http.NewRequest("GET", "/aluno/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Deveria retornar status 200")

}

func TestBuscaAlunoPorIDHandler(t *testing.T){
	database.Connect()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	

	r := SetupDasRotasDeTeste()
	req, _ := http.NewRequest("GET", "/aluno/"+strconv.Itoa(ID), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMock model.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	assert.Equal(t, "João", alunoMock.Nome, "Deveria retornar o nome João")
	assert.Equal(t, "12345678901", alunoMock.CPF, "Deveria retornar o CPF 12345678901")

}

func TestDeletaAlunoHandler(t *testing.T){
	database.Connect()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	

	r := SetupDasRotasDeTeste()
	req, _ := http.NewRequest("DELETE", "/aluno/"+strconv.Itoa(ID), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Deveria retornar status 200")

}

func TestEditaAlunoHandler(t *testing.T){
	database.Connect()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	
	aluno := model.Aluno{Nome: "Maria", CPF: "12345678901", RG: "123456789", Email: "teste@email.com"}
	payload, _ := json.Marshal(aluno)

	r := SetupDasRotasDeTeste()
	pathParaEditar := "/aluno/"+strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(payload))

	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMock model.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)

	assert.Equal(t, http.StatusOK, resposta.Code, "Deveria retornar status 200")
	assert.Equal(t, "Maria", alunoMock.Nome, "Deveria retornar o nome Maria")
}
