package gogopro

import (
	"errors"
	"fmt"
	"io/ioutil"
)

type StatusTranslater struct {
	Result         byte
	ExpectedReturn string
}

type StatusCommand struct {
	Endpoint    string
	ResultByte  int
	Translaters []StatusTranslater
}

func (s StatusCommand) RunStatusCommand(APIRequester *APIRequester) (string, error) {
	resp, err := APIRequester.get(s.Endpoint)
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
	fmt.Println(len(result))
	fmt.Println(result)
	found_result := byte(0)
	if s.ResultByte == -1 {
		found_result = result[len(result)-1]
	} else {
		found_result = result[s.ResultByte]
	}

	for _, translater := range s.Translaters {
		if translater.Result == found_result {
			return translater.ExpectedReturn, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Failed to find a expected return val for result %d", found_result))
}
