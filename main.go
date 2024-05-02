package main

import (
	"fmt"
	"github.com/Pr3ston2035/Squire-Final-Him/Controllers"
	"github.com/Pr3ston2035/Squire-Final-Him/Utils"
	"github.com/rs/cors"
	"log"
	"net/http"

	_ "github.com/Pr3ston2035/Squire-Final-Him/Controllers"
	"github.com/gorilla/mux"
)

func main() {

	err := Utils.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/user/register", Controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/api/user/login", Controllers.LoginUser).Methods("POST")
	router.HandleFunc("/api/insert-squirrel/AddSquire", Controllers.AddSquire).Methods("POST")
	router.HandleFunc("/api/view/view-squire", Controllers.GetSquire).Methods("GET")

	// Opciones CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Permitir solicitudes desde cualquier origen
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(router)
	// Iniciar el servidor
	port := ":7250"
	fmt.Println("Servidor escuchando en el puerto", port)
	log.Fatal(http.ListenAndServe(port, handler))

}
