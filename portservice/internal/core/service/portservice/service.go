package portservice

import (
	"github.com/pkg/errors"
	"portservice/internal/core/domain"
	"portservice/internal/core/ports"
)

var (
	ErrKeyRequired = errors.Errorf("Key is required")
)

type service struct {
	repository ports.PortRepository
}

func New(repository ports.PortRepository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ports []domain.Port) error {
	for _, p := range ports {
		if p.Key == "" {
			return ErrKeyRequired
		}
	}
	return s.repository.Save(ports)
}

//start, limit int32
func (s *service) GetAll() ([]domain.Port, error) {
	return s.repository.GetAll()
}
