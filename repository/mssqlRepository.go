package repository

import (
	"../entity"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

type MssqlRepository interface {
	GetSecurityData() ([]entity.SecurityData, error)
	GetSecurityDatabyStkCode(stk string) (*entity.SecurityData, error)
	GetForeignData() ([]entity.ForeignData, error)
	GetForeignDatabyStkCode(stk string) (*entity.ForeignData, error)
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
			return nil, err2
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
		return nil, err1
	}

	return &sData, nil
}

func (m *mssqlRepository) GetForeignData() ([]entity.ForeignData, error) {

	db, _ := m.DBConn()
	defer db.Close()

	var SecurityDatas []entity.ForeignData

	rows, err1 := db.Query("select stkcode, stkid, DateOfRec, QtyAvail from mf_StkInfoBK")
	if err1 != nil {
		log.Fatal(err1)
	}
	for rows.Next() {
		var sData entity.ForeignData
		err2 := rows.Scan(&sData.SecurityName, &sData.SecurityId, &sData.DateOfRec, &sData.QtyAvail)
		if err2 != nil {
			return nil, err2
		}
		fmt.Println(sData)
		SecurityDatas = append(SecurityDatas, sData)
	}
	return SecurityDatas, nil
}

func (m *mssqlRepository) GetForeignDatabyStkCode(stk string) (*entity.ForeignData, error) {

	db, _ := m.DBConn()
	defer db.Close()

	var sData entity.ForeignData

	err1 := db.QueryRow("select stkcode, stkid, DateOfRec, QtyAvail from mf_StkInfoBK where stkcode = $1", stk).Scan(&sData.SecurityName, &sData.SecurityId, &sData.DateOfRec, &sData.QtyAvail)
	if err1 != nil {
		return nil, err1
	}

	return &sData, nil
}
