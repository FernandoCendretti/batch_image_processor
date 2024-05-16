package image

import "batchImageProcessing/internal/event"

type Service struct {
	e event.Event
}

func NewService(e event.Event) *Service {
	return &Service{
		e: e,
	}
}
