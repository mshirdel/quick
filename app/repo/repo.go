package repo

import "gorm.io/gorm"

type Repo struct {
	database *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{database: db}
}

func Init() {
}
