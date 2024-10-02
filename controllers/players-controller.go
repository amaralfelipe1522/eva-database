package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/amaralfelipe1522/eva-database/models"
	"github.com/gorilla/mux"
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
	w.WriteHeader(http.StatusOK)
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
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(player)
}

// DeletePlayer godoc
// @Summary Delete a player
// @Description Delete a player by ID and return a success message
// @Tags players
// @Accept json
// @Produce json
// @Param id path int true "Player ID"
// @Success 200 {object} ResponseMessage
// @Failure 404 {object} ErrorMessage
// @Router /players/{id} [delete]

func DeletePlayer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// Extrair o ID na URL
	vars := mux.Vars(r)
	id := vars["id"]

	var player models.Player
	if err := db.First(&player, id).Error; err != nil {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	db.Delete(&player)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Player deleted successfully"})
}
