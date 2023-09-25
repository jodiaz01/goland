package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jodiaz01/go-gorm-restapi/db"
	"github.com/jodiaz01/go-gorm-restapi/models"
	"github.com/jodiaz01/go-gorm-restapi/routes"
)

func main() {

	db.DBConexion()                       //inicializa la conex
	db.DB.AutoMigrate(models.Articulos{}) //crea la migraciones
	db.DB.AutoMigrate(models.Categoria{}) //crea la migraciones

	Myrutas := mux.NewRouter()
	//Myrutas.Use(mux.CORSMethodMiddleware(Myrutas))
	//Myrutas.Headers("Access-Control-Request-Method", "*")

	Myrutas.HandleFunc("/", routes.HomeHandler)
	/// rutas de categorias
	Myrutas.HandleFunc("/get_categoria", routes.GetCategoriaHandler).Methods("GET")
	Myrutas.HandleFunc("/get_categoriaByid/{id}", routes.GetCategoriaByIdHandler).Methods("GET")
	Myrutas.HandleFunc("/add_categoria", routes.PostCategoriaHandler).Methods("POST")
	Myrutas.HandleFunc("/del_categoria/{id}", routes.DeleteCategoriaHandler).Methods("DELETE")
	/// rutas de articulos
	Myrutas.HandleFunc("/get_articulos", routes.GetArticulosHand).Methods("GET")
	Myrutas.HandleFunc("/get_articulosByid/{id}", routes.GetArticulosHandById).Methods("GET")
	Myrutas.HandleFunc("/get_articulos_catId/{id}", routes.GetArticulosHandCatId).Methods("GET")

	Myrutas.HandleFunc("/add_articulos", routes.PostArticulosHand).Methods("POST")
	Myrutas.HandleFunc("/del_articulos/{id}", routes.DeleteArticulosHand).Methods("DELETE")

	// Configurar CORS antes de iniciar el servidor
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "X-Requested-With", "Content-Type", "Authorization"}),
	)

	http.Handle("/", corsHandler(Myrutas)) // Aplicar CORS a todas las rutas

	http.ListenAndServe("172.21.0.2:8005", nil) // Iniciar el servidor HTTP
}
