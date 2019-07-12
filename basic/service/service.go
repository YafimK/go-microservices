package service

import (
	"github.com/Yafimk/go-microservices/common"
)

type Service struct {
	server *common.WebServer
}

func (service *Service) Start(host string) {
	service.server = common.NewWebServer(host)
	service.server.RegisterRoutes(basicServiceRoutes)
	service.server.Start()
}
