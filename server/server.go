package server

import (
    "context"
    "fmt"
    "log"
    "time"

    "go-grpc-report/proto"
)

var reports = make(map[string]string)

type ReportServer struct {
    proto.UnimplementedReportServiceServer
}

func (s *ReportServer) GenerateReport(ctx context.Context, req *proto.UserRequest) (*proto.ReportResponse, error) {
    reportID := fmt.Sprintf("report-%s-%d", req.UserId, time.Now().Unix())
    reports[req.UserId] = reportID

    log.Printf("[GenerateReport] UserID: %s, ReportID: %s", req.UserId, reportID)

    return &proto.ReportResponse{ReportId: reportID}, nil
}

func (s *ReportServer) HealthCheck(ctx context.Context, req *proto.HealthRequest) (*proto.HealthResponse, error) {
    log.Println("[HealthCheck] Service is healthy")
    return &proto.HealthResponse{Status: "SERVING"}, nil
}
