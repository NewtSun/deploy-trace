/**
 * @Author : NewtSun
 * @Date : 2023/8/12 16:25
 * @Description :
 **/

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string
	RealName string
	NickName string
	Email    string
	Phone    string
	Status   string
}
