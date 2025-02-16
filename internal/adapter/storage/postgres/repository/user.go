package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/sugaml/authserver/internal/core/domain"
	"github.com/sugaml/authserver/internal/core/port"
)

type UserGetter interface {
	User() port.UserRepository
}

type UserRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create creates a new user in the database
func (r *UserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := r.db.Model(&domain.User{}).Create(user).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

// GetByID gets a user by ID from the database
func (r *UserRepository) GetByID(ctx context.Context, id uint64) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.Model(domain.User{}).Where("id = ? and is_active = true", id).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

// GetByEmailAndPassword gets a user by email from the database
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.Model(domain.User{}).Where("email = ?", email).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r *UserRepository) GetByMobileNum(ctx context.Context, mobileNum string) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.Model(domain.User{}).Where("mobile_num = ?", mobileNum).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r *UserRepository) SetPassword(ctx context.Context, id uint64, password string) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.Model(domain.User{}).Where("id = ?", id).UpdateColumns(
		map[string]interface{}{
			"password": password,
		},
	).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

// List lists all users from the database
func (r *UserRepository) List(ctx context.Context, skip, limit uint64) ([]*domain.User, error) {
	users := []*domain.User{}
	err := r.db.Model(&domain.User{}).Order("id desc").Find(users).Error
	if err != nil {
		return nil, err
	}
	return users, err
}

// Update updates a user by ID in the database
func (r *UserRepository) Update(ctx context.Context, data *domain.User) (*domain.User, error) {
	user := &domain.User{}
	if data.Email != "" {
		user.Email = data.Email
	}
	err := r.db.Model(&domain.User{}).Where("id = ?", data.ID).Updates(user).Error
	if err != nil {
		return &domain.User{}, err
	}
	return user, nil
}

// DeleteUser deletes a user by ID from the database
func (r *UserRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.Model(&domain.User{}).Where("id = ?", id).Delete(&domain.User{}).Error
}
