package main

import (
	"fmt"
	"github.com/wagnerm/gogopro"
)

func main() {
	gopro, err := gopro.CreateGoPro("10.5.5.9").Init()
	if err != nil {
		panic(err)
	}
	status, err := gopro.Status()
	fmt.Println(status)
}
