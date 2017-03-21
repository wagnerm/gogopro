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
	return sc
}

func (c *Camera) GetMode() (string, error) {
	result, err := c.StatusCommands["mode"].RunStatusCommand(c.APIRequester)
	if err != nil {
		return "", err
	}
	return result, nil
}
