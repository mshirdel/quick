package db

import (
	"fmt"

	"github.com/mshirdel/nebula/db/postgres"
	"github.com/mshirdel/quick/config"
	"gorm.io/gorm"
)

type DB struct {
	cfg      *config.Config
	Database *gorm.DB
}

func New(cfg *config.Config) *DB {
	return &DB{
		cfg: cfg,
	}
}

func (db *DB) Init() (err error) {
	if db.Database != nil {
		return err
	}

	db.Database, err = postgres.NewPostgres(&db.cfg.Database)
	if err != nil {
		return fmt.Errorf("connect to database : %w", err)
	}

	return err
}

func (db *DB) Close() {
	if db.Database != nil {
		postgres.Close(db.Database)
	}
}
