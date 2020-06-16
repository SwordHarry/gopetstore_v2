package domain

type Category struct {
	CategoryId  string `gorm:"primary_key;column:catid"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:descn"`
}

// 设置 Category 的表名为`category`
func (*Category) TableName() string {
	return "category"
}
