package joinerdataaccess

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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
	Id                             int
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

// GetJoiner
func Get(input int) (Joiner, error) {

	config := configReader()

	// Connect to database
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		config.Database.Server, config.Database.Username, config.Database.Password, config.Database.Port, config.Database.DBName)

	conn, err := sql.Open("mssql", connString)

	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	var joiner Joiner

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
