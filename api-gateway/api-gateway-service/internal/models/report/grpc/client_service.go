package grpc

import (
	"context"

	"github.com/mikemonzo/play-microservices/api-gateway/api-gateway-service/config"
	gr "github.com/mikemonzo/play-microservices/api-gateway/api-gateway-service/pkg/grpc"
	"github.com/mikemonzo/play-microservices/api-gateway/api-gateway-service/pkg/logger"
	"github.com/mikemonzo/play-microservices/api-gateway/api-gateway-service/proto"
)

type ReportGRPCClient struct {
	log logger.Logger
	cfg *config.Config
	gr.GRPC_Client
}

func NewReportGRPCClient(log logger.Logger, cfg *config.Config) *ReportGRPCClient {
	return &ReportGRPCClient{log: log, cfg: cfg}
}

func (jc *ReportGRPCClient) GRPC_ListReports(c context.Context, listReportsRequest *proto.ListReportsRequest) (*proto.ListReportResponse, error) {
	jc.log.Info("ReportGRPCClient.GRPC_ListReports: Connecting to grpc server...")
	conn, err := jc.Connect(jc.cfg.ReportServiceURL)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto.NewReportServiceClient(conn)
	jc.log.Info("ReportGRPCClient.GRPC_ListReports: Conneced to grpc server...")
	jc.log.Infof("ReportGRPCClient.GRPC_ListReports: calling server for ListReports: %v", listReportsRequest)
	return client.ListReports(c, listReportsRequest)
}