package repository

// import (
// 	"context"

// 	"github.com/jinzhu/gorm"
// 	"github.com/sugaml/authserver/internal/core/domain"
// 	"github.com/sugaml/authserver/internal/core/port"
// )

// type OauthGetter interface {
// 	Client() port.ClientRepository
// }

// type OauthRepository struct {
// 	db *gorm.DB
// }

// func newOauthepository(db *gorm.DB) *OauthRepository {
// 	return &OauthRepository{
// 		db: db,
// 	}
// }

// func (r *OauthRepository) Get(ctx context.Context, id string) (*domain.Client, error) {
// 	var data domain.Client
// 	if err := r.db.Model(&domain.Client{}).
// 		Take(&data, "id = ?", id).Error; err != nil {
// 		return nil, err
// 	}
// 	return &data, nil
// }
