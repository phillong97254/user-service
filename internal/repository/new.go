package repository

import (
	"sync"

	"gorm.io/gorm"

	"github.com/phillong97254/user-service/internal/repository/profile"
)

type impl struct {
	db *gorm.DB

	profile     profile.Repository
	profileOnce sync.Once
}

func New(db *gorm.DB) Registry {
	return &impl{db: db}
}

func (i *impl) Profile() profile.Repository {
	i.profileOnce.Do(func() { i.profile = profile.New(i.db) })
	return i.profile
}
