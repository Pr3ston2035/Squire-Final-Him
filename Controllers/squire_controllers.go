package Controllers

import (
	"encoding/json"
	"github.com/JuanEstebanAstaiza/Squirrel/Models"
	"github.com/JuanEstebanAstaiza/Squirrel/Services"
	"net/http"
)

func AddSquire(w http.ResponseWriter, r *http.Request) {
	var squire Models.Squire
	err := json.NewDecoder(r.Body).Decode(&squire)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = Services.AddSquireToDB(squire)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetSquire(w http.ResponseWriter, _ *http.Request) {
	// Obtener las contraseñas de la base de datos
	passwords, err := Services.GetSquire()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convertir las contraseñas a formato JSON
	jsonPasswords, err := json.Marshal(passwords)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Establecer el encabezado de respuesta
	w.Header().Set("Content-Type", "application/json")

	// Escribir la respuesta JSON
	w.Write(jsonPasswords)
}
