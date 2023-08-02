package server

import (
    "log"
    "net"

    "github.com/mikemonzo/play-microservices/scheduler/scheduler-service/config"
	"github.com/mikemonzo/play-microservices/scheduler/scheduler-service/pkg/logger"
    MyJobGRPCService "github.com/mikemonzo/play-microservices/scheduler/scheduler-service/internal/models/job/grpc"
    JobGRPCServiceProto "github.com/mikemonzo/play-microservices/scheduler/scheduler-service/proto"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

type server struct {
    log       logger.Logger
    cfg       *config.Config
}

// NewServer constructor
func NewServer(log logger.Logger, cfg *config.Config) *server {
    return &server{log: log, cfg: cfg}
}

func (s *server) Run() error {
    lis, err := net.Listen("tcp", ":"+s.cfg.ServerPort)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpc_server := grpc.NewServer()

    job_service := MyJobGRPCService.NewJobService()
    JobGRPCServiceProto.RegisterJobServiceServer(grpc_server, job_service)
    reflection.Register(grpc_server)

    log.Printf("server listening at %v", lis.Addr())
    if err := grpc_server.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
        return err
    }

    return nil
}