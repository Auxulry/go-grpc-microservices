// Package gateway is describe reusable package for create gateway server
package gateway

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

var (
	MaxHeaderBytes = 1 << 20
	ReadTimeOut    = 10 * time.Second
	WriteTimeOut   = 10 * time.Second
)

type RestServer struct {
	*runtime.ServeMux
	Addr           string
	MaxHeaderBytes int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}

func (rest *RestServer) Run(ctx context.Context) error {
	gwServer := &http.Server{
		Addr:           fmt.Sprintf(":%v", rest.Addr),
		Handler:        rest.ServeMux,
		ReadTimeout:    rest.ReadTimeout,
		WriteTimeout:   rest.WriteTimeout,
		MaxHeaderBytes: rest.MaxHeaderBytes,
	}

	go func() {
		<-ctx.Done()
		log.Println("Shutting down the http gateway server")
		if err := gwServer.Shutdown(ctx); err != nil {
			log.Fatalf("Failed to shutdown http gateway server: %v", err)
		}
	}()
	if err := gwServer.ListenAndServe(); err != nil {
		log.Fatalln("Server gRPC-Gateway exited with error:", err)
	}

	return nil
}

func NewGateway(addr string) *RestServer {
	gwmux := runtime.NewServeMux()

	return &RestServer{
		ServeMux:       gwmux,
		Addr:           addr,
		MaxHeaderBytes: MaxHeaderBytes,
		ReadTimeout:    ReadTimeOut,
		WriteTimeout:   WriteTimeOut,
	}
}
