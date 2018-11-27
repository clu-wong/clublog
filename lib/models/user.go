package models

import (
	"github.com/jinzhu/gorm"
	"strings"
)

type User struct{
	gorm.Model
	Name string `gorm:"type:varchar(100);unique_index"`
	Password string
}

func (user User) ValidPassword(passwd string) bool{
	var flag bool
	flag = user.Password == strings.TrimSpace(passwd)
	return flag
}
