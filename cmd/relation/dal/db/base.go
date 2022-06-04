package db

import "time"

type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    int64 `gorm:"type:tinyint; default:1"` // 1:关注；0：取关
}
