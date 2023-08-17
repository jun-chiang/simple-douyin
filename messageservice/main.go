package main

import (
	api "github.com/jun-chiang/simple-douyin/messageservice/kitex_gen/api/messageservice"
	"log"
)

func main() {
	svr := api.NewServer(new(MessageServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
