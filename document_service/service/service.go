package service

import (
	. "github.com/Yafimk/go-microservices/common"
	"log"
)

const BucketName = "documents"
const DbName = "db"

type Service struct {
	server           *WebServer
	serviceDbHandler *DbHandler
}

func (service Service) Routes() Routes {
	return Routes{
		{
			Name:    "GetDocument",
			Method:  "GET",
			Pattern: "/documents/{Id}",
			Handler: service.serviceDbHandler.GetDocument(BucketName),
		},
		{
			Name:    "HealthCheck",
			Method:  "GET",
			Pattern: "/HealthCheck",
			Handler: service.serviceDbHandler.CheckDocumentServiceHealth(),
		},
	}

}

func NewService(host string) *Service {
	dbClient, err := NewDbDriver(DbName)
	if err != nil {
		log.Fatalf(err.Error())
	}
	service := &Service{
		server: NewWebServer(host),
		serviceDbHandler: &DbHandler{
			client: dbClient,
		},
	}

	service.server.RegisterRoutes(service.Routes())

	return service
}

func (service *Service) Start() {
	service.server.Start()
}
