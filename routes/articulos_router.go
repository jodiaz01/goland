package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jodiaz01/go-gorm-restapi/db"
	"github.com/jodiaz01/go-gorm-restapi/models"
)

// todo los articulos
func GetArticulosHand(w http.ResponseWriter, r *http.Request) {

	var lsart []models.Articulos
	db.DB.Find(&lsart)
	json.NewEncoder(w).Encode(&lsart)
	// w.Write([]byte("Get Articulos All"))
}

// filtrar  por id
func GetArticulosHandById(w http.ResponseWriter, r *http.Request) {

	var articulos models.Articulos
	params := mux.Vars(r)
	db.DB.First(&articulos, params["id"])
	if articulos.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("no se Encontro Usuario"))
		return
	}
	json.NewEncoder(w).Encode(&articulos)
}

//por articulos contengan categoria id

func GetArticulosHandCatId(w http.ResponseWriter, r *http.Request) {

	var articulos models.Articulos
	params := mux.Vars(r)
	db.DB.First(&articulos, params["id"])
	if articulos.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("no se Encontro Usuario"))
		return
	}
	db.DB.Model(&articulos).Association("Categoria").Find(&articulos.Idcategoria)
	json.NewEncoder(w).Encode(&articulos)
}

// crete
func PostArticulosHand(w http.ResponseWriter, r *http.Request) {
	var create models.Articulos
	json.NewDecoder(r.Body).Decode(&create)
	createArt := db.DB.Create(&create)
	err := createArt.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("createNo Creado Revices los campo"))
	}
	//retorna lo que se creo
	json.NewEncoder(w).Encode(&create)

}

// borrando por id
func DeleteArticulosHand(w http.ResponseWriter, r *http.Request) {
	var art models.Articulos
	params := mux.Vars(r)
	db.DB.First(&art, params["id"])
	if art.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("no hay usuario con este ID"))
		return
	}
	/* +++++++ lo saca de la lista de articulo como una actualizacion  y llena la columna delete_at con la fecha Tambien lla no se puede filtral pero sigue en la bd*/
	db.DB.Delete(&art)

	/* +++++++++++++++++++ OJO OJO OJO OJO +++++++++++++++++++++++	lo borra por completo de la BD  solo usaur Si enverdad Quiere eliminar de la bd */
	// db.DB.Unscoped().Delete(&art)
}
