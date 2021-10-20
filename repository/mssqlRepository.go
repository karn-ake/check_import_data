package repository

import (
	"../entity"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	//"strconv"
)

type MssqlRepository interface {
	GetSecurityData() ([]entity.SecurityData, error)
	GetSecurityDatabyStkCode(stk string) (*entity.SecurityData, error)
	DBConn() (*sql.DB, error)
}

type mssqlRepository struct{}

func NewMssqlRepository() MssqlRepository {
	return &mssqlRepository{}
}

func (m *mssqlRepository) DBConn() (*sql.DB, error) {

	var cfg entity.DBConfig
	m.LoadConfig(&cfg)

	connConfig := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", cfg.Server, cfg.User, cfg.Pass, cfg.Port, cfg.DB)

	db, err := sql.Open("mssql", connConfig)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (m *mssqlRepository) GetSecurityData() ([]entity.SecurityData, error) {

	db, _ := m.DBConn()
	defer db.Close()

	var SecurityDatas []entity.SecurityData

	rows, err1 := db.Query("select stkcode, stkid, isincode from mf_StkInfoBK")
	if err1 != nil {
		log.Fatal(err1)
	}
	for rows.Next() {
		var sData entity.SecurityData
		err2 := rows.Scan(&sData.StkCode, &sData.StkId, &sData.IsinCode)
		if err2 != nil {
			log.Fatal(err2)
		}
		fmt.Println(sData)
		SecurityDatas = append(SecurityDatas, sData)
	}
	return SecurityDatas, nil
}

func (m *mssqlRepository) GetSecurityDatabyStkCode(stk string) (*entity.SecurityData, error) {

	db, _ := m.DBConn()
	defer db.Close()

	var sData entity.SecurityData

	err1 := db.QueryRow("select stkcode, stkid, isincode from mf_StkInfoBK where stkcode = $1", stk).Scan(&sData.StkCode, &sData.StkId, &sData.IsinCode)
	if err1 != nil {
		log.Fatal(err1)
	}

	return &sData, nil
}
