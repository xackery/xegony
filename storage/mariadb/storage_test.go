package mariadb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	var err error
	stor := Storage{}

	err = stor.Initialize("", nil)
	assert.Nil(t, err)
	err = stor.Initialize("", nil)
	assert.Nil(t, err)

	err = stor.DropTables()
	assert.Nil(t, err)
	err = stor.VerifyTables()
	assert.Nil(t, err)
	err = stor.InsertTestData()
	assert.Nil(t, err)
}
