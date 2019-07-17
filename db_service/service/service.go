package service

import (
	. "github.com/Yafimk/go-microservices/common"
	"log"
)

type Service struct {
	server *WebServer
	*DbHandler
	bucketName string
}

func (service Service) Routes() Routes {
	return Routes{
		{
			Name:    "GetDocument",
			Method:  "GET",
			Pattern: "/document/{Id}",
			Handler: service.GetDocument(service.bucketName),
		},
		{
			Name:    "AddDocument",
			Method:  "POST",
			Pattern: "/documents/Add",
			Handler: service.AddDocument(service.bucketName),
		},
		{
			Name:    "HealthCheck",
			Method:  "GET",
			Pattern: "/HealthCheck",
			Handler: service.CheckDocumentServiceHealth(),
		},
	}

}

func NewService(host string, dbName string, bucketName string) *Service {
	dbClient, err := NewDbDriver(dbName)
	if err != nil {
		log.Fatalf(err.Error())
	}
	service := &Service{
		server: NewWebServer(host),
		DbHandler: &DbHandler{
			client: dbClient,
		},
		bucketName: bucketName,
	}

	service.server.RegisterRoutes(service.Routes())

	return service
}

func (service *Service) Start() {
	service.server.Start()
}
