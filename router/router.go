package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type MuxRouter interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERV(port string)
}

type muxRouter struct{}

var muxDispatcher = mux.NewRouter()

func NewMuxRouter() MuxRouter {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) SERV(port string) {
	fmt.Println("MuxDispatcher's running on port:", port)
	http.ListenAndServe(port, muxDispatcher)
}