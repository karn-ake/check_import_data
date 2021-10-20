package main

import (
	"./controller"
	"./router"
	"fmt"
	"net/http"
)

var (
	muxRouter router.MuxRouter           = router.NewMuxRouter()
	msctrl    controller.MssqlController = controller.NewMssqlController()
)

func main() {
	const port string = ":8080"
	muxRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and Running")
	})
	muxRouter.GET("/api/securityupdated", msctrl.GetSecurity)
	muxRouter.GET("/api/securityupdatedbystkcode/{stkcode}", msctrl.GetSecuritybyStkCode)
	muxRouter.SERV(port)
}
