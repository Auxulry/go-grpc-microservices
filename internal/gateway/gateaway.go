// Package gateway is described reusable package for create gateway server
package gateway

import (
	"context"
	"fmt"
	"github.com/MochamadAkbar/go-grpc-microservices/pkg/middleware"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"strings"
	"time"

	"github.com/MochamadAkbar/go-grpc-microservices/api"
	"github.com/MochamadAkbar/go-grpc-microservices/third_party"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

var (
	MaxHeaderBytes = 1 << 20
	ReadTimeOut    = 10 * time.Second
	WriteTimeOut   = 10 * time.Second
)

type Gateway struct {
	*runtime.ServeMux
	Addr           string
	MaxHeaderBytes int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}

func NewGateway(addr string, opts ...runtime.ServeMuxOption) *Gateway {
	gwMux := runtime.NewServeMux(opts...)

	return &Gateway{
		ServeMux:       gwMux,
		Addr:           addr,
		MaxHeaderBytes: MaxHeaderBytes,
		ReadTimeout:    ReadTimeOut,
		WriteTimeout:   WriteTimeOut,
	}
}

func swaggerUIHandler() http.Handler {
	err := mime.AddExtensionType(".svg", "image/svg+xml")
	if err != nil {
		panic(err)
	}
	subFS, err := fs.Sub(third_party.SwaggerUI, "swagger-ui")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}
	return http.FileServer(http.FS(subFS))
}

func (gw *Gateway) Run(ctx context.Context) error {
	sw := swaggerUIHandler()

	fileServer := http.FileServer(http.FS(api.FS))
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.Handle("/", sw)

	gwServer := &http.Server{
		Addr: fmt.Sprintf(":%v", gw.Addr),
		Handler: middleware.CORS(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if strings.HasPrefix(request.URL.Path, "/api/v1") {
				gw.ServeMux.ServeHTTP(writer, request)
				return
			}
			mux.ServeHTTP(writer, request)
		})),
		ReadTimeout:    gw.ReadTimeout,
		WriteTimeout:   gw.WriteTimeout,
		MaxHeaderBytes: gw.MaxHeaderBytes,
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
