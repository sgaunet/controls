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
	Dbhost      string `yaml:"dbhost" validate:"required"`
	Dbport      string `yaml:"dbport" validate:"required"`
	Dbuser      string `yaml:"dbuser" validate:"required"`
	Dbpassword  string `yaml:"dbpassword" validate:"required"`
	Dbname      string `yaml:"dbname" validate:"required"`
	Dbsizelimit int    `yaml:"sizelimit" validate:"required"`
}
