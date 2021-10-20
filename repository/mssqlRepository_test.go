package repository

import (
	"testing"
	"fmt"
	//"log"
)

var repo MssqlRepository = NewMssqlRepository()

func TestDBConn(t *testing.T) {
	db, err := repo.DBConn()
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(db)
}

func TestGetSecurityData(t *testing.T) {
	db,_ := repo.GetSecurityData()

	for _,d := range db {
		fmt.Println(d)
	}
}

func TestGetSecurityDatabyStkCode(t *testing.T) {
	db,_ := repo.GetSecurityDatabyStkCode("PTT")

	fmt.Println(db)
}