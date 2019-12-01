package main

import (
	"bytes"
	"fmt"
)

type Logger struct{}

func (logger *Logger) Log(s string) {
	fmt.Println(s)
}

type HttpClient struct {
	logger *Logger
}

func (client *HttpClient) Get(url string) string {
	client.logger.Log("getting info from " + url)

	return "Get info from url " + url
}

func NewHttpClient(logger *Logger) *HttpClient {
	return &HttpClient{logger}
}

type ServiceGroup struct {
	logger     *Logger
	httpClient *HttpClient
}

func CreateServiceGroup(logger *Logger, httpClient *HttpClient) *ServiceGroup {
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

func main() {
	logger := &Logger{}
	httpClient := &HttpClient{logger}
	serviceGroup := CreateServiceGroup(logger, httpClient)
	result := serviceGroup.GetAll("https://google.com", "https://youtube.com")
	logger.Log(result)
}
