package auth

import "github.com/api-monolith-template/internal/model/contract"

type Service struct {
	userRepository contract.UserRepository
}

func NewService() *Service {
	return new(Service)
}

func (s *Service) WithUserRepository(repo contract.UserRepository) *Service {
	s.userRepository = repo
	return s
}
