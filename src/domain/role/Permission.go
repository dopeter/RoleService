package role

import "github.com/jinzhu/gorm"

type Permission struct{
	gorm.Model
	Name string
	Operation string
	Path string
	Desc string
	RoleId int
}