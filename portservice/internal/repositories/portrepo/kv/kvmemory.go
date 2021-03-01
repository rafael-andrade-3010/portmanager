package kv

import (
	"log"
	"portservice/internal/core/domain"
)

type repo struct {
	database map[string]*domain.Port
}

func New() *repo {
	return &repo{
		database: make(map[string]*domain.Port),
	}
}

func (r *repo) Save(ports []domain.Port) error {
	log.Printf("saving %v", len(ports))
	for _, p := range ports {
		r.database[p.Key] = &p
	}
	return nil
}

//start, limit int32
func (r *repo) GetAll() ([]domain.Port, error) {
	list := make([]domain.Port, 0)
	for _, v := range r.database {
		list = append(list, *v)
	}
	return list, nil
}
