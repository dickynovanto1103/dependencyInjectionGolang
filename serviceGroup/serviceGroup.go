package serviceGroup

import (
	"bytes"

	"github.com/dickynovanto1103/dependencyInjectionGolang/httpClient"
	"github.com/dickynovanto1103/dependencyInjectionGolang/logger"
)

type ServiceGroup struct {
	logger     *logger.Logger
	httpClient *httpClient.HttpClient
}

func NewServiceGroup(logger *logger.Logger, httpClient *httpClient.HttpClient) *ServiceGroup {
	return &ServiceGroup{logger, httpClient}
}

func (service *ServiceGroup) GetAll(urls ...string) string {
	service.logger.Log("Getting all from url")
	var result bytes.Buffer

	for _, val := range urls {
		result.WriteString(service.httpClient.Get(val))
	}
	return result.String()
}
