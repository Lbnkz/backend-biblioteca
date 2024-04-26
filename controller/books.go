package controller

import (
	"backend-biblioteca/database"
	"backend-biblioteca/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase_SQL()
	var books []models.Book
	db.Find(&books)
	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
    db := database.GetDatabase_SQL()
    vars := mux.Vars(r)
    id := vars["id"]
    var books models.Book
    result := db.First(&books, id)
    if result.Error != nil {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "O livro com o ID %s não foi encontrado", id)
        return
    }
    json.NewEncoder(w).Encode(books)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
    db := database.GetDatabase_SQL()
    var books models.Book
    json.NewDecoder(r.Body).Decode(&books)
    result := db.Create(&books)
    if result.Error != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Erro ao adicionar o livro: %v", result.Error)
        return
    }
    json.NewEncoder(w).Encode(books)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
    db := database.GetDatabase_SQL()
    vars := mux.Vars(r)
    id := vars["id"]
    var books models.Book
    db.First(&books, id)
    json.NewDecoder(r.Body).Decode(&books)
    result := db.Save(&books)
    if result.Error != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Erro ao atualizar o livro: %v", result.Error)
        return
    }
    json.NewEncoder(w).Encode(books)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
    db := database.GetDatabase_SQL()
    vars := mux.Vars(r)
    id := vars["id"]
    var books models.Book
    result := db.Delete(&books, id)
    if result.Error != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Erro ao deletar o livro: %v", result.Error)
        return
    }
    fmt.Fprintf(w, "Livro com ID %s deletado com sucesso", id)
}