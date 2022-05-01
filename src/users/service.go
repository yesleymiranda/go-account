package users

import (
	"context"
)

type service struct {
	repository *repository
}

type Service interface {
	GetByID(ctx context.Context, id uint) (*User, error)
}

func NewService(r *repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetByID(ctx context.Context, id uint) (*User, error) {
	res, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
