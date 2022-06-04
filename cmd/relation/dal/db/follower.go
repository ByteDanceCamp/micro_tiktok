package db

import (
	"context"
	"micro_tiktok/pkg/constants"
)

type Follower struct {
	BaseModel  `gorm:"embedded"`
	Uid        int64 `gorm:"column:uid; not null; unique_index:U_FR"`
	FollowerId int64 `gorm:"column:follower_id; unique_index:U_FR"`
}

func (f *Follower) TableName() string {
	return constants.FollowerTableName
}

func MGetFollowerList(ctx context.Context, uid int64) (res []*Follower, err error) {
	if err = DB.WithContext(ctx).Where("uid = ? AND status <> ?", uid, 0).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
