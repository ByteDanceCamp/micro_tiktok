package pack

import (
	"micro_tiktok/kitex_gen/favorite"
	"micro_tiktok/kitex_gen/video"
)

//	Video video.Video to favorite.Video
func Video(v *video.Video) *favorite.Video {
	if v == nil {
		return nil
	}
	return &favorite.Video{
		Id: v.Id,
		//	这里的报错不知道咋搞的。。
		Author:        v.Author,
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    v.IsFavorite,
		Title:         v.Title,
	}
}

//	Videos []*video.Video to []*favorite.Video
func Videos(vs []*video.Video) []*favorite.Video {
	if vs == nil {
		return nil
	}
	fvs := make([]*favorite.Video, len(vs))
	for i, v := range vs {
		if v != nil {
			fvs[i] = Video(v)
		}
	}
	return fvs
}
