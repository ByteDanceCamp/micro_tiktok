package pack

import (
	"micro_tiktok/cmd/video/dal/db"
	"micro_tiktok/kitex_gen/video"
	"micro_tiktok/pkg/constants"
)

// Video db to idl
func Video(v *db.Video) *video.Video {
	if v == nil {
		return nil
	}
	return &video.Video{
		Id: int64(v.ID),
		Author: &video.User{
			Id:            v.Uid,
			Name:          "",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		},
		PlayUrl:       constants.QiNiuServer + v.PlayUrl,
		CoverUrl:      constants.QiNiuServer + v.CoverUrl,
		Title:         v.Title,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}
}

func Videos(videos []*db.Video) []*video.Video {
	vs := make([]*video.Video, 0)
	for _, vid := range videos {
		if v := Video(vid); v != nil {
			vs = append(vs, v)
		}
	}
	return vs
}
