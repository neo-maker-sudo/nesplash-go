package main

import (
	"fmt"
	"gin/api/router"
)

func main(){
	r := router.InitRouter()
	err := r.Run()
	if err != nil {
		fmt.Println("Error", err)
	}
}

