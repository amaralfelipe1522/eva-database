package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amaralfelipe1522/eva-database/config"
	"github.com/amaralfelipe1522/eva-database/routes"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	// Configurações do banco de dados
	config.LoadEnvVariables()
	DB = config.InitDB()

	// Configuração das rotas
	router := routes.SetupRouter(DB)

	fs := http.FileServer(http.Dir("./docs"))
	router.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", fs))

	// Servindo o Swagger na rota /swagger/
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/docs/swagger.json"), // Caminho para o swagger.json
	))

	// Inicialização do servidor
	fmt.Println("Eva DB is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
