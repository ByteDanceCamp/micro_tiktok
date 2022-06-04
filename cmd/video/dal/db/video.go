package db

import (
	"context"
	"micro_tiktok/pkg/constants"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Author   string `json:"user_id" gorm:"not null; type: varchar(80)"`
	PlayUrl  string `json:"play_url" gorm:"not null"`
	CoverUrl string `json:"cover_url" gorm:"not null"`
	Title    string `json:"title" gorm:"not null"`
}

func (v *Video) TableName() string {
	return constants.VideoTableName
}

// Create ，将新上传的视频插入数据库
func Create(ctx context.Context, video *Video) error {
	if err := DB.WithContext(ctx).Create(&video).Error; err != nil {
		return err
	}
	return nil
}

// GetMyList 获取登陆用户自己的食品发布列表
func GetMyList(ctx context.Context, user_id int64) ([]*Video, error) {
	res := make([]*Video, 0)

	if err := DB.WithContext(ctx).Where("author = ?", user_id).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetNewestList(ctx context.Context, latest_time int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if result := DB.WithContext(ctx).Order("id desc").Limit(30).Find(&res); result.Error != nil {
		return nil, result.Error
	}
	return res, nil
}

func QuertyVideos(ctx context.Context, vid_list []int64) ([]*Video, error) {
	res := make([]*Video, 0)
	for _, id := range vid_list {
		vi := &Video{}
		if err := DB.WithContext(ctx).Where("id = ?", id).Take(&vi).Error; err != nil {
			return nil, err
		}
		res = append(res, vi)
	}
	return res, nil
}
