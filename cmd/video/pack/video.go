package pack

import (
	"micro_tiktok/cmd/video/dal/db"
	"micro_tiktok/kitex_gen/video"
	"strconv"
)

// video db to idl
func Video(v *db.Video) *video.Video {
	if v == nil {
		return nil
	}
	tmp, _ := strconv.ParseInt(v.Author, 10, 64)
	return &video.Video{
		Id: int64(v.ID),
		Author: &video.User{
			Id: tmp,
		},
		PlayUrl:  v.PlayUrl,
		CoverUrl: v.CoverUrl,
		Title:    v.Title,
	}
}

func Videos(vedios []*db.Video) []*video.Video {
	vs := make([]*video.Video, 0)
	for _, vid := range vedios {
		if onevi := Video(vid); onevi != nil {
			vs = append(vs, onevi)
		}
	}
	return vs
}
