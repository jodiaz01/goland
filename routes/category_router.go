package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jodiaz01/go-gorm-restapi/db"
	"github.com/jodiaz01/go-gorm-restapi/models"
)

func GetCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	var cat []models.Categoria
	db.DB.Find(&cat)
	json.NewEncoder(w).Encode(&cat)
}

func GetCategoriaByIdHandler(w http.ResponseWriter, r *http.Request) {
	var cat models.Categoria
	params := mux.Vars(r)

	db.DB.First(&cat, params["id"])
	if cat.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Categoria No existe"))
		return
	}
	json.NewEncoder(w).Encode(&cat)
}
func PostCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	var cat models.Categoria
	json.NewDecoder(r.Body).Decode(&cat)
	create := db.DB.Create(&cat)
	err := create.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No se Pudo Add Esta Categoria"))
		return

	}
	json.NewEncoder(w).Encode(&cat)
}
func DeleteCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	var cat models.Categoria
	params := mux.Vars(r)

	db.DB.First(&cat, params["id"])
	if cat.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Categoria No existe"))
		return
	}
	db.DB.Delete(&cat)
	json.NewEncoder(w).Encode(&cat)
}
