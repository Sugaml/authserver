package service

import (
	"context"

	"github.com/sugaml/authserver/internal/adapter/storage/postgres/repository"
	"github.com/sugaml/authserver/internal/core/domain"
	"github.com/sugaml/authserver/internal/core/port"
	"github.com/sugaml/authserver/internal/core/util"
)

type UserServiceGetter interface {
	User() port.UserService
}

type UserService struct {
	repo repository.IRepository
}

func newUserService(repo repository.IRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// Register creates a new user
func (us *UserService) Register(ctx context.Context, user *domain.RegisterRequest) (*domain.UserResponse, error) {
	data := domain.Convert[domain.RegisterRequest, domain.User](user)
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, domain.ErrInternal
	}
	user.Password = hashedPassword
	_, err = us.repo.User().GetByEmail(ctx, user.Email)
	if err == nil {
		return nil, err
	}
	result, err := us.repo.User().Create(ctx, data)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	return domain.Convert[domain.User, domain.UserResponse](result), nil
}

// Get gets a user by ID
func (us *UserService) Get(ctx context.Context, id uint64) (*domain.UserResponse, error) {
	result, err := us.repo.User().GetByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	return domain.Convert[domain.User, domain.UserResponse](result), nil
}

// List lists all users
func (us *UserService) List(ctx context.Context, skip, limit uint64) ([]*domain.UserResponse, error) {
	var userResponse []*domain.UserResponse

	users, err := us.repo.User().List(ctx, skip, limit)
	if err != nil {
		return nil, domain.ErrInternal
	}
	for _, user := range users {
		userResponse = append(userResponse, domain.Convert[domain.User, domain.UserResponse](user))
	}

	return userResponse, nil
}

// Update updates a user's name, email, and password
func (us *UserService) Update(ctx context.Context, user *domain.User) (*domain.UserResponse, error) {
	_, err := us.repo.User().GetByID(ctx, uint64(user.ID))
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	var hashedPassword string

	if user.Password != "" {
		hashedPassword, err = util.HashPassword(user.Password)
		if err != nil {
			return nil, domain.ErrInternal
		}
	}

	user.Password = hashedPassword

	result, err := us.repo.User().Update(ctx, user)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return domain.Convert[domain.User, domain.UserResponse](result), nil
}

// Delete deletes a user by ID
func (us *UserService) Delete(ctx context.Context, id uint64) error {
	_, err := us.repo.User().GetByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return err
		}
		return domain.ErrInternal
	}

	return us.repo.User().Delete(ctx, id)
}
