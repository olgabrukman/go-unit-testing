package init

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var initValue int

//init -Golang feature, not recommended for unit tests setup!
func init() {
	initValue = 50
}

func TestOne(t *testing.T) {
	test := require.New(t)
	test.Equal(50, initValue)
	initValue = 30
}

func TestTwo(t *testing.T) {
	test := require.New(t)
	test.Equal(50, initValue)
	initValue = 20
}
