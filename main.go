package main

import (
	"fmt"
)

func main() {
	serviceGroup := CreateServiceGroup()
	result := serviceGroup.GetAll("https://google.com", "https://youtube.com")

	fmt.Println("result", result)
}
