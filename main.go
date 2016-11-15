package main

import (
	"fmt"

	"github.com/gwmoura/deployer/service"
)

func main() {
	fmt.Println("Hello Deployer")
	err := service.Run()
	if err != nil {
		fmt.Println(err)
	}
}
