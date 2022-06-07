package dal

import "micro_tiktok/cmd/favorite/dal/redis"

func Init() {
	redis.Init()
}
