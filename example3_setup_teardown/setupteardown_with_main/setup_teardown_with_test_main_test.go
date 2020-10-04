package setupteardown_with_main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var initValue int

func TestOne(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(50, initValue)
	initValue = 30
}

func TestTwo(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(50, initValue)
	initValue = 20
}

// M is a type passed to a TestMain function to run the actual tests.
func TestMain(m *testing.M) {
	//setup
	initValue = 50
	//run all unit tests
	code := m.Run()
	//teardown
	initValue = 50
	os.Exit(code)
}
