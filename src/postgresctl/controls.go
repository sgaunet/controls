package postgresctl

import (
	"fmt"
	"strings"

	"sgaunet/controls/config"

	"github.com/fatih/color"
)

func LaunchControls(cfgdbs []config.DbConfig) [][]string {
	var resultTable [][]string
	red := color.New(color.FgRed, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	// nbCfg := len(cfgdbs)
	// t := doc.NewTable(nbCfg, 5)
	// t.SetTitle(0, "DB")
	// t.SetTitle(1, "Nb CNX")
	// t.SetTitle(2, "Nb max CNX")
	// t.SetTitle(3, "Size (Go)")
	// t.SetTitle(4, "Result")
	resultTable = append(resultTable, []string{"DB", "Nb cnx", "Nb Max Cnx", "Size (Go)", "Result"})

	//fmt.Printf("%30s | %10s | %10s | %10s | %40s |\n", "DB", "NB CNX", "NB MAX CNX", "Size (Go)", "Result")
	idx := 0
	for _, db := range cfgdbs {
		onedb := PostgresDB{
			Cfg: db,
		}
		err := onedb.Connect()
		defer onedb.Close()

		if err != nil {
			comment := "Erreur de connexion du serveur "
			red.Printf("%30s | %10s | %10s | %10s | %40s |\n", strings.Split(onedb.Cfg.Dbhost, ".")[0], "N/A", "N/A", "N/A", comment)
			// t.SetContent(idx, 0, strings.Split(onedb.Cfg.Dbhost, ".")[0])
			// t.SetContent(idx, 1, "N/A")
			// t.SetContent(idx, 2, "N/A")
			// t.SetContent(idx, 3, "N/A")
			// t.SetContent(idx, 4, "<span style=\"color:red\">"+comment+"</span>")
			resultTable = append(resultTable, []string{strings.Split(onedb.Cfg.Dbhost, ".")[0], "N/A", "N/A", "N/A", "<span style=\"color:red\">" + comment + "</span>"})
		} else {
			onedb.CalcDatabaseSize()
			onedb.CalcCnx()
			comment := "OK"
			green.Printf("%30s | %10d | %10d | %10d | %40s |\n", strings.Split(onedb.Cfg.Dbhost, ".")[0], onedb.NbUsedConnections, onedb.NbMaxConnections, onedb.Size/1024/1024/1024, comment)
			// t.SetContent(idx, 0, strings.Split(onedb.Cfg.Dbhost, ".")[0])
			// t.SetContent(idx, 1, fmt.Sprintf("%d", onedb.NbUsedConnections))
			// t.SetContent(idx, 2, fmt.Sprintf("%d", onedb.NbMaxConnections))
			// t.SetContent(idx, 3, fmt.Sprintf("%d", onedb.Size/1024/1024/1024))
			// t.SetContent(idx, 4, "<span style=\"color:green\">"+comment+"</span>")
			resultTable = append(resultTable, []string{strings.Split(onedb.Cfg.Dbhost, ".")[0], fmt.Sprintf("%d", onedb.NbUsedConnections), fmt.Sprintf("%d", onedb.NbMaxConnections), fmt.Sprintf("%d", onedb.Size/1024/1024/1024), "<span style=\"color:green\">" + comment + "</span>"})
		}
		idx++
	}
	return resultTable
}
