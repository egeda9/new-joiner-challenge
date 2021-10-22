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
