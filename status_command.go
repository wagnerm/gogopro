package gogopro

import (
	"errors"
	"fmt"
	"io/ioutil"
)

type StatusCommand struct {
	ResultByte  int
	Translaters map[byte]string
}

func (s StatusCommand) RunStatusCommand(Endpoint string, APIRequester *APIRequester) (string, error) {
	resp, err := APIRequester.get(Endpoint)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	result, err := s.EvalTranslater(body)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s StatusCommand) EvalTranslater(result []byte) (string, error) {
	commandResultByte := byte(0)
	if s.ResultByte == -1 {
		commandResultByte = result[len(result)-1]
	} else if s.ResultByte < len(result) && s.ResultByte >= 0 {
		commandResultByte = result[s.ResultByte]
	}

	v, ok := s.Translaters[commandResultByte]
	if ok == false {
		return v, errors.New(fmt.Sprintf("No related status for %d", commandResultByte))
	}

	return v, nil
}
