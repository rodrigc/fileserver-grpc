package server

import (
	"context"
	"os"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/rodrigc/fileserver/pb"
)

type fileServer struct {
	pb.UnimplementedFileServiceServer
}

func (f *fileServer) GetFileMetaData(_ context.Context, file *pb.File) (*pb.FileMetadata, error) {
	fileInfo, err := os.Stat(file.FileName)
	if err != nil {
		return nil, err
	}

	fileMetadata := &pb.FileMetadata{
		FileName: file.FileName,
		Size:     uint64(fileInfo.Size()),
		ModifiedTs: &timestamppb.Timestamp{
			Seconds: fileInfo.ModTime().Unix(),
		},
	}

	return fileMetadata, nil
}

func (f *fileServer) Exists(_ context.Context, file *pb.File) (*emptypb.Empty, error) {
	_, err := os.Stat(file.FileName)
	return nil, err
}

func NewFileserver() *fileServer {
	return &fileServer{}
}
