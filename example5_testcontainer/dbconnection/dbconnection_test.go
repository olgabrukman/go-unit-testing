package dbconnection

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDBConnectionCRUD(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	myCollection, container := NewTestDBConnection(t)
	defer container.Terminate(Ctx)

	require.NotNil(myCollection)
	require.NotNil(container)

	key := "aaa"
	value := "333"
	insert, err := myCollection.Insert(key, value)
	require.NoError(err)
	assert.NotNil(insert)
	assert.NotZero(insert)

	dbValue, err := myCollection.GetValue(key)
	require.NoError(err)
	assert.Equal(value, dbValue)

	ok, err := myCollection.Delete(key)
	require.NoError(err)
	assert.True(ok)
}

func panicOnTimeout(d time.Duration) {
	<-time.After(d)
	panic("Test timed out")
}

//equivalent to running go test --timeout 5s
func TestMain(m *testing.M) {
	println("in test main")
	go panicOnTimeout(5 * time.Second) // custom timeout

	code := m.Run()
	os.Exit(code)
}

/*
Access mongoDB:
docker exec -it mongodb-container-id /bin/bash
mongo test
show dbs
use local
show collections
db.keyvalue.find()
*/
