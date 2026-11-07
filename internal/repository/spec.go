package repository

import "github.com/phillong97254/user-service/internal/repository/profile"

// Registry gom các repository của service (Postgres qua gorm).
type Registry interface {
	Profile() profile.Repository
}
