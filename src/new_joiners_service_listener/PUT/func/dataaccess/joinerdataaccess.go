package joinerdataaccess

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

type Joiner struct {
	Id                             int
	Name                           string
	Stack                          string
	Role                           string
	Languages                      string
	JoinerMessageAcknowledgementId int
}

// UpdateJoiner update a record status
func (j Joiner) UpdateJoiner(id int) (int64, error) {

	connStr := os.Getenv("DATABASE_CONNECTION_STRING")
	if connStr == "" {
		fmt.Println("FATAL: expected environment variable DATABASE_CONNECTION_STRING not set")
		return -1, fmt.Errorf("FATAL: expected environment variable DATABASE_CONNECTION_STRING not set")
	}

	conn, err := sql.Open("mssql", connStr)

	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	fmt.Printf("Connected!\n")
	defer conn.Close()

	sql := fmt.Sprintf("UPDATE Joiner SET Stack = '%s', Role = '%s', Languages = '%s' WHERE Id= %d", j.Stack, j.Role, j.Languages, id)

	result, err := conn.Exec(sql)

	if err != nil {
		fmt.Println("Error updating row: " + err.Error())
		return -1, err
	}
	return result.LastInsertId()
}

// GetJoiner
func Get(input int) (Joiner, error) {

	var joiner Joiner

	connStr := os.Getenv("DATABASE_CONNECTION_STRING")
	if connStr == "" {
		fmt.Println("FATAL: expected environment variable DATABASE_CONNECTION_STRING not set")
		return joiner, fmt.Errorf("FATAL: expected environment variable DATABASE_CONNECTION_STRING not set")
	}

	conn, err := sql.Open("mssql", connStr)

	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	var name, stack, role, languages string
	var id, joinerMessageAcknowledgementId int

	// Query for a value based on a single row.
	if err := conn.QueryRow("SELECT * from Joiner where id = ?", input).Scan(&id, &name, &stack, &role, &languages, &joinerMessageAcknowledgementId); err != nil {
		if err == sql.ErrNoRows {
			return joiner, err
		}
		return joiner, err
	}

	joiner = Joiner{
		Id:                             id,
		Name:                           name,
		Stack:                          stack,
		Role:                           role,
		Languages:                      languages,
		JoinerMessageAcknowledgementId: joinerMessageAcknowledgementId,
	}

	return joiner, nil
}
