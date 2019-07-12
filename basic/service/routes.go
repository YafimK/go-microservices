package service

import (
	"github.com/Yafimk/go-microservices/common"
	"net/http"
)

var basicServiceRoutes = common.Routes{
	{
		"GetBasicInfo",
		"GET",
		"/info/{Id}",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write([]byte("{\"result\":\"OK\"}"))
		},
	},
}
