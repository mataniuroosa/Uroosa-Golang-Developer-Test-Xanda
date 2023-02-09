package DBUtilities

// This file contain's all the common function for overall application

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Database struct {
		DATABASE_NAME             string `json:"database_name"`
		DATABASE_HOSTNAME         string `json:"database_host"`
		DATABASE_PORT             string `json:"database_port"`
		DATABASE_USER             string `json:"database_user"`
		DATABASE_PSWRD            string `json:"database_password"`
		SpaceCraptTableName       string `json:"space_craft_table_name"`
		ArmamentTableName         string `json:"armament_table_name"`
		AuthenticateUserTableName string `json:"authenticateUser_table_name"`
	} `json:"database"`
	SERVER_HOST string `json:"server_host"`
	SERVER_PORT string `json:"server_port"`
}

func LoadConfiguration(filename string) (Config, error) {
	var config Config
	ConfigFile, err := os.Open(filename)
	defer ConfigFile.Close()

	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(ConfigFile)
	err = jsonParser.Decode(&config)
	fmt.Println(config)
	return config, err
}

// GetDBConnection have database connection and error

func GetDbConnection() *sql.DB {

	fmt.Println("Starting the Application.......")
	config, _ := LoadConfiguration("ApplicationConfigProperties/config.json")
	fmt.Println("config:", config)
	dbconn, dberror := sql.Open("mysql", config.Database.DATABASE_USER+":"+config.Database.DATABASE_PSWRD+"@tcp("+config.Database.DATABASE_HOSTNAME+":"+config.Database.DATABASE_PORT+")/"+config.Database.DATABASE_NAME)
	if dberror != nil {
		panic(dberror.Error())
	}
	return dbconn
}

// CloseDbConnection have definition of connection Close
func CloseDbConnection(dbconn *sql.DB) {
	dbconn.Close()
}

// CheckDbErrorIsNotNull is a function to check all the db error
func CheckDbErrorIsNotNull(err error) {
	if err != nil {
		panic(err.Error())
	}
}
