package postgres

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	"github.com/sugaml/authserver/internal/adapter/config"
)

func NewPostgres(config *config.DB) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Connection,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Cannot connect %s to  database: %v", config.Connection, err)
	}
	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
