package routes

import (
	"net/http"

	"github.com/amaralfelipe1522/eva-database/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	// Rotas dos jogadores
	router.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetPlayers(db, w, r)
	}).Methods("GET")

	router.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreatePlayer(db, w, r)
	}).Methods("POST")

	router.HandleFunc("/players/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeletePlayer(db, w, r)
	}).Methods("DELETE")

	return router
}
