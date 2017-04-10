package gogopro

import (
	"testing"
)

func TestEvalTranslator(t *testing.T) {
	s := StatusCommand{
		ResultByte:  1,
		Translaters: map[byte]string{0: "video", 1: "photo"},
	}
	expected := s.Translaters[1]
	r, _ := s.EvalTranslater([]byte{0, 1, 2})
	if r != expected {
		t.Errorf("Expected %s, got %s", expected, r)
	}
}

func TestNoMatchEvalTranslaters(t *testing.T) {
	s := StatusCommand{
		ResultByte:  0,
		Translaters: map[byte]string{},
	}
	r, _ := s.EvalTranslater([]byte{0})
	if r != "" {
		t.Errorf("Expected empty result, got %s", r)
	}
}

func TestEvalTranslaterLastValue(t *testing.T) {
	s := StatusCommand{
		ResultByte:  -1,
		Translaters: map[byte]string{0: "no", 1: "yes"},
	}
	expected := s.Translaters[0]
	r, _ := s.EvalTranslater([]byte{0, 3, 1, 2, 3, 1, 0})
	if r != expected {
		t.Errorf("Expected %s , got %s", expected, r)
	}
}
