package repository

import (
	"errors"

	"gorm.io/gorm"

	"go-devops/internal/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetAll lấy tất cả người dùng
func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)
	return users, result.Error
}

// GetByID lấy người dùng theo ID
func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Create tạo người dùng mới
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// Update cập nhật thông tin người dùng
func (r *UserRepository) Update(user *models.User) error {
	// Kiểm tra người dùng có tồn tại không
	var existingUser models.User
	result := r.db.First(&existingUser, user.ID)
	if result.Error != nil {
		return errors.New("Không tìm thấy người dùng")
	}

	// Cập nhật thông tin
	return r.db.Save(user).Error
}

// Delete xóa người dùng
func (r *UserRepository) Delete(id uint) error {
	result := r.db.Delete(&models.User{}, id)
	if result.RowsAffected == 0 {
		return errors.New("Không tìm thấy người dùng")
	}
	return result.Error
}
