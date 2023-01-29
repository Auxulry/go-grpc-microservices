package main

import (
	"fmt"
	"github.com/MochamadAkbar/go-grpc-microservices/cmd"
	"github.com/MochamadAkbar/go-grpc-microservices/configs"
)

func init() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)
}

func main() {
	cmd.Execute()
}
