package main

import (
	"backend-biblioteca/database"
	"backend-biblioteca/database/migrations"
	"backend-biblioteca/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	godotenv.Load(".env")
	database.StartDB_SQL()
	db := database.GetDatabase_SQL()
	migrations.RunMigrations(db)
	r := mux.NewRouter()
	routes.StartRoutes(r)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "OPTIONS", "DELETE", "PUT", "POST"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)
	PORT := os.Getenv("PORT")
	fmt.Println("Server running on port ", PORT)
	log.Fatal(http.ListenAndServe(PORT, handler))
	print(db)
}
