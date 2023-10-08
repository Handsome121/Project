package models

import "gorm.io/gorm"

type GroupBasic struct {
	gorm.Model
	Name  string
	OwnId uint
	Icon  string
	Type  string
	Desc  string
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
