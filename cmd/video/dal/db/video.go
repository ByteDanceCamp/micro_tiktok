package db

import (
	"context"
	"micro_tiktok/pkg/constants"
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Uid      int64  `gorm:"column:user_id; not null"`
	PlayUrl  string `jgorm:"not null; type: varchar(255)"`
	CoverUrl string `gorm:"not null; type: varchar(255)"`
	Title    string `gorm:"not null; type: varchar(255)"`
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
func GetMyList(ctx context.Context, userId int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Order("updated_at desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetNewestList(ctx context.Context, latestTime int64) ([]*Video, int64, error) {
	res := make([]*Video, 0)
	if latestTime <= 0 {
		latestTime = time.Now().Unix()
	}
	if result := DB.WithContext(ctx).Where("updated_at <= ?", time.Unix(latestTime, 0).Local()).Order("updated_at desc").Limit(5).Find(&res); result.Error != nil {
		return nil, 0, result.Error
	}
	var nextTime int64
	if len(res) > 0 {
		nextTime = res[len(res)-1].CreatedAt.Unix()
	} else {
		nextTime = time.Now().Unix()
	}

	return res, nextTime, nil
}

func MGet(ctx context.Context, vidList []int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("id in ?", vidList).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryById(ctx context.Context, id int64) (res *Video, err error) {
	if err = DB.WithContext(ctx).Where("id = ?", id).Take(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
