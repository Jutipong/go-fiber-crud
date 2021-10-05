package repository

import (
	"fiber-crud/app/address/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IRepository interface {
	Inquiry(auth *model.Address) ([]model.Address, error)
	Create(auth *model.Address) (model.Address, error)
	Update(auth *model.Address) (model.Address, error)
	Delete(auth *model.Address) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return &repository{db: db}
}

func (r *repository) Inquiry(address *model.Address) (result []model.Address, err error) {
	//## หากต้องการดู string query ที่ gorm generate ให้ให้ใช่ .Debuger()
	//## ตัวอย่าง err = r.db.Debug().Find(&result).Error
	err = r.db.Debug().Find(&result, &address).Error
	if err != nil {
		return result, err
	}

	if len(result) == 0 {
		return result, fiber.ErrNotFound
	}

	return result, nil
}

func (r *repository) Create(address *model.Address) (result model.Address, err error) {

	return result, nil
}

func (r *repository) Update(address *model.Address) (result model.Address, err error) {

	return result, nil
}

func (r *repository) Delete(address *model.Address) error {

	return nil
}
