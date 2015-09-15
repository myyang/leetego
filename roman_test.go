package leetego

import (
	"errors"
	"testing"
)

func TestEncode(t *testing.T) {
	ro := Roman{}
	if v, err := ro.Encode(1234); v != "MCCXXXIV" && err != nil {
		t.Error("Error to encode 1234")
	}
	if v, err := ro.Encode(4000); v != "4000" && err != errors.New("Value should be 1~3999") {
		t.Error("Error to encode 4000")
	}
	if v, err := ro.Encode(-1); v != "-1" && err != errors.New("Value should be 1~3999") {
		t.Error("Error to encode -1")
	}
}

func TestDecode(t *testing.T) {
	ro := Roman{}
	if v, err := ro.Decode("MMMCDLXVI"); v != 3966 && err != nil {
		t.Error("Error to decode MMMCDLXVI")
	}

	if v, err := ro.Decode("ABCDE"); v != 0 && err != errors.New("Invalid charater: A") {
		t.Error("Error to decode ABCDE")
	}

	if v, err := ro.Decode("CXIB"); v != 0 && err != errors.New("Invalid charater: B") {
		t.Error("Error to decode CXIB")
	}
}

func TestSetDecodable(t *testing.T) {
	ro := Roman{}
	if ro.Decodable != false {
		t.Error("Roamn.Decodable is not false by default")
	}
	ro.SetDecodable(true)
	if ro.Decodable != true {
		t.Error("Fail to set Roman.Decodable through Roman.SetDecodable")
	}
}
