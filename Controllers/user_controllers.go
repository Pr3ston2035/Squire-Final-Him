package Controllers

import (
	"encoding/json"
	"github.com/JuanEstebanAstaiza/Squirrel/Models"
	"github.com/JuanEstebanAstaiza/Squirrel/Services"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user Models.Credentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	// Registrar al usuario
	err = Services.RegisterUser(user)
	if err != nil {
		http.Error(w, "Error al registrar el usuario", http.StatusInternalServerError)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusCreated)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Parsear los datos del cuerpo de la solicitud
	var user Models.Credentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtener el usuario con las credenciales proporcionadas
	userInfo, err := Services.LoginUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if userInfo == nil {
		// Las credenciales son inválidas
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	// Si las credenciales son válidas, enviar la información del usuario
	if userInfo != nil {
		// Escribir encabezado de respuesta exitosa
		w.WriteHeader(http.StatusOK)
		// Codificar el mensaje de "Acceso concedido" junto con los detalles del usuario
		response := map[string]interface{}{
			"message": "Acceso concedido",
			"user":    userInfo,
		}
		// Codificar la respuesta en formato JSON y enviarla
		json.NewEncoder(w).Encode(response)
	}

}
