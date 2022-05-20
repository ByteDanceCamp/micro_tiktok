package db

import (
	"context"
	"gorm.io/gorm"
	"micro_tiktok/pkg/constants"
)

type User struct {
	gorm.Model
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// Create 新建用户，将新记录插入数据库
func Create(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).Create(user).Error
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
	user := User{}
	if err := DB.WithContext(ctx).Where("user_name = ?", username).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
