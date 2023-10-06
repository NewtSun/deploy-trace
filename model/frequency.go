package model

import "gorm.io/gorm"

type Frequency struct {
	gorm.Model
	UserName string
	RealName string
	NickName string
	Value    int64
}
