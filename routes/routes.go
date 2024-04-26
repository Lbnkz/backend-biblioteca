package routes

import (
	"backend-biblioteca/controller"

	"github.com/gorilla/mux"
)

func StartRoutes(r *mux.Router) {

	//Books
	r.HandleFunc("/books", controller.GetAllBooks).Methods("GET")
	r.HandleFunc("/books/{id}", controller.GetBookById).Methods("GET")
	r.HandleFunc("/books", controller.AddBook).Methods("POST")
	r.HandleFunc("/books/{id}", controller.UpdateBookById).Methods("PUT")
	r.HandleFunc("/books/{id}", controller.DeleteBookById).Methods("DELETE")
	
}