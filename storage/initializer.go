package storage

import ()

//Initializer is a generic interface of all storage types
type Initializer interface {
	DropTables() (err error)
	VerifyTables() (err error)
	InsertTestData() (err error)
}
