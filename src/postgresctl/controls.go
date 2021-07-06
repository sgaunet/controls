package postgresctl

import (
	"fmt"
	"strings"

	"sgaunet/controls/config"

	"github.com/fatih/color"
)

func (db *PostgresDB) PrintFailedCnx() {
	red := color.New(color.FgRed, color.Bold)
	red.Printf("DB Host    : %s\n", strings.Split(db.Cfg.Dbhost, ".")[0])
	red.Println("Nb cnx     : Connection error")
	red.Println("Nb Max cnx : Connection error")
	red.Println("Size (Go)  : Connection error")
	red.Println("Result     : Connection error")
}

func (db *PostgresDB) PrintResult() {
	green := color.New(color.FgGreen, color.Bold)
	green.Printf("DB Host    : %s\n", strings.Split(db.Cfg.Dbhost, ".")[0])
	green.Printf("Nb cnx     : %d\n", db.NbUsedConnections)
	green.Printf("Nb Max cnx : %d\n", db.NbMaxConnections)
	green.Printf("Size (Go)  : %d\n", db.Size/1024/1024/1024)
	green.Println("Result     : OK")
}

func LaunchControls(cfgdbs []config.DbConfig) [][]string {
	var resultTable [][]string
	resultTable = append(resultTable, []string{"DB", "Nb cnx", "Nb Max Cnx", "Size (Go)", "Result"})

	idx := 0
	for _, db := range cfgdbs {
		onedb := PostgresDB{
			Cfg: db,
		}
		err := onedb.Connect()
		defer onedb.Close()

		if err != nil {
			onedb.PrintFailedCnx()
			resultTable = append(resultTable, []string{strings.Split(onedb.Cfg.Dbhost, ".")[0], "N/A", "N/A", "N/A", "<span style=\"color:red\">Connection error</span>"})
		} else {
			onedb.CalcDatabaseSize()
			onedb.CalcCnx()
			onedb.PrintResult()
			resultTable = append(resultTable, []string{strings.Split(onedb.Cfg.Dbhost, ".")[0], fmt.Sprintf("%d", onedb.NbUsedConnections), fmt.Sprintf("%d", onedb.NbMaxConnections), fmt.Sprintf("%d", onedb.Size/1024/1024/1024), "<span style=\"color:green\">OK</span>"})
		}
		idx++
	}
	return resultTable
}
