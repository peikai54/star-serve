package model

import (
	"time"
)

type UserModel struct {
	UserId    int64 `gorm:"primaryKey"`
	UserName  string
	NickName  string
	Password  string
	CreatedAt time.Time
	IsAdmin   int64
}

func (UserModel) TableName() string {
	return "user"
}

func BatchGetUserByIds(ids []int64) ([]UserModel, error) {
	var userList = []UserModel{}
	result := DbConnect.Where(ids).Find(&userList)
	return userList, result.Error
}
