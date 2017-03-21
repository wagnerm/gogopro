package main

import (
	"fmt"
	"github.com/wagnerm/gogopro"
)

func main() {
	gopro, err := gogopro.CreateGoPro("10.5.5.9").Init()
	if err != nil {
		panic(err)
	}
	fmt.Println(gopro)
	status, err := (*gopro.Power).GetPowerStatus()
	if err != nil {
		panic(err)
	}
	fmt.Println(status)

	status, err := (*gopro.Camera).GetMode()
	if err != nil {
		panic(err)
	}
	fmt.Println(status)
}
