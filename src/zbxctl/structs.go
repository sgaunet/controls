package zbxctl

// Struct to represent the informations to connect to the Zabbix API
type ZbxCtl struct {
	ApiEndpoint       string `yaml:"apiEndpoint"`
	User              string `yaml:"user"`
	Password          string `yaml:"password"`
	Since             int    `yaml:"since"`
	SeverityThreshold int    `yaml:"severityThreshold"`
}
