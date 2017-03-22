package gogopro

import ()

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
		Endpoint:   "/camera/sx",
		ResultByte: 1,
		Translaters: []StatusTranslater{
			StatusTranslater{
				Result:         0,
				ExpectedReturn: "video"},
			StatusTranslater{
				Result:         1,
				ExpectedReturn: "photo"},
			StatusTranslater{
				Result:         2,
				ExpectedReturn: "burst"},
			StatusTranslater{
				Result:         3,
				ExpectedReturn: "timelapse"}}}
	sc["defaultmode"] = StatusCommand{
		Endpoint:   "/camera/sx",
		ResultByte: 3,
		Translaters: []StatusTranslater{
			StatusTranslater{
				Result:         0,
				ExpectedReturn: "video"},
			StatusTranslater{
				Result:         1,
				ExpectedReturn: "photo"},
			StatusTranslater{
				Result:         2,
				ExpectedReturn: "burst"},
			StatusTranslater{
				Result:         3,
				ExpectedReturn: "timelapse"}}}
	sc["spotmeter"] = StatusCommand{
		Endpoint:   "/camera/sx",
		ResultByte: 4,
		Translaters: []StatusTranslater{
			StatusTranslater{
				Result:         0,
				ExpectedReturn: "off"},
			StatusTranslater{
				Result:         1,
				ExpectedReturn: "on"}}}
	sc["timelapse_interval"] = StatusCommand{
		Endpoint:   "/camera/sx",
		ResultByte: 5,
		Translaters: []StatusTranslater{
			StatusTranslater{
				Result:         0,
				ExpectedReturn: "0.5s"},
			StatusTranslater{
				Result:         1,
				ExpectedReturn: "1s"},
			StatusTranslater{
				Result:         2,
				ExpectedReturn: "2s"},
			StatusTranslater{
				Result:         5,
				ExpectedReturn: "5s"},
			StatusTranslater{
				Result:         10,
				ExpectedReturn: "10s"},
			StatusTranslater{
				Result:         40,
				ExpectedReturn: "30s"},
			StatusTranslater{
				Result:         60,
				ExpectedReturn: "60s"}}}
	sc["fov"] = StatusCommand{
		Endpoint:   "/camera/sx",
		ResultByte: 7,
		Translaters: []StatusTranslater{
			StatusTranslater{
				Result:         0,
				ExpectedReturn: "wide"},
			StatusTranslater{
				Result:         1,
				ExpectedReturn: "medium"},
			StatusTranslater{
				Result:         2,
				ExpectedReturn: "narrow"}}}
	sc["photores"] = StatusCommand{
		Endpoint:   "/camera/sx",
		ResultByte: 8,
		Translaters: []StatusTranslater{
			StatusTranslater{
				Result:         3,
				ExpectedReturn: "5MP_med"},
			StatusTranslater{
				Result:         4,
				ExpectedReturn: "7MP_wide"},
			StatusTranslater{
				Result:         8,
				ExpectedReturn: "10MP_wide"}}}
	sc["minselapsed"] = StatusCommand{
		Endpoint:    "/camera/sx",
		ResultByte:  13,
		Translaters: nil}
	sc["secselapsed"] = StatusCommand{
		Endpoint:    "/camera/sx",
		ResultByte:  14,
		Translaters: nil}
	sc["volume"] = StatusCommand{
		Endpoint:   "/camera/sx",
		ResultByte: 16,
		Translaters: []StatusTranslater{
			StatusTranslater{
				Result:         0,
				ExpectedReturn: "off"},
			StatusTranslater{
				Result:         1,
				ExpectedReturn: "70%"},
			StatusTranslater{
				Result:         2,
				ExpectedReturn: "100%"}}}
	sc["led"] = StatusCommand{
		Endpoint:   "/camera/sx",
		ResultByte: 17,
		Translaters: []StatusTranslater{
			StatusTranslater{
				Result:         0,
				ExpectedReturn: "off"},
			StatusTranslater{
				Result:         1,
				ExpectedReturn: "led2"},
			StatusTranslater{
				Result:         2,
				ExpectedReturn: "led4"}}}

	/* TODO: Photo/video counts for two bytes? */
	sc["photoremaining"] = StatusCommand{
		Endpoint:    "/camera/sx",
		ResultByte:  21,
		Translaters: nil}
	sc["photocount"] = StatusCommand{
		Endpoint:    "/camera/sx",
		ResultByte:  23,
		Translaters: nil}
	sc["videoremaining"] = StatusCommand{
		Endpoint:    "/camera/sx",
		ResultByte:  25,
		Translaters: nil}
	sc["videocount"] = StatusCommand{
		Endpoint:    "/camera/sx",
		ResultByte:  25,
		Translaters: nil}

	sc["recording"] = StatusCommand{
		Endpoint:   "/camera/sx",
		ResultByte: 29,
		Translaters: []StatusTranslater{
			StatusTranslater{
				Result:         0,
				ExpectedReturn: "off"},
			StatusTranslater{
				Result:         1,
				ExpectedReturn: "on"}}}
	sc["videoresolution"] = StatusCommand{
		Endpoint:   "/camera/sx",
		ResultByte: 50,
		Translaters: []StatusTranslater{
			StatusTranslater{
				Result:         0,
				ExpectedReturn: "WVGA"},
			StatusTranslater{
				Result:         1,
				ExpectedReturn: "720"},
			StatusTranslater{
				Result:         2,
				ExpectedReturn: "960"},
			StatusTranslater{
				Result:         3,
				ExpectedReturn: "1080"}}}
	sc["fps"] = StatusCommand{
		Endpoint:   "/camera/sx",
		ResultByte: 51,
		Translaters: []StatusTranslater{
			StatusTranslater{
				Result:         3,
				ExpectedReturn: "25"},
			StatusTranslater{
				Result:         6,
				ExpectedReturn: "50"}}}

	/*
		/*
			TODO: Need support for checking bits
			sc["orientation"] = StatusCommand{
				Endpoint:   "/camera/sx",
				ResultByte: 18,
				Translaters: []StatusTranslater{
					StatusTranslater{
						Result:         0,
						ExpectedReturn: "up"},
					StatusTranslater{
						Result:         4,
						ExpectedReturn: "down"}}}
			sc["iso_sharpness"] = StatusCommand{
					Endpoint:   "/camera/sx",
					ResultByte: 6,
					Translaters: []StatusTranslater{
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

func (c *Camera) Status(Command string) (string, error) {
	result, err := c.StatusCommands[Command].RunStatusCommand(c.APIRequester)
	if err != nil {
		return "", err
	}
	return result, nil
}
