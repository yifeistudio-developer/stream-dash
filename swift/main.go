package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"net/http"
	pb "swift/proto"
	"time"
)

type server struct {
	pb.UnimplementedStreamServiceServer
}

func (s *server) StreamData(grpc.BidiStreamingServer[pb.DataStream, pb.DataStream]) error {
	fmt.Println("StreamData")
	return status.Errorf(codes.Unimplemented, "method StreamData not implemented")
}

func startKafkaProducer() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "192.168.18.149:9092"})
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer p.Close()

	topic := "input-topic"

	for {
		data := &pb.DataStream{
			Points: []*pb.DataPoint{
				{Timestamp: "2025-03-05", Value: 10},
				{Timestamp: "2025-03-06", Value: 20},
			},
		}
		payload, err := proto.Marshal(data)
		if err != nil {
			log.Printf("Failed to marshal: %v", err)
			continue
		}

		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          payload,
		}, nil)
		if err != nil {
			log.Printf("Failed to produce: %v", err)
		}

		p.Flush(1000)
		time.Sleep(1 * time.Second)
	}
}

func startGRPCServer() {
	// 创建 gRPC 服务
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterStreamServiceServer(grpcServer, &server{})

	// 包装为 gRPC-Web
	grpcWebServer := grpcweb.WrapServer(grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true // 允许所有来源（开发用）
		}),
	)

	// 创建 HTTP 服务器，支持 gRPC 和 gRPC-Web
	httpServer := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if grpcWebServer.IsGrpcWebRequest(r) {
				grpcWebServer.ServeHTTP(w, r)
			} else {
				grpcServer.ServeHTTP(w, r) // 原生 gRPC
			}
		}),
	}

	log.Printf("Swift gRPC and gRPC-Web running on :8080")
	if err := httpServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func main() {
	go startKafkaProducer()
	startGRPCServer()
}
