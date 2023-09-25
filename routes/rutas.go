package routes

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, router *http.Request) {
	w.Write([]byte("hola"))
}
