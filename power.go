package gogopro

import ()

type Power struct {
	APIRequester   *APIRequester
	StatusCommands map[string]StatusCommand
}

func (p *Power) Init() *Power {
	return p
}

func CreatePower(APIRequester *APIRequester) *Power {
	power := &Power{}
	power.APIRequester = APIRequester
	statusCommands := CreatePowerStatusCommands()
	power.StatusCommands = statusCommands
	return power
}

func CreatePowerStatusCommands() map[string]StatusCommand {
	sc := make(map[string]StatusCommand)
	sc["power"] = StatusCommand{Endpoint: "/bacpac/se", ResultByte: -1,
		Translaters: []StatusTranslater{
			StatusTranslater{
				Result:         0,
				ExpectedReturn: "off"},
			StatusTranslater{
				Result:         1,
				ExpectedReturn: "on"}}}
	return sc
}

func (p *Power) Status(Command string) (string, error) {
	result, err := p.StatusCommands[Command].RunStatusCommand(p.APIRequester)
	if err != nil {
		return "", err
	}
	return result, nil
}
