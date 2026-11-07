package repository

import "github.com/phil/user-service/internal/repository/profile"

// Registry gom các repository của service. Hiện lưu in-memory, đổi sang DB thì chỉ sửa tầng này.
type Registry interface {
	Profile() profile.Repository
}
