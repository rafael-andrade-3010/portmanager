package ports

import "portservice/internal/core/domain"

type PortRepository interface {
	GetAll() ([]domain.Port, error)
	Save(ports []domain.Port) error
}
