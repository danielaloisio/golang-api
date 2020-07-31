package methods

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)

func Config() *mux.Router {
	router := mux.NewRouter()

	router.Handle("/contact", negroni.New(
		negroni.HandlerFunc(ConfigJwt().HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(GetPeople)))).Methods("GET")

	router.HandleFunc("/contact/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/contact", CreatePerson).Methods("POST")
	router.HandleFunc("/contact/{id}", DeletePerson).Methods("DELETE")

	return router
}
