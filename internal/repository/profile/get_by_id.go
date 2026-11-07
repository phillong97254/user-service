package profile

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/phillong97254/user-service/internal/domain"
)

func (i *impl) GetByID(ctx context.Context, id string) (domain.Profile, error) {
	var p domain.Profile
	err := i.db.WithContext(ctx).Where("id = ?", id).Take(&p).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.Profile{}, ErrNotFound
	}
	return p, err
}
