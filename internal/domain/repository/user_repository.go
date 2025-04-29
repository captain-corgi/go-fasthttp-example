package repository

import "github.com/captain-corgi/go-fasthttp-example/internal/domain/model"

// UserRepository defines the interface for user data operations
type UserRepository interface {
	GetByID(id string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id string) error
}

// InMemoryUserRepository implements UserRepository interface with in-memory storage
type InMemoryUserRepository struct {
	users map[string]*model.User
}

// NewInMemoryUserRepository creates a new instance of InMemoryUserRepository
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*model.User),
	}
}

// GetByID retrieves a user by ID
func (r *InMemoryUserRepository) GetByID(id string) (*model.User, error) {
	if user, exists := r.users[id]; exists {
		return user, nil
	}
	return nil, nil
}

// Create adds a new user
func (r *InMemoryUserRepository) Create(user *model.User) error {
	r.users[user.ID] = user
	return nil
}

// Update modifies an existing user
func (r *InMemoryUserRepository) Update(user *model.User) error {
	r.users[user.ID] = user
	return nil
}

// Delete removes a user
func (r *InMemoryUserRepository) Delete(id string) error {
	delete(r.users, id)
	return nil
}
