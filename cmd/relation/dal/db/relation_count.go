package db

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"micro_tiktok/pkg/constants"
	"micro_tiktok/pkg/errno"
)

type RelationCount struct {
	gorm.Model
	Uid           int64 `gorm:"column:uid; unique; not null;"`
	FollowCount   int64 `gorm:"not null; type:bigint; default:0"`
	FollowerCount int64 `gorm:"not null; type:bigint; default:0"`
}

func (r *RelationCount) TableName() string {
	return constants.RelationCountTableName
}

func MGetCount(ctx context.Context, ids []int64) ([]*RelationCount, error) {
	res := make([]*RelationCount, 0)
	if len(ids) == 0 {
		return res, nil
	}
	if err := DB.WithContext(ctx).Where("id in ?", ids).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func FollowAction(ctx context.Context, uid, toUid int64) (err error) {
	if uid == toUid {
		return errno.RelationErr.WithMsg("can't follow yourself")
	}
	err = DB.Transaction(func(tx *gorm.DB) error {
		// 对 follow 表操作
		follow := &Follow{
			Uid:        uid,
			FollowedId: toUid,
		}
		err = tx.WithContext(ctx).Model(follow).Where(&follow).Take(&follow).Error
		// 查询出错
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		// 已经关注
		if err == nil && follow.Status != 0 {
			return errno.RelationErr.WithMsg("already followed the user")
		}
		// 记录不存在，插入新纪录
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			if err = tx.WithContext(ctx).Create(&follow).Error; err != nil {
				return err
			}
		}
		// 已取关，重新关注
		if err := tx.WithContext(ctx).Model(follow).Update("status", 1).Error; err != nil {
			return err
		}

		// 对 follower 表操作
		follower := &Follower{
			Uid:        toUid,
			FollowerId: uid,
		}
		err = tx.WithContext(ctx).Model(follower).Where(&follower).Take(&follower).Error
		// 查询出错
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		// 已被关注
		if err == nil && follower.Status != 0 {
			return errno.RelationErr.WithMsg("has been followed")
		}
		// 记录不存在，插入新纪录
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			if err = tx.WithContext(ctx).Create(&follower).Error; err != nil {
				return err
			}
		}
		// 被取关，重新被关注
		if err := tx.WithContext(ctx).Model(follower).Update("status", 1).Error; err != nil {
			return err
		}

		// 对 relation_count 表操作
		r1 := &RelationCount{Uid: uid}
		err = tx.WithContext(ctx).Where(&r1).Take(&r1).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			if err = tx.WithContext(ctx).Create(&r1).Error; err != nil {
				return err
			}
		}
		r2 := &RelationCount{Uid: toUid}
		err = tx.WithContext(ctx).Unscoped().Where(&r2).Take(&r2).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			if err = tx.WithContext(ctx).Create(&r2).Error; err != nil {
				return err
			}
		}
		if err = tx.WithContext(ctx).Model(r1).Unscoped().Where(&r1).Update("follow_count", r1.FollowCount+1).Error; err != nil {
			return err
		}
		if err = tx.WithContext(ctx).Model(r2).Unscoped().Where(&r2).Update("follower_count", r2.FollowerCount+1).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func UnFollowAction(ctx context.Context, uid, toUid int64) (err error) {
	if uid == toUid {
		return errno.RelationErr.WithMsg("can't unfollow yourself")
	}
	err = DB.Transaction(func(tx *gorm.DB) error {
		// 操作 follow 表
		follow := &Follow{
			Uid:        uid,
			FollowedId: toUid,
		}
		err = tx.WithContext(ctx).Model(follow).Where(&follow).Take(&follow).Error
		if err != nil {
			return err
		}
		if err == nil && follow.Status == 0 {
			return errno.RelationErr.WithMsg("not follow the user")
		}
		if err = tx.WithContext(ctx).Model(follow).Where(&follow).Update("status", 0).Error; err != nil {
			return err
		}
		// 操作 follower 表
		follower := &Follower{
			Uid:        toUid,
			FollowerId: uid,
		}
		err = tx.WithContext(ctx).Model(follower).Where(&follower).Take(&follower).Error
		if err != nil {
			return err
		}
		if err == nil && follower.Status == 0 {
			return errno.RelationErr.WithMsg("has not been followed")
		}
		if err = tx.WithContext(ctx).Model(follower).Where(&follower).Update("status", 0).Error; err != nil {
			return err
		}

		// 操作 relation_count 表
		r1 := &RelationCount{Uid: uid}
		if err = tx.WithContext(ctx).Model(r1).Where(&r1).Take(&r1).Error; err != nil {
			return err
		}
		r2 := &RelationCount{Uid: toUid}
		if err = tx.WithContext(ctx).Model(r2).Where(&r2).Take(&r2).Error; err != nil {
			return err
		}
		if err = tx.WithContext(ctx).Model(r1).Where(&r1).Update("follow_count", r1.FollowCount-1).Error; err != nil {
			return err
		}
		if err = tx.WithContext(ctx).Model(r2).Where(&r2).Update("follower_count", r2.FollowerCount-1).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func AddRecord(ctx context.Context, uid int64) (res *RelationCount, err error) {
	res = &RelationCount{
		Uid: uid,
	}
	if err = DB.WithContext(ctx).Create(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
