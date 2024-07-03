package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 7433
// Hash 1426
// Hash 5361
// Hash 7926
// Hash 2783
// Hash 3476
// Hash 3666
// Hash 4742
// Hash 6520
// Hash 5260
// Hash 9301
// Hash 4619
// Hash 3547
// Hash 8388
// Hash 9574
// Hash 9113
// Hash 9459
// Hash 2202
// Hash 2500
// Hash 6611
// Hash 7041
// Hash 7332
// Hash 6544
// Hash 2762
// Hash 6293