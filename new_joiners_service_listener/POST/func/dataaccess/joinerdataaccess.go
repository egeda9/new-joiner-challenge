package joinerdataaccess

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		DBName   string `yaml:"dbname"`
		Port     int    `yaml:"port"`
		Server   string `yaml:"server"`
	} `yaml:"database"`
}

type Joiner struct {
	Name                           string
	Stack                          string
	Role                           string
	Languages                      string
	JoinerMessageAcknowledgementId int
}

func configReader() Config {

	path, _ := os.Getwd()
	f, err := os.Open(path + "\\config.yml")

	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		log.Panic(err)
	}

	return cfg
}

// CreateJoiner create a joiner
func (j Joiner) CreateJoiner() (int, error) {

	config := configReader()

	// Connect to database
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		config.Database.Server, config.Database.Username, config.Database.Password, config.Database.Port, config.Database.DBName)

	conn, err := sql.Open("mssql", connString)

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

	config := configReader()

	// Connect to database
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		config.Database.Server, config.Database.Username, config.Database.Password, config.Database.Port, config.Database.DBName)

	conn, err := sql.Open("mssql", connString)

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

	config := configReader()

	// Connect to database
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		config.Database.Server, config.Database.Username, config.Database.Password, config.Database.Port, config.Database.DBName)

	conn, err := sql.Open("mssql", connString)

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

	config := configReader()

	// Connect to database
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		config.Database.Server, config.Database.Username, config.Database.Password, config.Database.Port, config.Database.DBName)

	conn, err := sql.Open("mssql", connString)

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
