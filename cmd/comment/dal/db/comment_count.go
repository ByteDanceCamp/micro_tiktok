package db

import (
	"context"
	"errors"
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

func Create(ctx context.Context, params *CommentCount) error {
	if err := DB.WithContext(ctx).Create(&params).Error; err != nil {
		return err
	}
	return nil
}

func GetCount(ctx context.Context, params *CommentCount) (res int64, err error) {
	err = DB.WithContext(ctx).Where(&params).Take(&params).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = Create(ctx, &CommentCount{
			VideoId:      params.VideoId,
			CommentCount: 0,
		})
		if err != nil {
			return 0, err
		}
		return 0, nil
	}
	return params.CommentCount, nil
}
