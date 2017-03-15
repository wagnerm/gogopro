package main

import (
	"fmt"
	"github.com/wagnerm/gopro"
)

func main() {
	gopro, err := gopro.CreateGoPro("10.5.5.9").Init()
	if err != nil {
		panic(err)
	}
	fmt.Println(gopro.Status())
}
