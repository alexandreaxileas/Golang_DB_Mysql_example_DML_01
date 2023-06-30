package dbfullexercise

import (
	"database/sql"
	"fmt"
)

// FnConnect - Function created in order to connecet to Database
func FnConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "principal:skinner@/high_school")
	if err != nil {
		return nil, err
	}
	fmt.Println("Conexão estabelecida")
	return db, nil
}
