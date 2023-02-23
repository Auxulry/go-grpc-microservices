// Package cmd is described Main applications for this project.
package cmd

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"

	"github.com/MochamadAkbar/go-grpc-microservices/configs"
	"github.com/MochamadAkbar/go-grpc-microservices/domain"
	"github.com/MochamadAkbar/go-grpc-microservices/internal/gateway"
	"github.com/MochamadAkbar/go-grpc-microservices/pkg/orm"
	"github.com/MochamadAkbar/go-grpc-microservices/pkg/rpcclient"
	"github.com/MochamadAkbar/go-grpc-microservices/pkg/rpcserver"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"gorm.io/gorm"
)

var (
	svPort  string
	gwPort  string
	rootCmd = &cobra.Command{
		Use:   "service",
		Short: "Running the gRPC service",
		Long:  "Used to run gRPC Service including rpc server, rpc client and gateway",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logrus.New()
			ctx := context.Background()

			cfg, err := configs.LoadConfig(".")
			if err != nil {
				panic(err)
			}

			db, err := orm.NewPSQL(ctx,
				cfg.DBDNS,
				&orm.ConfigConnProvider{
					ConnMaxLifetime: time.Hour,
					ConnMaxIdleTime: time.Minute,
					MaxOpenConns:    10,
					MaxIdleConns:    10,
				}, &gorm.Config{})
			if err != nil {
				panic(err)
			}

			rpcServer := rpcserver.NewRPCServer(svPort, "tcp", false)
			defer rpcServer.StopListener()

			domain.RegisterAuthServiceServer(ctx, db, rpcServer.Server)
			grpc_health_v1.RegisterHealthServer(rpcServer.Server, health.NewServer())
			logger.Infoln("Serving gRPC on", svPort)
			if err := rpcServer.Run(); err != nil {
				logger.Fatalln("Failed to listen grpc server")
			}

			rpcClient, err := rpcclient.NewRPCClient(ctx, svPort, false, grpc.WithBlock())
			if err != nil {
				logger.Fatalln("failed to dial server:", err)
				log.Fatalln("Failed to dial server:", err)
			}

			rpcGateway := gateway.NewGateway(gwPort)
			domain.RegisterAuthServiceHandler(ctx, rpcGateway.ServeMux, rpcClient)
			logger.Infoln("Serving gRPC-Gateway on", gwPort)
			if err := rpcGateway.Run(ctx); err != nil {
				logger.Fatalln("Failed to listen grpc server")
			}

			rpcServer.Terminate(ctx)
		},
	}
)

func Execute() {
	rootCmd.Flags().StringVarP(&svPort, "svport", "s", "", "define rpc server port")
	rootCmd.Flags().StringVarP(&gwPort, "gwport", "g", "", "define gateway port")
	rootCmd.MarkFlagsRequiredTogether("svport", "gwport")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
