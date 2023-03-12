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

func BatchGetUserByIds(ids []int64, creator string) ([]UserModel, error) {
	var userList = []UserModel{}
	query := DbConnect.Where(ids)

	if creator != "" {
		query = query.Where("user_name LIKE BINARY ?", "%"+creator+"%")
	}

	result := query.Find(&userList)
	return userList, result.Error
}
