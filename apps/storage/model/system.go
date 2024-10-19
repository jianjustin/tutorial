package model

type SystemParam struct {
	Id    string `gorm:"id" json:"id"`
	Key   string `gorm:"key" json:"key"`
	Value string `gorm:"value" json:"value"`
	G     string `gorm:"g" json:"g"`
}
