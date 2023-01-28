// Package cmd is described Main applications for this project.
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/MochamadAkbar/go-grpc-microservices/pkg/gateway"
	"github.com/MochamadAkbar/go-grpc-microservices/pkg/rpcclient"
	"github.com/MochamadAkbar/go-grpc-microservices/pkg/rpcserver"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	svPort  string
	gwPort  string
	rootCmd = &cobra.Command{
		Use:   "service",
		Short: "Running the gRPC service",
		Long:  "Used to run gRPC Service including rpc server, rpc client and gateway",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()

			rpcServer := rpcserver.NewRPCServer(svPort, "tcp", false)
			defer rpcServer.StopListener()

			log.Println("Serving gRPC on ", svPort)
			if err := rpcServer.Run(); err != nil {
				log.Fatalln("Failed to listen grpc server")
			}

			_, err := rpcclient.NewRPCClient(ctx, svPort, false, grpc.WithBlock())
			fmt.Println("Dial Server")
			if err != nil {
				log.Fatalln("Failed to dial server:", err)
			}

			rpcGateway := gateway.NewGateway(gwPort)
			log.Println("Serving gRPC-Gateway on http://0.0.0.0:5001")
			if err := rpcGateway.Run(ctx); err != nil {
				log.Fatalln("Failed to listen grpc server")
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
