package controller

import (
	"../repository"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var repo repository.MssqlRepository = repository.NewMssqlRepository()

type MssqlController interface {
	GetSecurity(res http.ResponseWriter, req *http.Request)
	GetSecuritybyStkCode(res http.ResponseWriter, req *http.Request)
	GetForeign(res http.ResponseWriter, req *http.Request)
	GetForeignbyStkCode(res http.ResponseWriter, req *http.Request)
}

type mssqlController struct{}

func NewMssqlController() MssqlController {
	return &mssqlController{}
}

func (*mssqlController) GetSecurity(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	post, _ := repo.GetSecurityData()
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}

func (*mssqlController) GetSecuritybyStkCode(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	res.Header().Set("Content-Type", "application/json")
	post, _ := repo.GetSecurityDatabyStkCode(string(vars["stkcode"]))
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}

func (*mssqlController) GetForeign(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	post, _ := repo.GetForeignData()
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}

func (*mssqlController) GetForeignbyStkCode(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	res.Header().Set("Content-Type", "application/json")
	post, _ := repo.GetForeignDatabyStkCode(string(vars["stkcode"]))
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}
