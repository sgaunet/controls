package main

import (
	"flag"
	"fmt"
	"os"

	"sgaunet/controls/config"
	"sgaunet/controls/httpctl"
	"sgaunet/controls/postgresctl"
	"sgaunet/controls/reportmd"
	"sgaunet/controls/sshctl"
	zbxctl "sgaunet/controls/zbxctl"

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

var version string = "development"

func printVersion() {
	fmt.Println(version)
}

func main() {
	var configFile string
	var reportPath string
	var debugLevel string
	var vOption bool
	flag.StringVar(&configFile, "f", "", "YAML file to parse.")
	flag.StringVar(&reportPath, "o", "", "Report tot create (md format")
	flag.StringVar(&debugLevel, "d", "debug", "Debug level (info,warn,debug)")
	flag.BoolVar(&vOption, "v", false, "Get version")
	flag.Parse()

	initTrace(debugLevel)

	if vOption {
		printVersion()
		os.Exit(0)
	}

	if len(reportPath) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if configFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	configApp, err := config.ReadyamlConfigFile(configFile)
	if err != nil {
		os.Exit(1)
	}

	r := reportmd.New()
	r.AddTitle("Controls")

	if len(configApp.Db) != 0 {
		r.AddSection("Databases controls")
		//r.AddTable(toto)
		r.AddTable("", postgresctl.LaunchControls(configApp.Db))
	}

	if len(configApp.SshServers) != 0 {
		r.AddSection("SSH Asserts")
		// Loop over group of servers
		for serverGroupName, servers := range configApp.SshServers {

			// fmt.Printf("\nSSH Asserts for group %s :\n\n", serverGroupName)
			// report.WriteLines(1)
			// report.WriteTitle("SSH asserts for group "+serverGroupName, 3)
			// report.WriteLines(1)
			// Loop over group of asserts
			for assertsGroupName, asserts := range configApp.SshAsserts {
				if serverGroupName == assertsGroupName {
					r.AddTable("SSH asserts for group "+serverGroupName, sshctl.LaunchControls(servers, asserts))
					break
				}
			}
		}
	}
	fmt.Println()
	r.AddPAgeBreak()

	if len(configApp.AssertsHTTP) != 0 {
		r.AddSection("HTTP controls")
		r.AddTable("", httpctl.LaunchControls(configApp.AssertsHTTP))
		fmt.Println()
	}

	if len(configApp.ZbxCtl.ApiEndpoint) != 0 {
		r.AddPAgeBreak()
		r.AddSection("Zabbix controls")
		fmt.Println()

		z, err := zbxctl.New(configApp.ZbxCtl.User, configApp.ZbxCtl.Password, configApp.ZbxCtl.ApiEndpoint, configApp.ZbxCtl.Since, configApp.ZbxCtl.SeverityThreshold)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot login into zabbix API")
			r.AddTable("", z.FailedResultControls(err))
		} else {
			defer z.Logout()
			res, err := z.LaunchControls()
			if err == nil {
				r.AddTable("", res)
			} else {
				r.AddTable("", z.FailedResultControls(err))
			}
		}
	}

	r.AddFooter()

	err = r.Export(reportPath)
	if err != nil {
		log.Fatal(err)
	}
}
