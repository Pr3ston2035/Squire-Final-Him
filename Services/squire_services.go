package Services

import (
	"github.com/Pr3ston2035/Squire-Final-Him/Models"
	"github.com/Pr3ston2035/Squire-Final-Him/Utils"
)

//Insertar datos en la base de datos

func AddSquireToDB(squire Models.Squire) error {
	query := "INSERT INTO mainpage (url,User, password) VALUES (?, ?, ?)"

	// Ejecutar la consulta SQL con los valores proporcionados
	_, err := Utils.DB.Exec(query, squire.Url, squire.Username, squire.Password)
	if err != nil {
		return err
	}

	return nil
}

func GetSquire() ([]Models.Squire, error) {
	query := "SELECT url, user, password FROM mainpage"

	// Ejecutar la consulta SQL
	rows, err := Utils.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterar sobre las filas y guardar los datos en una lista
	var passwords []Models.Squire
	for rows.Next() {
		var password Models.Squire
		err := rows.Scan(&password.Url, &password.Username, &password.Password)
		if err != nil {
			return nil, err
		}
		passwords = append(passwords, password)
	}

	return passwords, nil
}
