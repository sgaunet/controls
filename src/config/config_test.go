package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// setUp function, add a number to numbers slice
func setUp() {
	fmt.Printf("setUp tests\n")
}

// tearDown function, delete a number to numbers slice
func tearDown() {
	fmt.Printf("TEARDOWN\n")
}

func TestMain(m *testing.M) {

	setUp()
	code := m.Run()
	tearDown()

	os.Exit(code)
}

func TestIsSSHPortDefined(t *testing.T) {

	s := SSHServer{
		Host: "srv.mysociety.com",
	}
	//assert.Nil(t, err)
	assert.Equal(t, false, s.IsSSHPortDefined())
	s.Host = "srv.society.com:22"
	assert.Equal(t, true, s.IsSSHPortDefined())
	s.Host = "srv.society.com:2222"
	assert.Equal(t, true, s.IsSSHPortDefined())
}

func TestAddPortIfNotDefined(t *testing.T) {

	s := SSHServer{
		Host: "srv.mysociety.com",
	}
	//assert.Nil(t, err)
	s.AddPortIfNotPresent()
	assert.Equal(t, s.Host, "srv.mysociety.com:22")
	s.AddPortIfNotPresent()
	assert.Equal(t, s.Host, "srv.mysociety.com:22")
	s.Host = "srv.mysociety.com:6548"
	s.AddPortIfNotPresent()
	assert.Equal(t, s.Host, "srv.mysociety.com:6548")
}
