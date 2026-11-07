package profile

import "github.com/phil/user-service/internal/domain"

type impl struct {
	profiles map[string]domain.Profile
}

func New() Repository {
	return &impl{profiles: seed()}
}

// Dữ liệu demo. Thực tế: lưu DB.
func seed() map[string]domain.Profile {
	return map[string]domain.Profile{
		"alice": {ID: "alice", Name: "Alice Nguyen", Email: "alice@example.com", Department: "Finance Ops"},
		"bob":   {ID: "bob", Name: "Bob Tran", Email: "bob@example.com", Department: "Treasury"},
		"carol": {ID: "carol", Name: "Carol Le", Email: "carol@example.com", Department: "Support"},
	}
}
