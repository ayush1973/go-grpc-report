package main

import (
    "context"
    "log"
    "time"

    "go-grpc-report/proto"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    client := proto.NewReportServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // Call GenerateReport
    res, err := client.GenerateReport(ctx, &proto.UserRequest{UserId: "ayush-client"})
    if err != nil {
        log.Fatalf("GenerateReport error: %v", err)
    }
    log.Printf("Generated Report ID: %s", res.ReportId)

    // Call HealthCheck
    health, err := client.HealthCheck(ctx, &proto.HealthRequest{})
    if err != nil {
        log.Fatalf("HealthCheck error: %v", err)
    }
    log.Printf("Health Status: %s", health.Status)
}
