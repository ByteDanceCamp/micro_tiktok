package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"micro_tiktok/pkg/constants"
	"time"
)

type ActionParams struct {
	Uid string
	Vid string
}

func Like(ctx context.Context, params *ActionParams) error {
	pipe := RDB.WithContext(ctx).TxPipeline()

	// 点赞列表中添加数据
	likeTime := time.Now().Unix()
	likeKey := constants.FavoriteLikePre + params.Uid
	pipe.ZAddNX(ctx, likeKey, &redis.Z{
		Score:  float64(likeTime),
		Member: params.Vid,
	})

	// 视频点赞数+1
	pipe.Incr(ctx, constants.FavoriteVideoPre+params.Vid)

	_, err := pipe.Exec(ctx)
	return err
}

func UnLike(ctx context.Context, params *ActionParams) error {
	pipe := RDB.WithContext(ctx).TxPipeline()

	// 点赞列表中删除数据
	likeKey := constants.FavoriteLikePre + params.Uid
	pipe.ZRem(ctx, likeKey, params.Vid)

	// 视频点赞数-1
	pipe.Decr(ctx, constants.FavoriteVideoPre+params.Vid)

	_, err := pipe.Exec(ctx)
	return err
}

// GetLikeCount 获取点赞数量
func GetLikeCount(ctx context.Context, findType, vid string) (int64, error) {
	key := findType + vid
	return RDB.WithContext(ctx).ZCard(ctx, key).Result()
}

// IsLike 是否点赞
func IsLike(ctx context.Context, keyPre, uid, vid string) (bool, error) {
	res := RDB.WithContext(ctx).ZScore(ctx, keyPre+uid, vid).Val()
	return res != 0, nil
}

// GetLikeList 获取点赞列表
func GetLikeList(ctx context.Context, key string, offset, limit int64) ([]string, error) {
	res := RDB.WithContext(ctx).ZRevRange(ctx, key, offset, limit).Val()
	return res, nil
}
