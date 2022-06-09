package db

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"micro_tiktok/pkg/constants"
)

type Follow struct {
	BaseModel  `gorm:"embedded"`
	Uid        int64 `gorm:"column:uid; not null; unique_index:U_FD"`
	FollowedId int64 `gorm:"column:followed_id; unique_index:U_FD"`
}

func (r *Follow) TableName() string {
	return constants.FollowTableName
}

func IsFollow(ctx context.Context, uid, toUid int64) (bool, error) {
	data := &Follow{
		BaseModel: BaseModel{
			Status: 1,
		},
		Uid:        uid,
		FollowedId: toUid,
	}
	if uid == toUid {
		return false, nil
	}
	err := DB.WithContext(ctx).Where(&data).Take(&data).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	return true, nil
}

func MGetFollowList(ctx context.Context, uid int64) (res []*Follow, err error) {
	if err = DB.WithContext(ctx).Where("uid = ? AND status <> ?", uid, 0).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
