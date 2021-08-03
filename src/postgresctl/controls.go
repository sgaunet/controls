package postgresctl

import (
	"fmt"
	"strings"

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

func (db *PostgresDB) PrintResult(msg string, testPassed bool) {
	var std *color.Color
	if testPassed {
		std = color.New(color.FgGreen, color.Bold)
	} else {
		std = color.New(color.FgRed, color.Bold)
	}
	std.Printf("DB Host    : %s\n", strings.Split(db.Cfg.Dbhost, ".")[0])
	std.Printf("Nb cnx     : %d\n", db.NbUsedConnections)
	std.Printf("Nb Max cnx : %d\n", db.NbMaxConnections)
	std.Printf("Size (Go)  : %d\n", db.Size/1024/1024/1024)
	std.Printf("Result     : %s\n", msg)
}

func LaunchControls(cfgdbs []DbConfig) [][]string {
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
			msg := "<span style=\"color:green\">OK</span>"
			msgStdout := "OK"
			testPassed := true

			onedb.CalcDatabaseSize()
			onedb.CalcCnx()

			size := onedb.Size / 1024 / 1024 / 1024
			if size > onedb.Size {
				msg = "<span style=\"color:red\">Size limit reached</span>"
				msgStdout = "Size limit reached"
				testPassed = false
			}
			if onedb.NbUsedConnections == onedb.NbMaxConnections {
				msg = "<span style=\"color:red\">Size limit reached</span>"
				msgStdout = "Size limit reached"
				testPassed = false
			}
			onedb.PrintResult(msgStdout, testPassed)
			resultTable = append(resultTable, []string{strings.Split(onedb.Cfg.Dbhost, ".")[0], fmt.Sprintf("%d", onedb.NbUsedConnections), fmt.Sprintf("%d", onedb.NbMaxConnections), fmt.Sprintf("%d", size), msg})
		}
		idx++
	}
	return resultTable
}
