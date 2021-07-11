package sshserver

// Struct representing the connection to an SSH Server
type SSHServer struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Sshkey   string `yaml:"sshkey"`
}
