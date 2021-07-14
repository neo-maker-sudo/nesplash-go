package main

import (
	"gin/api/router"
)

func main(){
	r := router.InitRouter()
	r.Run()
}

