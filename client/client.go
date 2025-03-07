package main

import (
	"context"
	"crypto/tls"

	//"crypto/tls"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/rodrigc/fileserver/pb"
)

func main() {
	serverAddr := flag.String(
		"server", "localhost:8080",
		"The server address in the format of host:port",
	)
	file := flag.String(
		"file", "",
		"The file to query on the fileserver",
	)
	disableTLS := flag.Bool(
		"insecure",
		true,
		"Disable TLS",
	)
	flag.Parse()

	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})

	opts := []grpc.DialOption{}
	if *disableTLS {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.Fatalln("fail to create client:", err)
	}
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)

	res, err := client.GetFileMetaData(ctx, &pb.File{
		FileName: *file,
	})
	if err != nil {
		log.Fatalln("error sending request:", err)
	}

	fmt.Printf("result: %v\n", res)

	_, err = client.Exists(ctx, &pb.File{
		FileName: *file,
	})
	if err != nil {
		log.Fatalf("error sending request: %v", err)
	}
}
