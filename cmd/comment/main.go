package main

import (
	"log"
	comment "micro_tiktok/kitex_gen/comment/commentvideoserver"
)

func main() {
	svr := comment.NewServer(new(CommentVideoServerImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
