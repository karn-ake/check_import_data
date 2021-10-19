package repository

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
)

var (
	server = "192.168.142.51"
	port = 1433
	user = "n2nthuat\\0138811"
	pass = "@Liverp00l"
)

type queryDatabase struct{}

func NewQueryDatabase() Repository {
	return &queryDatabase{}
}



