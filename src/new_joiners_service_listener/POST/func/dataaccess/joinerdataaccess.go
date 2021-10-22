package joinerdataaccess

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

type Joiner struct {
	Name                           string
	Stack                          string
	Role                           string
	Languages                      string
	JoinerMessageAcknowledgementId int
}

// CreateJoiner create a joiner
func (j Joiner) CreateJoiner() (int, error) {

	connStr := os.Getenv("DATABASE_CONNECTION_STRING")
	if connStr == "" {
		fmt.Println("FATAL: expected environment variable DATABASE_CONNECTION_STRING not set")
		return 0, fmt.Errorf("FATAL: expected environment variable DATABASE_CONNECTION_STRING not set")
	}

	conn, err := sql.Open("mssql", connStr)

	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	fmt.Printf("Connected!\n")
	defer conn.Close()

	sql := fmt.Sprintf("INSERT INTO Joiner (Name, Stack, Role, Languages, JoinerMessageAcknowledgementId) OUTPUT Inserted.Id VALUES ('%s','%s','%s','%s',%d);",
		j.Name, j.Stack, j.Role, j.Languages, j.JoinerMessageAcknowledgementId)

	lastInsertId := 0
	err = conn.QueryRow(sql).Scan(&lastInsertId)

	if err != nil {
		fmt.Println("Error inserting new row: " + err.Error())
		return -1, err
	}

	return lastInsertId, nil
}

// ReadJoiner read a joiner
func (j Joiner) GetJoiner(name string) (int, error) {

	connStr := os.Getenv("DATABASE_CONNECTION_STRING")
	if connStr == "" {
		fmt.Println("FATAL: expected environment variable DATABASE_CONNECTION_STRING not set")
		return 0, fmt.Errorf("FATAL: expected environment variable DATABASE_CONNECTION_STRING not set")
	}

	conn, err := sql.Open("mssql", connStr)

	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	sql := fmt.Sprintf("SELECT * FROM Joiner WHERE Name = '%s';", name)
	rows, err := conn.Query(sql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return -1, err
	}

	defer rows.Close()
	count := 0

	for rows.Next() {
		var name, stack, role, language string
		var id, joinerMessageAcknowledgementId int
		err := rows.Scan(&id, &name, &stack, &role, &language, &joinerMessageAcknowledgementId)

		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return -1, err
		}

		fmt.Printf("ID: %d, Name: %s, Stack: %s, Role: %s, Language: %s\n", id, name, stack, role, language)
		count++
	}
	return count, nil
}

// CreateJoiner create a joiner
func (j Joiner) CreateJoinerMessageAcknowledgement(message string) (int, error) {

	connStr := os.Getenv("DATABASE_CONNECTION_STRING")
	if connStr == "" {
		fmt.Println("FATAL: expected environment variable DATABASE_CONNECTION_STRING not set")
		return 0, fmt.Errorf("FATAL: expected environment variable DATABASE_CONNECTION_STRING not set")
	}

	conn, err := sql.Open("mssql", connStr)

	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	fmt.Printf("Connected!\n")
	defer conn.Close()

	sql := fmt.Sprintf("INSERT INTO JoinerMessageAcknowledgement (CreatedDate, Status, Message) OUTPUT Inserted.Id VALUES ('%s','%s','%s');",
		time.Now().Format("2006-01-02 15:04:05"), "Incomplete", message)

	lastInsertId := 0
	err = conn.QueryRow(sql).Scan(&lastInsertId)

	if err != nil {
		fmt.Println("Error inserting new row: " + err.Error())
		return -1, err
	}

	return lastInsertId, nil
}

// UpdateJoinerMessageAcknowledgementStatus update a record status
func (j Joiner) UpdateJoinerMessageAcknowledgementStatus() (int64, error) {

	connStr := os.Getenv("DATABASE_CONNECTION_STRING")
	if connStr == "" {
		fmt.Println("FATAL: expected environment variable DATABASE_CONNECTION_STRING not set")
		return 0, fmt.Errorf("FATAL: expected environment variable DATABASE_CONNECTION_STRING not set")
	}

	conn, err := sql.Open("mssql", connStr)

	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	fmt.Printf("Connected!\n")
	defer conn.Close()

	sql := fmt.Sprintf("UPDATE JoinerMessageAcknowledgement SET Status = 'Complete' WHERE Id= %d", j.JoinerMessageAcknowledgementId)

	result, err := conn.Exec(sql)

	if err != nil {
		fmt.Println("Error updating row: " + err.Error())
		return -1, err
	}
	return result.LastInsertId()
}
