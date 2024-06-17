package auth

import "github.com/api-monolith-template/internal/model/contract"

type Service struct {
	userRepository  contract.UserRepository
	cacheRepository contract.CacheRepository
}

func NewService() *Service {
	return new(Service)
}

func (s *Service) WithUserRepository(repo contract.UserRepository) *Service {
	s.userRepository = repo
	return s
}

func (s *Service) WithCacheRepository(repo contract.CacheRepository) *Service {
	s.cacheRepository = repo
	return s
}
