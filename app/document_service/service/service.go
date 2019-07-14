package service

import (
	"github.com/Yafimk/go-microservices/document-service/common"
	"log"
)

const BucketName = "documents"
const DbName = "db"

type Service struct {
	server           *common.WebServer
	serviceDbHandler *DbHandler
}

func (service Service) Routes() common.Routes {
	return common.Routes{
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
	dbClient, err := common.NewDbDriver(DbName)
	if err != nil {
		log.Fatalf(err.Error())
	}
	service := &Service{
		common.NewWebServer(host),
		&DbHandler{
			dbClient,
		},
	}

	service.server.RegisterRoutes(service.Routes())

	return service
}

func (service *Service) Start() {
	service.server.Start()
}
