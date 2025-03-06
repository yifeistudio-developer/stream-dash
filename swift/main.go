package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "swift/proto"
	"time"
)

type server struct {
	pb.UnimplementedStreamServiceServer
}

func (s *server) StreamData(stream pb.StreamService_StreamDataServer) error {
	for {
		data := &pb.DataStream{
			Points: []*pb.DataPoint{
				{Timestamp: "2025-03-05", Value: 10},
				{Timestamp: "2025-03-06", Value: 20},
			},
		}
		if err := stream.Send(data); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStreamServiceServer(s, &server{})
	log.Printf("Swift running on :8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
