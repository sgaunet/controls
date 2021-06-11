package postgresctl

import (
	"database/sql"

	"sgaunet/controls/config"
)

type PostgresDB struct {
	Cfg  config.DbConfig
	Conn *sql.DB
	Size int

	NbMaxConnections        int
	NbUsedConnections       int
	NbReservedForSuperUser  int
	NbReservedForNormalUser int
}
