package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sgaunet/controls/config"
	"github.com/sgaunet/controls/httpctl"
	"github.com/sgaunet/controls/postgresctl"
	"github.com/sgaunet/controls/reportpdf"
	"github.com/sgaunet/controls/sshctl"
	zbxctl "github.com/sgaunet/controls/zbxctl"

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

	isConfOk := configApp.CheckAsserts()
	if !isConfOk {
		os.Exit(1)
	}

	//r := reportmd.New()
	rPdf := reportpdf.New()
	rPdf.AddTitle("Controls")

	if len(configApp.ZbxCtl.ApiEndpoint) != 0 {
		z, err := zbxctl.New(configApp.ZbxCtl.User, configApp.ZbxCtl.Password, configApp.ZbxCtl.ApiEndpoint, configApp.ZbxCtl.Since, configApp.ZbxCtl.SeverityThreshold)
		//rPdf.AddLine()
		rPdf.AddSection("Zabbix controls")
		fmt.Println()

		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot login into zabbix API")
			rPdf.AddTable("", z.FailedResultControls(err))
		} else {
			defer z.Logout()
			res, err := z.LaunchControls()
			if err == nil {
				rPdf.AddTable("", res)
			} else {
				rPdf.AddTable("", z.FailedResultControls(err))
			}
		}
	}

	if len(configApp.Db) != 0 {
		rPdf.AddSection("Databases controls")
		//r.AddTable(toto)
		rPdf.AddTable("", postgresctl.LaunchControls(configApp.Db))
	}

	if len(configApp.SshServers) != 0 {
		//rPdf.AddLine()
		rPdf.AddSection("SSH controls")
		// Loop over group of servers
		for serverGroupName, servers := range configApp.SshServers {
			for assertsGroupName, asserts := range configApp.SshAsserts {
				if serverGroupName == assertsGroupName {
					rPdf.AddTable("SSH controls for group "+serverGroupName, sshctl.LaunchControls(servers, asserts))
					break
				}
			}
		}
	}
	fmt.Println()
	//rPdf.AddPAgeBreak()

	if len(configApp.AssertsHTTP) != 0 {
		//rPdf.AddLine()
		rPdf.AddSection("HTTP controls")
		//r.AddTable("", httpctl.LaunchControls(configApp.AssertsHTTP))
		results := httpctl.LaunchControls(configApp.AssertsHTTP)
		rPdf.AddTable("HTTP", results)
		fmt.Println()
	}

	//rPdf.AddLine()
	//r.AddFooter()
	rPdf.AddFooter(version)
	err = rPdf.Export(reportPath)

	//err = r.Export(reportPath)
	if err != nil {
		log.Fatal(err)
	}
}
