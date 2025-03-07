package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/rodrigc/fileserver/pb"
	server "github.com/rodrigc/fileserver/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	tls := false
	port := 8081
	certFile := ""
	keyFile := ""

	lis, err := net.Listen("tcp", fmt.Sprintf("0:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if tls {
		if certFile == "" {
			certFile = "x509/server_cert.pem"
		}
		if keyFile == "" {
			keyFile = "x509/server_key.pem"
		}
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials: %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterFileServiceServer(grpcServer, server.NewFileserver())
	grpcServer.Serve(lis)
}
