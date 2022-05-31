package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"micro_tiktok/pkg/constants"
	"time"
)

type ActionParams struct {
	Uid   string
	ToUid string
}

func Follow(ctx context.Context, params *ActionParams) error {
	var err error
	pipe := RDB.WithContext(ctx).TxPipeline()

	// 自己的关注表添加数据
	followTime := time.Now().Unix()
	followKey := constants.RelationFollowPre + params.Uid
	pipe.ZAddNX(ctx, followKey, &redis.Z{
		Score:  float64(followTime),
		Member: params.ToUid,
	})

	// 目标的粉丝表添加数据
	fansKey := constants.RelationFansPre + params.ToUid
	pipe.ZAddNX(ctx, fansKey, &redis.Z{
		Score:  float64(followTime),
		Member: params.Uid,
	})

	_, err = pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func UnFollow(ctx context.Context, params *ActionParams) error {
	pipe := RDB.WithContext(ctx).TxPipeline()

	// 自己的关注列表删除数据
	followKey := constants.RelationFollowPre + params.Uid
	pipe.ZRem(ctx, followKey, params.ToUid)

	// 目标粉丝列表中删除数据
	fansKey := constants.RelationFansPre + params.ToUid
	pipe.ZRem(ctx, fansKey, params.Uid)

	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

// GetCount 获取关注/粉丝总数
// findType 与 uid 拼接为带查询集合的 key
func GetCount(ctx context.Context, findType, uid string) (count int64, err error) {
	count, err = RDB.ZCard(ctx, findType+uid).Result()
	return count, err
}

// IsFollow 判断目标用户是否在当前用户关注列表中
func IsFollow(ctx context.Context, key, targetUid string) bool {
	res := RDB.ZScore(ctx, key, targetUid).Val()
	if res != 0 {
		return true
	}
	return false
}

// GetListStrSlice 以字符串切片形式返回需要的列表
func GetListStrSlice(ctx context.Context, key string, startIndex, endIndex int64) ([]string, error) {
	return RDB.ZRevRange(ctx, key, startIndex, endIndex).Result()
}
