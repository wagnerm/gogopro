package gogopro

import (
	"testing"
)

func NewStatusCommand(ResultByte int) *StatusCommand {
	return &StatusCommand{
		Endpoint:    "",
		ResultByte:  ResultByte,
		Translaters: NewStatusTranslaters(),
	}
}

func NewStatusTranslaters() []StatusTranslater {
	t := []StatusTranslater{
		StatusTranslater{
			Result:         0,
			ExpectedReturn: "low"},
		StatusTranslater{
			Result:         1,
			ExpectedReturn: "med"},
		StatusTranslater{
			Result:         2,
			ExpectedReturn: "high"},
	}
	return t
}

func RunEvalTranslater(ResultByte int, Results []byte) (string, error) {
	s := NewStatusCommand(ResultByte)
	return s.EvalTranslater(Results)
}

func TestEvalTranslator(t *testing.T) {
	if r, err := RunEvalTranslater(0, []byte{0, 1, 2}); err != nil {
		t.Error(err)
	} else if r != "low" {
		t.Errorf("Expected low, got %s", r)
	}

	if r, err := RunEvalTranslater(1, []byte{0, 1, 2}); err != nil {
		t.Error(err)
	} else if r != "med" {
		t.Errorf("Expected med, got %s", r)
	}

	if r, err := RunEvalTranslater(2, []byte{0, 1, 2}); err != nil {
		t.Error(err)
	} else if r != "high" {
		t.Errorf("Expected high, got %s", r)
	}
}
