package main

import (
    "log"
    "net"

    "go-grpc-report/server"
    "go-grpc-report/proto"

    "github.com/robfig/cron/v3"
    "google.golang.org/grpc"
)

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    reportServer := &server.ReportServer{}
    proto.RegisterReportServiceServer(grpcServer, reportServer)

    c := cron.New()
    c.AddFunc("@every 10s", func() {
        for _, userID := range []string{"user1", "user2", "user3"} {
            _, err := reportServer.GenerateReport(nil, &proto.UserRequest{UserId: userID})
            if err != nil {
                log.Printf("Error generating report: %v", err)
            }
        }
    })
    c.Start()

    log.Println("gRPC server started on port :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
