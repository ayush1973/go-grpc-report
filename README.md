#  go-grpc-report

This project implements a gRPC-based report generation service in Go. It periodically generates reports for predefined users using a cron job.

##  Features

- `GenerateReport(UserID)` – Generates a report with timestamped ID
- `HealthCheck()` – gRPC health endpoint
- Cron job runs every 10 seconds for `user1`, `user2`, `user3`
- In-memory storage (map)
- Logging with timestamps

##  How to Run

### 1. Clone or unzip

```bash
cd go-grpc-report
go mod tidy
go run main.go
