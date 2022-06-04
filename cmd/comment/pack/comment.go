package pack

import (
	"micro_tiktok/cmd/comment/dal/db"
	"micro_tiktok/kitex_gen/comment"
)

func Comment(c *db.Comment) *comment.Comment {
	if c == nil {
		return nil
	}
	return &comment.Comment{
		Id:         int64(c.ID),
		Content:    c.Content,
		CreateDate: c.CreatedAt.Format("04-02"),
		User: &comment.User{
			Id: c.Uid,
		},
	}
}

func Comments(comments []*db.Comment) []*comment.Comment {
	cs := make([]*comment.Comment, 0)
	for _, c := range comments {
		if c2 := Comment(c); c2 != nil {
			cs = append(cs, c2)
		}
	}
	return cs
}
