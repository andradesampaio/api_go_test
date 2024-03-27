# Curso Go Alura 
Go: validações, testes e páginas HTML

# Faça esse curso de GoLang e:
- Aprenda a validar os dados de uma API Go
- Saiba como escrever testes com Go de forma simples
- Aprenda a garantir um comportamento esperado das rotas de sua API
- Renderize páginas HTML com Gin 
- Configure e personalize a página 404 do Gin
- Mostre os dados de sua API nas páginas HTML

## Requisitos

- Docker
- Docker Compose
- Go

## Instruções de Uso

1. Clone este repositório para a sua máquina local.
2. Navegue até o diretório do projeto.
3. Antes de iniciar a aplicação, é necessário criar o banco de dados. Para isso, navegue até a pasta `infra` e execute `docker-compose up`.
4. O script de inicialização irá criar a tabela `alunos` e inserir 10 registros.

## Acessando o Banco de Dados

Para acessar o banco de dados em execução, use o seguinte comando:

``bash 
docker exec -it <container_id_or_name> /bin/bash

# Projeto com Docker e PostgreSQL

Este projeto utiliza Docker e PostgreSQL para criar um banco de dados.

4. Para testar a aplicação, você pode usar o seguinte comando curl:

```bash
curl --request POST \
  --url http://localhost:8080/alunos/ \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/8.6.1' \
  --data '{
	"nome": "",
	"CPF": "",
	"RG": "45.456.879-0",
	"Email" : "teste@gmail.com"
}'

curl --request PATCH \
  --url http://localhost:8080/aluno/1 \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/8.6.1' \
  --data '{
	"nome": "Sofia Herinque",
	"CPF": "123-000-435-01",
	"RG": "45.000.879-0",
	"Email" : "teste@gmail.com"
}'

curl --request DELETE \
  --url http://localhost:8080/aluno/11 \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/8.6.1'



  curl --request GET \
  --url http://localhost:8080/alunos \
  --header 'User-Agent: insomnia/8.6.1'


  curl --request GET \
  --url http://localhost:8080/aluno/1 \
  --header 'User-Agent: insomnia/8.6.1'


  curl --request GET \
  --url http://localhost:8080/aluno/cpf/34267543501 \
  --header 'User-Agent: insomnia/8.6.1'
