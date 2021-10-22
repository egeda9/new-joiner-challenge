package joinerdataaccess

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

type Joiner struct {
	Id        int
	Name      string
	Stack     string
	Role      string
	Languages string
}

// UpdateJoiner update a record status
func Get() ([]Joiner, error) {

	connStr := os.Getenv("DATABASE_CONNECTION_STRING")
	if connStr == "" {
		fmt.Println("FATAL: expected environment variable DATABASE_CONNECTION_STRING not set")
		return nil, fmt.Errorf("FATAL: expected environment variable DATABASE_CONNECTION_STRING not set")
	}

	conn, err := sql.Open("mssql", connStr)

	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	sql := "SELECT * FROM Joiner"
	rows, err := conn.Query(sql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return nil, err
	}

	defer rows.Close()
	var result []Joiner

	for rows.Next() {
		var name, stack, role, languages string
		var id, joinerMessageAcknowledgementId int
		err := rows.Scan(&id, &name, &stack, &role, &languages, &joinerMessageAcknowledgementId)

		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return nil, err
		}

		fmt.Printf("ID: %d, Name: %s, Stack: %s, Role: %s, Languages: %s\n", id, name, stack, role, languages)
		j := Joiner{
			Id:        id,
			Name:      name,
			Stack:     stack,
			Role:      role,
			Languages: languages,
		}
		result = append(result, j)
	}
	return result, nil
}
