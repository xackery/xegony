package storage

//Initializer is a generic interface of all storage types
type Initializer interface {
	VerifyTables() (err error)
	InsertTestData() (err error)
}
