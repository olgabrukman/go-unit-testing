package example1_basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloAssertions(t *testing.T) {
	assert := assert.New(t)
	emptyUserResult := hello("")
	assert.NotNil(emptyUserResult) //nil
	assert.NotEmpty(emptyUserResult)
	assert.Equal(GenericGreeting, emptyUserResult)

	result := hello("Olga")
	assert.NotNil(result)
	assert.NotEmpty(result)
	assert.Equal("Hello Olga!", result)
}

func TestHelloRequire(t *testing.T) {
	require := require.New(t)
	emptyUserResult := hello("")
	require.NotNil(emptyUserResult) //nil
	require.NotEmpty(emptyUserResult)
	require.Equal(GenericGreeting, emptyUserResult)

	result := hello("Olga")
	require.NotNil(result)
	require.NotEmpty(result)
	require.Equal("Hello Olga!", result)
}
