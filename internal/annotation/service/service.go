package annotation_service

import "github.com/Flak34/crowd-api/internal/entrypoint"

type Service struct {
	ep entrypoint.Entrypoint
}

func New(ep entrypoint.Entrypoint) *Service {
	return &Service{
		ep: ep,
	}
}
