package flyerdbutil

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v3"
)

// DBConfig contains metadata for the database
type DBConfig struct {
	AppVersion    string `yaml:"app_version"`
	AppDate       string `yaml:"app_date"`
	DBhost        string `yaml:"db_host"`
	DBname        string `yaml:"db_name"`
	DBport        int    `yaml:"db_port"`
	DBadmin       string `yaml:"db_admin"`
	DBpassword    string `yaml:"db_password"`
	DBdate        string `yaml:"db_date"`
	DefHost       string `yaml:"defender_host"`
	Defport       string `yaml:"defender_port"`
	DefDBhost     string `yaml:"defender_db_host"`
	DefDBname     string `yaml:"defender_db_name"`
	DefDBport     string `yaml:"defender_db_port"`
	DefDBadmin    string `yaml:"defender_db_admin"`
	DefDBpassword string `yaml:"defender_db_password"`
}

// ReadConfig reads server configuration from configuration file and returns the config struct
func ReadConfig(configFilePath string) (DBConfig, error) {
	var ServerConf DBConfig

	configYaml, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return ServerConf, err
	}

	err = yaml.Unmarshal(configYaml, &ServerConf)
	if err != nil {
		return ServerConf, err
	}

	return ServerConf, nil
}

// Connect is a function that establishes connection to chosen database
func Connect(database string, config DBConfig) (*sql.DB, error) {
	var postgres string

	switch database {
	case "server":
		postgres = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DBhost, config.DBport, config.DBadmin, config.DBpassword, config.DBname)

	case "defender":
		postgres = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DefDBhost, config.DefDBport, config.DefDBadmin, config.DefDBpassword, config.DefDBname)
	}
	return sql.Open("postgres", postgres)
}
