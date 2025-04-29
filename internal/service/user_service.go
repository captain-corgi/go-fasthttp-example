package service

import (
	"github.com/captain-corgi/go-fasthttp-example/internal/domain/model"
	"github.com/captain-corgi/go-fasthttp-example/internal/domain/repository"
)

// UserService handles business logic for user operations
type UserService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id string) (*model.User, error) {
	return s.userRepo.GetByID(id)
}

// CreateUser creates a new user
func (s *UserService) CreateUser(user *model.User) error {
	return s.userRepo.Create(user)
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(user *model.User) error {
	return s.userRepo.Update(user)
}

// DeleteUser removes a user
func (s *UserService) DeleteUser(id string) error {
	return s.userRepo.Delete(id)
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers() ([]*model.User, error) {
	return s.userRepo.GetAll()
}
