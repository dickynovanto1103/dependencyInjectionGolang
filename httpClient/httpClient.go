package httpClient

import (
	"fmt"

	"github.com/dickynovanto1103/dependencyInjectionGolang/logger"
)

type HttpClient struct {
	logger *logger.Logger
}

func (client *HttpClient) Get(url string) string {
	fmt.Println(client.logger)
	client.logger.Log("getting info from " + url)

	return "Get info from url " + url
}

func NewHttpClient(logger *logger.Logger) *HttpClient {
	return &HttpClient{logger}
}
