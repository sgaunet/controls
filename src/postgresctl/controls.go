package postgresctl

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/sgaunet/controls/results"
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

func (db *PostgresDB) resultsWhenErrCnx() (resultTable []results.Result) {
	result := results.Result{
		Title:  fmt.Sprintf("DB Size (%s)", db.Cfg.Dbhost),
		Result: "Connection error",
		Pass:   false,
	}
	result.PrintToStdout()
	resultTable = append(resultTable, result)

	result.Title = fmt.Sprintf("Nb cnx (%s)", db.Cfg.Dbhost)
	result.PrintToStdout()
	resultTable = append(resultTable, result)

	result.Title = fmt.Sprintf("Nb MAX cnx (%s)", db.Cfg.Dbhost)
	result.PrintToStdout()
	resultTable = append(resultTable, result)
	return
}

func LaunchControls(cfgdbs []DbConfig) (resultTable []results.Result) {
	//resultTable = append(resultTable, []string{"DB", "Nb cnx", "Nb Max Cnx", "Size (Go)", "Result"})

	idx := 0
	for _, db := range cfgdbs {
		onedb := PostgresDB{
			Cfg: db,
		}
		err := onedb.Connect()
		defer onedb.Close()

		if err != nil {
			for _, r := range onedb.resultsWhenErrCnx() {
				resultTable = append(resultTable, r)
			}

		} else {
			onedb.CalcDatabaseSize()
			onedb.CalcCnx()
			// Convert size to Go
			size := onedb.Size / 1024 / 1024 / 1024

			if size > onedb.Size {
				result := results.Result{
					Title:  fmt.Sprintf("DB Size (%s - limit %d)", onedb.GetDbHost(), onedb.Size),
					Result: fmt.Sprintf("%d Go", size),
					Pass:   false,
				}
				result.PrintToStdout()
				resultTable = append(resultTable, result)

			} else {
				result := results.Result{
					Title:  fmt.Sprintf("DB Size (%s - limit %d)", onedb.GetDbHost(), onedb.Size),
					Result: fmt.Sprintf("%d Go", size),
					Pass:   true,
				}
				result.PrintToStdout()
				resultTable = append(resultTable, result)
			}
			if onedb.NbUsedConnections == onedb.NbMaxConnections {
				result := results.Result{
					Result: "Nb used connections == Nb max connections",
					Pass:   false,
				}
				result.Title = fmt.Sprintf("Nb connections (%s)", onedb.Cfg.Dbhost)
				result.PrintToStdout()
				resultTable = append(resultTable, result)
				result.Title = fmt.Sprintf("Nb max connections (%s)", onedb.Cfg.Dbhost)
				result.PrintToStdout()
				resultTable = append(resultTable, result)
			} else {
				result := results.Result{
					Result: fmt.Sprintf("%d", onedb.NbUsedConnections),
					Pass:   true,
				}
				result.Title = fmt.Sprintf("Nb connections (%s)", onedb.Cfg.Dbhost)
				result.PrintToStdout()
				resultTable = append(resultTable, result)
				result.Title = fmt.Sprintf("Nb max connections (%s)", onedb.Cfg.Dbhost)
				result.Result = fmt.Sprintf("%d", onedb.NbMaxConnections)
				result.PrintToStdout()
				resultTable = append(resultTable, result)
			}
		}
		idx++
	}
	return resultTable
}
