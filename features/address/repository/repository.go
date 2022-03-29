package repository

import (
	"fiber-crud/features/address/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IRepository interface {
	Inquiry(auth *model.Address) ([]model.Address, error)
	Create(auth *model.Address) error
	Update(auth *model.Address) error
	Delete(auth *model.Address) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return &repository{db}
}

func (r *repository) Inquiry(address *model.Address) (result []model.Address, err error) {
	//## หากต้องการดู string query ที่ gorm generate ให้ให้ใช่ .Debuger()
	//## ตัวอย่าง err = r.db.Debug().Find(&result).Error
	err = r.db.Where("IsActive = ?", false).Find(&result, &address).Error
	if err != nil {
		return result, err
	}

	if len(result) == 0 {
		return result, fiber.ErrNotFound
	}

	return result, nil
}

func (r *repository) Create(address *model.Address) error {
	query := r.db.Select("AddressId", "Address", "Lat", "Long", "CreateDate", "CreateBy").Omit("UpdateDate").Create(&address)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *repository) Update(address *model.Address) error {
	query := r.db.Select("Address", "Lat", "Long", "UpdateBy").Updates(&address)
	if query.Error != nil {
		return query.Error
	}

	if query.RowsAffected == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

//Demo นี้จะไม่ทำการลบข้อมูลจริงๆ แต่จะทำการ set IsActive = false
func (r *repository) Delete(address *model.Address) error {
	query := r.db.Find(address)
	if query.Error != nil {
		return query.Error
	}

	if query.RowsAffected == 0 {
		return fiber.ErrNotFound
	}

	address.IsActive = false
	err := r.db.Save(address).Error
	if err != nil {
		return err
	}

	return nil
}
