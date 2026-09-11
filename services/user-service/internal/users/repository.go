package user

import (
	"context"
	"services/user-service/internal/model"

	"gorm.io/gorm"
)

// Interface as repository contract
type IUserRepository interface {
	FindById(ctx context.Context, id uint) (*model.User, error)
	ExistsById(ctx context.Context, id uint) (bool, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	ExistsByUsername(ctx context.Context, username string) (bool, error)
	// TODO - FindBySpokenDialects(ctx context.Context, spoken_dialects []string) (*[]model.User, error)
	// TODO - FindBySpokenLanguages(ctx context.Context, spoken_languages []string) (*[]model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint) error
}

// Database connection
type userRepository struct {
	db *gorm.DB
}

// Constructor to initialize repository + Dependency injection of database connection within it
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

// Implementation
func (r *userRepository) FindById(ctx context.Context, id uint) (*model.User, error) {
	var user model.User

	if err := r.db.WithContext(ctx).Model(&model.User{}).First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) ExistsById(ctx context.Context, id uint) (bool, error) {
	var count int64

	if err := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Limit(1).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var user *model.User

	if err := r.db.WithContext(ctx).Model(&model.User{}).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	var count int64

	if err := r.db.WithContext(ctx).Model(&model.User{}).Where("username = ?", username).Limit(1).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *userRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	if err := r.db.WithContext(ctx).Model(&model.User{}).Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	if err := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&model.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
