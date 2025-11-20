package db

//go:generate mockgen -source=interface.go -destination=../../../mocks/db_mock.go -package=mocks
type DB interface {
	Find(dest interface{}, conds ...interface{}) *DBResult
	First(dest interface{}, conds ...interface{}) *DBResult
	Create(value interface{}) *DBResult
}

type DBResult struct {
	Error error
}
