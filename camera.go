package gogopro

import (
	"errors"
	"fmt"
)

type Camera struct {
	APIRequester   *APIRequester
	StatusCommands map[string]StatusCommand
}

func (c *Camera) Init() *Camera {
	return c
}
func CreateCamera(APIRequester *APIRequester) *Camera {
	camera := &Camera{}
	camera.APIRequester = APIRequester
	statusCommands := CreateCameraStatusCommands()
	camera.StatusCommands = statusCommands
	return camera
}

func CreateCameraStatusCommands() map[string]StatusCommand {
	sc := make(map[string]StatusCommand)
	sc["mode"] = StatusCommand{
		ResultByte:  1,
		Translaters: map[byte]string{0: "video", 1: "photo", 2: "burst", 3: "timelapse"},
	}
	sc["defaultmode"] = StatusCommand{
		ResultByte:  3,
		Translaters: map[byte]string{0: "video", 1: "photo", 2: "burst", 3: "timelapse"},
	}
	sc["spotmeter"] = StatusCommand{
		ResultByte:  4,
		Translaters: map[byte]string{0: "off", 1: "on"},
	}
	sc["timelapse_interval"] = StatusCommand{
		ResultByte:  5,
		Translaters: map[byte]string{0: "0.5s", 1: "1s", 2: "2s", 5: "5s", 10: "10s", 40: "30s", 60: "60s"},
	}
	sc["fov"] = StatusCommand{
		ResultByte:  7,
		Translaters: map[byte]string{0: "wide", 1: "medium", 2: "narrow"},
	}
	sc["photores"] = StatusCommand{
		ResultByte:  8,
		Translaters: map[byte]string{3: "5MP_med", 4: "7MP_wide", 8: "10MP_wide"},
	}
	/* TODO nil translaters
	sc["minselapsed"] = StatusCommand{
		ResultByte:  13,
		Translaters: nil,
	}
	sc["secselapsed"] = StatusCommand{
		ResultByte:  14,
		Translaters: nil,
	}
	*/
	sc["volume"] = StatusCommand{
		ResultByte:  16,
		Translaters: map[byte]string{0: "off", 1: "70%", 2: "100%"},
	}
	sc["led"] = StatusCommand{
		ResultByte:  17,
		Translaters: map[byte]string{0: "off", 1: "led2", 2: "led4"},
	}
	/* TODO: Photo/video counts for two bytes?
		sc["photoremaining"] = StatusCommand{
			ResultByte:  21,
			Translaters: nil}
		sc["photocount"] = StatusCommand{
			ResultByte:  23,
			Translaters: nil}
		sc["videoremaining"] = StatusCommand{
			ResultByte:  25,
			Translaters: nil}
		sc["videocount"] = StatusCommand{
			ResultByte:  25,
			Translaters: nil}
		sc["recording"] = StatusCommand{
			ResultByte: 29,
			Translaters: map[byte]string{0:"off", 1:"on"},
	    }
		sc["videoresolution"] = StatusCommand{
			ResultByte: 50,
			Translaters: map[byte]string{0:"WVGA", 1:"720", 2:"960", 3:"1080"},
		}
		sc["fps"] = StatusCommand{
			ResultByte: 51,
			Translaters: map[byte]string{3:"25", 6:"50"},
		}
		/*
			/*
				TODO: Need support for checking bits
				sc["orientation"] = StatusCommand{
					ResultByte: 18,
					Translaters: map[byte]string{
						StatusTranslater{
							Result:         0,
							ExpectedReturn: "up"},
						StatusTranslater{
							Result:         4,
							ExpectedReturn: "down"}}}
				sc["iso_sharpness"] = StatusCommand{
						ResultByte: 6,
						Translaters: map[byte]string{
							StatusTranslater{
								Result:         0,
								ExpectedReturn: "off"},
							StatusTranslater{
								Result:         1,
								ExpectedReturn: "on"}}}
				TODO: Byte 30
				====
				Protune
				Low light
				Color
	*/
	return sc
}
func (c *Camera) Status(command string) (string, error) {
	sCmd, ok := c.StatusCommands[command]
	if ok == false {
		return "", errors.New(fmt.Sprintf("No camera status command %s", command))
	}
	result, err := sCmd.RunStatusCommand("/camera/sx", c.APIRequester)
	if err != nil {
		return "", err
	}
	return result, nil
}
