package leetego

import (
	"errors"
	"testing"
)

func TestMorseEncode(t *testing.T) {
	mo := Morse{}
	expected_str, expected_err := ".... . .-.. .-.. --- / .-- --- .-. .-.. -..", errors.New("Invalid charater: \\")
	if v, err := mo.Encode("hello world"); v != expected_str && err != nil {
		t.Error("error to encode: hello world")
	}
	if v, err := mo.Encode("\\"); v != "" && err != expected_err {
		t.Error("error to encode: \\")
	}
}

func TestMorseDecode(t *testing.T) {
	mo := Morse{}
	expected_str, expected_err := "HELLO WORLD", errors.New("Invalid charater: .......")
	if v, err := mo.Decode(".... . .-.. .-.. --- / .-- --- .-. .-.. -.."); v != expected_str && err != nil {
		t.Error("error decode to HELLO WORLD ")
	}
	if v, err := mo.Decode("......."); v != "" && err != expected_err {
		t.Error("error to decode: .......")
	}
}

func TestMorseSetDecodable(t *testing.T) {
	mo := Morse{}
	if mo.Decodable != false {
		t.Error("Morse.Decodable is not false by default")
	}
	mo.SetDecodable(true)
	if mo.Decodable != true {
		t.Error("Fail to set Morse.Decodable through Morse.SetDecodable")
	}
}
