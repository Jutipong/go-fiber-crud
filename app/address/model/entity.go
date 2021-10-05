package model

import (
	"database/sql"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
)

// ** ใช้สำหรับ query database on gorm.
// ===================== model entity =====================//
type Address struct {
	AddressId  string       `gorm:"primaryKey; column:AddressId" json:"AddressId"`
	Address    string       `gorm:"column:Address" json:"Address"`
	Lat        *float64     `gorm:"column:Lat" json:"Lat"`
	Long       *float64     `gorm:"column:Long" json:"Long"`
	CreateDate time.Time    `gorm:"column:CreateDate;autoCreateTime" json:"-"`
	CreateBy   string       `gorm:"column:CreateBy;" json:"-"`
	UpdateDate sql.NullTime `gorm:"column:UpdateDate; autoUpdateTime;" json:"-"`
	UpdateBy   string       `gorm:"column:UpdateBy" json:"-"`
	IsActive   bool         `gorm:"column:IsActive; default:true" json:"-"`
}

func (a *Address) TableName() string {
	return "Address"
}

// Validation struct Auth
func (a *Address) Inquiry_Validation(c *fiber.Ctx) error {
	if utf8.RuneCountInString(strings.TrimSpace(a.AddressId)) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "required: AddressId")
	}
	return nil
}

func (a *Address) Create_Validation(c *fiber.Ctx) error {
	v := validate.Struct(a)
	v.AddRule("Address", "required")
	v.AddRule("Lat", "required")
	if !v.Validate() {
		return fiber.NewError(fiber.StatusBadRequest, v.Errors.Error())
	}
	return nil
}

func (a *Address) Update_Validation(c *fiber.Ctx) error {
	v := validate.Struct(a)
	v.AddRule("AddressId", "required")
	v.AddRule("Address", "required")
	v.AddRule("Lat", "required")
	v.AddRule("Long", "required")
	if !v.Validate() {
		return fiber.NewError(fiber.StatusBadRequest, v.Errors.Error())
	}
	return nil
}
func (a *Address) Delete_Validation(c *fiber.Ctx) error {
	v := validate.Struct(a)
	v.AddRule("AddressId", "required")
	if !v.Validate() {
		return fiber.NewError(fiber.StatusBadRequest, v.Errors.Error())
	}
	return nil
}
