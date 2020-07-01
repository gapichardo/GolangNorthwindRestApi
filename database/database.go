package database

import "database/sql"

func InitDB() *sql.DB {
	connectionString := "root:picge02@tcp(localhost:3306)/northwind"
	databaseConnection, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err.Error()) // Error Handling = manejo de errores
	}
	return databaseConnection
}
