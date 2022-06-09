package db

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"micro_tiktok/pkg/constants"
	"micro_tiktok/pkg/errno"
)

type Comment struct {
	gorm.Model
	VideoId int64  `gorm:"not null"`
	Uid     int64  `gorm:"not null"`
	Content string `gorm:"type: varchar(255); not null"`
}

func (c *Comment) TableName() string {
	return constants.CommentTableName
}

func Publish(ctx context.Context, params *Comment) (res *Comment, err error) {
	err = DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 发布评论
		if err = tx.WithContext(ctx).Create(&params).Error; err != nil {
			return err
		}
		// 更新计数表
		countParam := &CommentCount{
			VideoId: params.VideoId,
		}
		err = tx.WithContext(ctx).Model(&countParam).Where(&countParam).Take(&countParam).Error
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			// 新增记录
			if err = tx.WithContext(ctx).Create(&countParam).Error; err != nil {
				return err
			}
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		// 自增
		if err = tx.WithContext(ctx).Model(countParam).Where("id = ?", countParam.ID).Update("comment_count", countParam.CommentCount+1).Error; err != nil {
			return err
		}
		return nil
	})
	return params, nil
}

func Delete(ctx context.Context, params *Comment, uid int64) (err error) {

	err = DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.WithContext(ctx).Model(params).Where("id = ?", params.ID).Take(&params).Error; err != nil {
			return errno.ParamsErr.WithMsg("has been deleted")
		}
		if params.Uid != uid {
			return errno.CommentErr.WithMsg("can't del others comment")
		}
		if err = tx.WithContext(ctx).Model(params).Where("id = ?", params.ID).Delete(params.ID).Error; err != nil {
			return err
		}

		countParam := &CommentCount{
			VideoId: params.VideoId,
		}
		if err = tx.WithContext(ctx).Where(&countParam).Take(&countParam).Error; err != nil {
			return err
		}
		if countParam.CommentCount < 1 {
			return errno.CommentErr
		}
		if err = DB.WithContext(ctx).Model(countParam).Where("id = ?", countParam.ID).Update("comment_count", countParam.CommentCount-1).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func GetList(ctx context.Context, params *Comment) (res []*Comment, err error) {
	err = DB.WithContext(ctx).Where("video_id = ?", params.VideoId).Find(&res).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return res, nil
	}
	return res, nil
}
