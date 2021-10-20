package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"../entity"
)

func (*mssqlRepository) processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
func (m *mssqlRepository) LoadConfig(cfg *entity.DBConfig) {
	f, err := os.Open("F:\\Go\\checkImportedData\\config\\config.json")
	if err != nil {
		m.processError(err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		m.processError(err)
	}
}
