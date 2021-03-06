// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"github.com/shopspring/decimal"
)

const TableNameAddress = "Address"

// Address mapped from table <Address>
type Address struct {
	AddressID  string          `gorm:"column:AddressId;type:nvarchar;primaryKey;default:newid()" json:"AddressId"`
	Address    string          `gorm:"column:Address;type:text;not null" json:"Address"`
	Lat        decimal.Decimal `gorm:"column:Lat;type:decimal" json:"Lat"`
	Long       decimal.Decimal `gorm:"column:Long;type:decimal" json:"Long"`
	CreateDate time.Time       `gorm:"column:CreateDate;type:datetime;not null;default:getdate()" json:"CreateDate"`
	CreateBy   string          `gorm:"column:CreateBy;type:nvarchar;not null" json:"CreateBy"`
	UpdateDate *time.Time      `gorm:"column:UpdateDate;type:datetime" json:"UpdateDate"`
	UpdateBy   *string         `gorm:"column:UpdateBy;type:nvarchar" json:"UpdateBy"`
	IsActive   bool            `gorm:"column:IsActive;type:bit;not null;default:1" json:"IsActive"`
	Price      decimal.Decimal `gorm:"column:Price;type:decimal" json:"Price"`
}

// TableName Address's table name
func (*Address) TableName() string {
	return TableNameAddress
}
