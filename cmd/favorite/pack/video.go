package pack

import (
	"micro_tiktok/kitex_gen/favorite"
	"micro_tiktok/kitex_gen/video"
)

//	Video  service Video to favorite.Video
func Video(v *video.Video) *favorite.Video {
	if v == nil {
		return nil
	}
	return &favorite.Video{
		Id: v.Id,
		Author: &favorite.User{
			Id:            v.Author.Id,
			Name:          v.Author.Name,
			FollowCount:   v.Author.FollowCount,
			FollowerCount: v.Author.FollowerCount,
			IsFollow:      v.Author.IsFollow,
		},
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
	fvs := make([]*favorite.Video, 0)
	for _, v := range vs {
		fvs = append(fvs, Video(v))
	}
	return fvs
}
