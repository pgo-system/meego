// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	UID               int32     `gorm:"column:uid;primaryKey;autoIncrement:true" json:"uid"`
	UserName          string    `gorm:"column:user_name;not null" json:"user_name"`                   // 用户名
	Password          string    `gorm:"column:password;not null" json:"password"`                     // 哈希密码
	Points            int32     `gorm:"column:points;not null" json:"points"`                         // 积分
	AvatarURL         string    `gorm:"column:avatar_url;not null" json:"avatar_url"`                 // 头像
	Email             string    `gorm:"column:email;not null" json:"email"`                           // 邮箱
	PhoneNumber       string    `gorm:"column:phone_number;not null" json:"phone_number"`             // 加密电话
	PersonalSignature string    `gorm:"column:personal_signature;not null" json:"personal_signature"` // 个性签名
	CreateTime        time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime        time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
