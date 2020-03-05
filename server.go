package server

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

//Srv contains main server components eg. database and http service
type Srv struct {
	DB     *sql.DB
	HTTP   *http.Server
	Router *mux.Router
}

//KairosService as main Storage for all services
var KairosService Srv

//RegisterRoutes ... registers object routes
func RegisterRoutes(method string, endPoint string, f func(http.ResponseWriter, *http.Request)) {
	KairosService.Router.HandleFunc(endPoint, f).Methods(method)
}
