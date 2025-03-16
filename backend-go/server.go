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

// Hash 2004
// Hash 3973
// Hash 4071
// Hash 6113
// Hash 9952
// Hash 7200
// Hash 8689
// Hash 5262
// Hash 9903
// Hash 6827
// Hash 2976
// Hash 7333
// Hash 2442
// Hash 1373
// Hash 9189
// Hash 9443
// Hash 4190
// Hash 8209
// Hash 3595
// Hash 2571
// Hash 5821
// Hash 3263
// Hash 4041
// Hash 7645
// Hash 6992
// Hash 1172
// Hash 3354
// Hash 5503
// Hash 6054
// Hash 4893
// Hash 3788
// Hash 2740
// Hash 2430
// Hash 6651
// Hash 3153
// Hash 9994
// Hash 7055
// Hash 2423
// Hash 2886
// Hash 1901
// Hash 7143
// Hash 8940
// Hash 9641
// Hash 2124
// Hash 7486
// Hash 1716
// Hash 4979
// Hash 2722
// Hash 1954
// Hash 4419
// Hash 2412
// Hash 2345
// Hash 9041
// Hash 1326
// Hash 3646