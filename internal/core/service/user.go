package service

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/sugaml/authserver/internal/core/domain"
	"github.com/sugaml/authserver/internal/core/util"
	"golang.org/x/crypto/bcrypt"
)

// Register creates a new user
func (us *Service) RegisterUser(ctx context.Context, req *domain.RegisterRequest) (*domain.UserResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	data := domain.Convert[domain.RegisterRequest, domain.User](req)
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, domain.ErrInternal
	}
	data.Password = hashedPassword
	_, err = us.repo.User().GetByEmail(ctx, data.Email)
	if err == nil {
		return nil, errors.New("email already exists")
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
func (us *Service) GetUser(ctx context.Context, id uint64) (*domain.UserResponse, error) {
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
func (us *Service) ListUser(ctx context.Context, skip, limit uint64) ([]*domain.UserResponse, error) {
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
func (us *Service) UpdateUser(ctx context.Context, user *domain.User) (*domain.UserResponse, error) {
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
func (us *Service) DeleteUser(ctx context.Context, id uint64) error {
	_, err := us.repo.User().GetByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return err
		}
		return domain.ErrInternal
	}
	return us.repo.User().Delete(ctx, id)
}

func (s *Service) LoginUser(ctx context.Context, req *domain.LoginRequest) (*domain.UserResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	result, err := s.repo.User().GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	err = util.VerifyPassword(result.Password, req.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, err
	}
	logrus.Info("Loggedin user id :: ", result.ID)
	return domain.Convert[domain.User, domain.UserResponse](result), nil
}
