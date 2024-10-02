package config

import (
	"log"
	"os"

	"github.com/amaralfelipe1522/eva-database/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}
}

func InitDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar com o banco de dados: ", err)
	}

	// Chame a migração
	MigrateDB(db)

	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.Player{}, &models.Character{}, &models.Adventure{}, &models.Session{}, &models.SessionParticipant{})
}
