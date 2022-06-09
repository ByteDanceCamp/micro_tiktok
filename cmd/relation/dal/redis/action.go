package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"micro_tiktok/cmd/relation/dal/db"
	"micro_tiktok/pkg/constants"
	"micro_tiktok/pkg/errno"
	"strconv"
)

type ActionParams struct {
	Uid   int64
	ToUid int64
}

type CountInfo struct {
	Uid           int64
	FollowCount   int64
	FollowerCount int64
}

// Action 关注/取关，均删除 redis 记录
func Action(ctx context.Context, data *ActionParams) (err error) {
	if data.Uid == data.ToUid {
		return errno.RelationErr.WithMsg("can't follow or unfollow yourself")
	}
	// 删除 redis 的缓存
	uid, toUid := strconv.Itoa(int(data.Uid)), strconv.Itoa(int(data.ToUid))
	pipe := RDB.WithContext(ctx).TxPipeline()
	// 删除自己列表的关注数据
	followKey := constants.RelationFollowPre + uid
	//pipe.ZRem(ctx, followKey, toUid)
	// 删除目标用户的粉丝数据
	fansKey := constants.RelationFansPre + toUid
	//pipe.ZRem(ctx, fansKey, uid)
	// 删除自己和目标用户的关系数表
	countUidKey := constants.RelationCountPre + uid
	countToUidKey := constants.RelationCountPre + toUid
	pipe.Del(ctx, followKey, fansKey, countUidKey, countToUidKey)
	_, err = pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func GetCount(ctx context.Context, uid int64) (*CountInfo, errno.ErrNo) {
	res := &CountInfo{
		Uid: uid,
	}
	var err error
	id := strconv.Itoa(int(uid))
	key := constants.RelationCountPre + id
	if RDB.Exists(ctx, key).Val() == 0 {
		return nil, errno.RelationCacheMissErr.WithMsg("cache isn't exists")
	}
	followCountStr := RDB.LIndex(ctx, key, 0).Val()
	fansCountStr := RDB.LIndex(ctx, key, 1).Val()
	followCount, err := strconv.Atoi(followCountStr)
	if err != nil {
		return nil, errno.ConvertErr(err)
	}
	res.FollowCount = int64(followCount)
	followerCount, err := strconv.Atoi(fansCountStr)
	if err != nil {
		return nil, errno.ConvertErr(err)
	}
	res.FollowerCount = int64(followerCount)
	return res, errno.Success
}

func IsFollow(ctx context.Context, data *ActionParams) (bool, errno.ErrNo) {
	if data.Uid == data.ToUid {
		return false, errno.Success
	}
	uid, toUid := strconv.Itoa(int(data.Uid)), strconv.Itoa(int(data.ToUid))
	followKey := constants.RelationFollowPre + uid
	if RDB.Exists(ctx, followKey).Val() == 0 {
		return false, errno.RelationCacheMissErr.WithMsg("cache not exists")
	}
	res := RDB.ZScore(ctx, followKey, toUid).Val()
	if res != 0 {
		return true, errno.Success
	}
	return false, errno.Success

}

func AddFollow(ctx context.Context, follow *db.Follow) (err error) {
	followTime := follow.CreatedAt.Unix()
	followKey := constants.RelationFollowPre + strconv.Itoa(int(follow.Uid))
	err = RDB.ZAddNX(ctx, followKey, &redis.Z{
		Score:  float64(followTime),
		Member: follow.FollowedId,
	}).Err()
	if err != nil {
		return err
	}
	return nil
}

func AddFollower(ctx context.Context, follower *db.Follower) (err error) {
	followerTime := follower.CreatedAt.Unix()
	followKey := constants.RelationFansPre + strconv.Itoa(int(follower.Uid))
	err = RDB.ZAddNX(ctx, followKey, &redis.Z{
		Score:  float64(followerTime),
		Member: follower.FollowerId,
	}).Err()
	if err != nil {
		return err
	}
	return nil
}

func UpdateCount(ctx context.Context, data *CountInfo) (err error) {
	key := constants.RelationCountPre + strconv.Itoa(int(data.Uid))
	err = RDB.LPush(ctx, key, data.FollowerCount, data.FollowCount).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetListStrSlice 以字符串切片形式返回需要的列表
func GetListStrSlice(ctx context.Context, key string, startIndex, endIndex int64) ([]string, errno.ErrNo) {
	if RDB.Exists(ctx, key).Val() == 0 {
		return nil, errno.RelationErr.WithMsg("cache is not exists")
	}
	res, err := RDB.ZRevRange(ctx, key, startIndex, endIndex).Result()
	if err != nil {
		return nil, errno.ConvertErr(err)
	}
	return res, errno.Success
}
