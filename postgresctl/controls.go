package postgresctl

import (
	"fmt"

	"github.com/sgaunet/controls/results"
)

func (db *PostgresDB) resultsWhenErrCnx() (resultTable []results.Result) {
	msgErr := "Connection error"

	tests := []string{
		fmt.Sprintf("Nb cnx (%s)", db.Cfg.Dbhost),
		fmt.Sprintf("Nb MAX cnx (%s)", db.Cfg.Dbhost),
	}
	for _, t := range tests {
		result := results.Result{
			Title:  t,
			Result: msgErr,
			Pass:   false,
		}
		resultTable = append(resultTable, result)
		result.PrintToStdout()
	}
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
			// Cannot connect to DB so
			// init all resultTable with failed results
			for _, r := range onedb.resultsWhenErrCnx() {
				resultTable = append(resultTable, r)
			}
		} else {
			// Connexion to DB is ok
			onedb.CollectInfos() // !TODO Check err
			results := onedb.LaunchControl()
			for _, r := range results {
				resultTable = append(resultTable, r)
			}
		}
		idx++
	}
	return resultTable
}
