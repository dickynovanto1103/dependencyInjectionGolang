package main

import (
	"fmt"

	"github.com/dickynovanto1103/dependencyInjectionGolang/serviceGroup"
)

func main() {
	serviceGroup := serviceGroup.CreateServiceGroup()
	result := serviceGroup.GetAll("https://google.com", "https://youtube.com")

	fmt.Println("result", result)
}
