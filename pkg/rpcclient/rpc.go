package rpcclient

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func NewRPCClient(ctx context.Context, addr string, tls bool, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	if tls {
		certFile := "ssl/certificates/ca.crt" // => file path location your certFile
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
			return &grpc.ClientConn{}, err
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		creds := grpc.WithTransportCredentials(insecure.NewCredentials())
		opts = append(opts, creds)
	}

	opts = append(opts, grpc.WithChainUnaryInterceptor(LogInterceptor()))

	conn, err := grpc.DialContext(ctx, fmt.Sprintf(":%v", addr), opts...)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
		return &grpc.ClientConn{}, err
	}

	return conn, nil
}
