package ports

import "portservice/internal/core/domain"

type PortService interface {
	GetAll() ([]domain.Port, error)
	Create([]domain.Port) error
}
