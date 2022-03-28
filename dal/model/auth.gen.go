// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"github.com/shopspring/decimal"
)

const TableNameAuth = "Auth"

// Auth mapped from table <Auth>
type Auth struct {
	ID         string          `gorm:"column:Id;type:nvarchar;not null;default:newid()" json:"Id"`
	UserName   string          `gorm:"column:UserName;type:nvarchar;not null" json:"UserName"`
	Password   string          `gorm:"column:Password;type:nvarchar;not null" json:"Password"`
	CreateDate time.Time       `gorm:"column:CreateDate;type:datetime;not null;default:getdate()" json:"CreateDate"`
	CreateBy   string          `gorm:"column:CreateBy;type:nvarchar;not null" json:"CreateBy"`
	UpdateDate *time.Time      `gorm:"column:UpdateDate;type:datetime" json:"UpdateDate"`
	UpdateBy   *string         `gorm:"column:UpdateBy;type:nvarchar" json:"UpdateBy"`
	IsActive   bool            `gorm:"column:IsActive;type:bit;not null;default:1" json:"IsActive"`
	UserID     string          `gorm:"column:UserId;type:nvarchar;not null;default:newid()" json:"UserId"`
	Amount     decimal.Decimal `gorm:"column:Amount;type:decimal" json:"Amount"`
}

// TableName Auth's table name
func (*Auth) TableName() string {
	return TableNameAuth
}
