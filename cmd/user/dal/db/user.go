package db

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"micro_tiktok/pkg/constants"
)

type User struct {
	gorm.Model
	UserName string `json:"username" gorm:"unique; not null; type: varchar(80)"`
	PassWord string `json:"password" gorm:"not null; type: varchar(64)"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// Create 新建用户，将新记录插入数据库
func Create(ctx context.Context, user *User) (*User, error) {
	if err := DB.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// MGet 批量获取用户信息
func MGet(ctx context.Context, ids []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(ids) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", ids).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// QueryUser 根据用户名获取用户信息
func QueryUser(ctx context.Context, username string) (*User, error) {
	user := &User{}
	if err := DB.WithContext(ctx).Where("user_name = ?", username).Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func IsExist(ctx context.Context, id int64) (bool, error) {
	res := &User{}
	err := DB.WithContext(ctx).Where("id = ?", id).Take(&res).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	return true, nil
}
