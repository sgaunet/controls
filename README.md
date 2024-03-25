[![Go Report Card](https://goreportcard.com/badge/github.com/sgaunet/controls)](https://goreportcard.com/report/github.com/sgaunet/controls)


# controls

Little CLI to launch series of controls:

* SSH
* Get zabbix Problems
* Get connections number of postgresql Database
* Check HTTP requests

The CLI prints controls on stdout and save it in a markdown report.

Want a PDF ? [Convert the markdown to html](https://github.com/sgaunet/mdtohtml) and [the html to PDF ... ](https://wkhtmltopdf.org/)

**This tool is under development.**


# Example

Configuration file :

```bash
# get an example of configuration
controls -config
```

# Tests

## Pre requisites

* vagrant
* virtualbox
* Golang
* Docker
* docker-compose

## Launch tests

```
$ cd tst
$ vagrant up
...
$ cd zbxserver
$ docker-compose up -d   # will launch a local zabbix server
...
$ ./tests.sh
...
```

There is a postgreSQL server in the stack but there is no directory mounted to save the data. Within 1 or 2 minutes, a problem will be triggered which will be reported in the test. No need to configure anything in zabbix so...
