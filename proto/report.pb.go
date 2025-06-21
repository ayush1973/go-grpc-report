package proto

import ()

type UserRequest struct {
	UserId string
}

type ReportResponse struct {
	ReportId string
}

type HealthRequest struct{}

type HealthResponse struct {
	Status string
}
