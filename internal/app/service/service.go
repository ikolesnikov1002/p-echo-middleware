package service

import (
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int64 {
	futureDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	currentDate := time.Now()
	dateDiff := futureDate.Sub(currentDate)

	return int64(dateDiff.Hours() / 24)
}
