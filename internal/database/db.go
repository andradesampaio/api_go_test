package database

import (
	"log"
	"api-go-test/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "5432")
	os.Setenv("USER", "root")
	os.Setenv("PASSWORD", "root")
	os.Setenv("DBNAME", "mydatabase")

	stringDeConexao := "host="+os.Getenv("HOST")+" user="+os.Getenv("USER")+" password="+os.Getenv("PASSWORD")+" dbname="+os.Getenv("DBNAME")+" port="+os.Getenv("PORT")+" sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&model.Aluno{})
}
