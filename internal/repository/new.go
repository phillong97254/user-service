package repository

import (
	"sync"

	"github.com/phillong97254/user-service/internal/repository/profile"
)

type impl struct {
	profile     profile.Repository
	profileOnce sync.Once
}

func New() Registry {
	return &impl{}
}

func (i *impl) Profile() profile.Repository {
	i.profileOnce.Do(func() { i.profile = profile.New() })
	return i.profile
}
