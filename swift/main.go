package main

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"google.golang.org/protobuf/proto"
	"log"
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

		p.Flush(1000) // 等待 1 秒确保消息发送
		time.Sleep(1 * time.Second)
	}
}
