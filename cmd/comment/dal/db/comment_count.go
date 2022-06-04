package db

import (
	"context"
	"gorm.io/gorm"
	"micro_tiktok/pkg/constants"
)

type CommentCount struct {
	gorm.Model
	VideoId      int64 `gorm:"unique; not null"`
	CommentCount int64 `gorm:"not null; default:0"`
}

func (c *CommentCount) TableName() string {
	return constants.CommentCountTableName
}

func GetCount(ctx context.Context, params *CommentCount) (res int64, err error) {
	if err = DB.WithContext(ctx).Where(&params).Take(&params).Error; err != nil {
		return -1, err
	}
	return params.CommentCount, nil
}
