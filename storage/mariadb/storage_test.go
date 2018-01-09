package mariadb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	var err error
	stor := Storage{}

	err = stor.Initialize("root@tcp(127.0.0.1:3306)/eqemu_test?charset=utf8&parseTime=true", nil)
	assert.Nil(t, err)
	err = stor.Initialize("root@tcp(127.0.0.1:3306)/eqemu_test?charset=utf8&parseTime=true", nil)
	assert.Nil(t, err)

	err = stor.DropTables()
	assert.Nil(t, err)
	err = stor.VerifyTables()
	assert.Nil(t, err)
	err = stor.InsertTestData()
	assert.Nil(t, err)
}
