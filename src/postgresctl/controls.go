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
			onedb.CalcDatabaseSize()
			onedb.CalcCnx()

			if onedb.GetDBSizeGo() > onedb.Cfg.Dbsizelimit {
				result := results.Result{
					Title:  fmt.Sprintf("DB Size (%s - limit %d)", onedb.GetDbHost(), onedb.Cfg.Dbsizelimit),
					Result: fmt.Sprintf("%d Go", onedb.GetDBSizeGo()),
					Pass:   false,
				}
				result.PrintToStdout()
				resultTable = append(resultTable, result)

			} else {
				result := results.Result{
					Title:  fmt.Sprintf("DB Size (%s - limit %d)", onedb.GetDbHost(), onedb.Cfg.Dbsizelimit),
					Result: fmt.Sprintf("%d Go", onedb.GetDBSizeGo()),
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
