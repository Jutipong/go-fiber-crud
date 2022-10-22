// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package schema

const TableNameSupplier = "Supplier"

// Supplier mapped from table <Supplier>
type Supplier struct {
	ID           int64   `gorm:"column:Id;type:int;primaryKey" json:"Id"`
	CompanyName  string  `gorm:"column:CompanyName;type:nvarchar;not null" json:"CompanyName"`
	ContactName  *string `gorm:"column:ContactName;type:nvarchar" json:"ContactName"`
	ContactTitle *string `gorm:"column:ContactTitle;type:nvarchar" json:"ContactTitle"`
	City         *string `gorm:"column:City;type:nvarchar" json:"City"`
	Country      *string `gorm:"column:Country;type:nvarchar" json:"Country"`
	Phone        *string `gorm:"column:Phone;type:nvarchar" json:"Phone"`
	Fax          *string `gorm:"column:Fax;type:nvarchar" json:"Fax"`
}

// TableName Supplier's table name
func (*Supplier) TableName() string {
	return TableNameSupplier
}
