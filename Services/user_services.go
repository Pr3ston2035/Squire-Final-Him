package Services

import (
	"database/sql"
	"errors"
	"github.com/Pr3ston2035/Squire-Final-Him/Database"
	"github.com/Pr3ston2035/Squire-Final-Him/Models"
	"github.com/Pr3ston2035/Squire-Final-Him/Utils"
)

func AuthenticateUser(credentials Models.Credentials) (bool, error) {

	// Buscar al usuario por email en la base de datos
	user, err := Database.FindUserByEmail(credentials.Email)
	if err != nil {
		return false, err
	}

	// verificar si se encontro un usuario con el email proporcionado
	if user == nil {
		return false, nil
	}

	var storedPassword string
	err = Utils.DB.QueryRow("SELECT password FROM users WHERE email = ?", credentials.Email).Scan(&storedPassword)
	if err != nil {
		return false, err
	}

	return true, nil
}

func RegisterUser(user Models.Credentials) error {
	// Generar un ID único para el usuario
	userID, err := Utils.GenerateUserID()
	if err != nil {
		return err
	}

	// Insertar el usuario en la base de datos con el ID único generado
	_, err = Utils.DB.Exec("INSERT INTO users (id, nickname, email, password) VALUES (?, ?, ?, ?)", userID, user.Nickname, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func LoginUser(credentials Models.Credentials) (*Models.Credentials, error) {
	// Verificar las credenciales del usuario

	// Verificar las credenciales del usuario
	authenticated, err := AuthenticateUser(credentials)
	if err != nil {
		return nil, err
	}

	if !authenticated {
		// Las credenciales son inválidas
		return nil, nil
	}

	// Obtener la contraseña almacenada del usuario con el email proporcionado
	var storedPassword string
	err = Utils.DB.QueryRow("SELECT password FROM users WHERE email = ?", credentials.Email).Scan(&storedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No se encontró ningún usuario con el email dado
			return nil, nil
		}
		return nil, err
	}

	if storedPassword != credentials.Password {
		return nil, nil
	}

	// Obtener el usuario con el email proporcionado
	var user Models.Credentials
	err = Utils.DB.QueryRow("SELECT id, nickname, email FROM users WHERE email = ?", credentials.Email).Scan(&user.ID, &user.Nickname, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No se encontró ningún usuario con el email dado
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
