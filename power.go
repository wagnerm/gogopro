package gogopro

import (
	"io/ioutil"
)

type Power struct {
	APIRequester *APIRequester
}

func (p *Power) Init() *Power {
	return p
}

func CreatePower(APIRequester *APIRequester) *Power {
	power := &Power{}
	power.APIRequester = APIRequester
	return power
}

func (p *Power) GetPowerStatus() (string, error) {
	resp, err := p.APIRequester.get("/bacpac/se")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if body[len(body)-1] == 0 {
		return "off", nil
	}
	return "on", nil
}
