package users

import (
	"gorm.io/gorm"
)

type repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		Db: db,
	}
}

func (r *repository) GetByID(id uint) (*User, error) {
	var entity *User
	res := r.Db.Where("id = ?", id).First(&entity)
	if res.Error != nil {
		return nil, res.Error
	}
	return entity, nil
}
