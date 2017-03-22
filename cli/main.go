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

	commands := []string{"mode", "defaultmode", "spotmeter", "timelapse_interval", "fov", "photores",
		"minselapsed", "secselapsed", "volume", "led", "recording", "videoresolution", "fps",
		"photoremaining", "photocount", "videoremaining", "videocount"}
	for _, c := range commands {
		fmt.Println(c + ":")
		status, err := (*gopro.Camera).Status(c)
		if err != nil {
			panic(err)
		}
		fmt.Println("\t" + status)
	}

}
