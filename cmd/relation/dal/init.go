package dal

import (
	"micro_tiktok/cmd/relation/dal/db"
	"micro_tiktok/cmd/relation/dal/redis"
)

func Init() {
	redis.Init()
	db.Init()
}
