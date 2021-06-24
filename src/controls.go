package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"sgaunet/controls/httpctl"
	"sgaunet/controls/postgresctl"
	"sgaunet/controls/sshctl"
	zabbixapi "sgaunet/controls/zabbixApi"

	"github.com/atsushinee/go-markdown-generator/doc"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func initTrace(debugLevel string) {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})
	// log.SetFormatter(&log.TextFormatter{
	// 	DisableColors: true,
	// 	FullTimestamp: true,
	// })

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	switch debugLevel {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	var configFile string
	var reportPath string
	var debugLevel string
	flag.StringVar(&configFile, "f", "", "YAML file to parse.")
	flag.StringVar(&reportPath, "o", "", "Report tot create (md format")
	flag.StringVar(&debugLevel, "d", "debug", "Debug level (info,warn,debug)")
	flag.Parse()

	initTrace(debugLevel)

	if len(reportPath) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if configFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	configApp, err := ReadyamlConfigFile(configFile)
	if err != nil {
		os.Exit(1)
	}

	report := doc.NewMarkDown()
	report.WriteTitle("Controls", doc.LevelTitle).WriteLines(2)

	if len(configApp.Db) != 0 {
		report.WriteTitle("Databases control", 2)
		fmt.Printf("DB controls :\n\n")
		report.WriteLines(1)
		report = postgresctl.LaunchControls(configApp.Db, report)
		report.WriteLines(2)
		fmt.Println()
	}

	if len(configApp.SshServers) != 0 {
		report.WriteTitle("SSH Asserts", 2)
		// Loop over group of servers
		for serverGroupName, servers := range configApp.SshServers {
			fmt.Printf("\nSSH Asserts for group %s :\n\n", serverGroupName)
			report.WriteLines(1)
			report.WriteTitle("SSH asserts for group "+serverGroupName, 3)
			report.WriteLines(1)
			// Loop over group of asserts
			for assertsGroupName, asserts := range configApp.SshAsserts {
				if serverGroupName == assertsGroupName {
					report = sshctl.LaunchControls(servers, asserts, report)
					break
				}
			}
		}
	}
	// fmt.Println(configApp.SshAsserts)
	// fmt.Println(configApp.SshServers)
	// report.WriteLines(2)
	fmt.Println()

	if len(configApp.AssertsHTTP) != 0 {
		report.WriteTitle("HTTP controls", 2)
		fmt.Printf("HTTP controls :\n\n")
		report.WriteLines(1)
		report = httpctl.LaunchControls(configApp.AssertsHTTP, report)
		report.WriteLines(2)
		fmt.Println()
	}

	err = report.Export(reportPath)
	if err != nil {
		log.Fatal(err)
	}

	if len(configApp.ZbxCtl.ApiEndpoint) != 0 {
		z, err := zabbixapi.New(configApp.ZbxCtl.User, configApp.ZbxCtl.Password, configApp.ZbxCtl.ApiEndpoint, configApp.ZbxCtl.Since, configApp.ZbxCtl.SeverityThreshold)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot login into zabbix API")
		} else {
			defer z.Logout()

			report.WriteTitle("Zabbix controls", 2)
			fmt.Printf("Zabbix controls :\n\n")
			report.WriteLines(1)
			report = z.LaunchControls(report)
			report.WriteLines(2)
			fmt.Println()
		}
	}

	report.WriteTitle("Infos", 2)
	report.WriteLines(1)
	report.Write("Generated at :" + time.Now().String() + "<br>")

	err = report.Export(reportPath)
	if err != nil {
		log.Fatal(err)
	}

}
