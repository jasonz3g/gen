// Code generated by github.com/jasonz3g/gen. DO NOT EDIT.
// Code generated by github.com/jasonz3g/gen. DO NOT EDIT.
// Code generated by github.com/jasonz3g/gen. DO NOT EDIT.

package model

const TableNameBank = "banks"

// Bank mapped from table <banks>
type Bank struct {
	ID      int64   `gorm:"column:id;primaryKey;autoIncrement:true" json:"-"`
	Name    *string `gorm:"column:name" json:"-"`
	Address *string `gorm:"column:address" json:"-"`
	Scale   *int64  `gorm:"column:scale" json:"-"`
}

// TableName Bank's table name
func (*Bank) TableName() string {
	return TableNameBank
}
