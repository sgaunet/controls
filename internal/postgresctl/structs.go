package postgresctl

import (
	"database/sql"
)

type PostgresDB struct {
	Cfg  DbConfig
	Conn *sql.DB
	Size int

	NbMaxConnections        int
	NbUsedConnections       int
	NbReservedForSuperUser  int
	NbReservedForNormalUser int
}

// Struct representing the connection to a postgres Database
type DbConfig struct {
	Dbhost      string `yaml:"dbhost"`
	Dbport      string `yaml:"dbport"`
	Dbuser      string `yaml:"dbuser"`
	Dbpassword  string `yaml:"dbpassword"`
	Dbname      string `yaml:"dbname"`
	Dbsizelimit int    `yaml:"sizelimit"`
}
