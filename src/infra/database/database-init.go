package database

import (
	"gorm.io/gorm"
)

type IFuncDialectOpen func(dsn string) gorm.Dialector

type DatabaseInitResult struct {
	database *gorm.DB
	Err      error
}

func DatabaseInit(dialect gorm.Dialector) DatabaseInitResult {
	var err error
	database, err := gorm.Open(dialect)
	return DatabaseInitResult{
		database: database,
		Err:      err,
	}
}
