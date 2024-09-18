package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/amaralfelipe1522/eva-database/models"
	"gorm.io/gorm"
)

// GetPlayers godoc
// @Summary Retrieve all players
// @Description Get a list of all players including their characters
// @Tags players
// @Accept json
// @Produce json
// @Success 200 {array} models.Player
// @Router /players [get]
func GetPlayers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var players []models.Player
	db.Preload("Characters").Find(&players)
	json.NewEncoder(w).Encode(players)
}

// CreatePlayer godoc
// @Summary Create a new player
// @Description Create a new player and return the created player object
// @Tags players
// @Accept json
// @Produce json
// @Param player body models.Player true "Player object to be created"
// @Success 201 {object} models.Player
// @Router /players [post]
func CreatePlayer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var player models.Player
	_ = json.NewDecoder(r.Body).Decode(&player)
	db.Create(&player)
	json.NewEncoder(w).Encode(player)
}
