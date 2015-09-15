// Package leetego provides roman convertion related funtions
package leetego

import (
	"errors"
	"strconv"
	"strings"
)

type RomanLevels struct {
	Unit, Five, Ten string
	Value           int
}

var (
	UNIT    RomanLevels = RomanLevels{"I", "V", "X", 1}
	TENS    RomanLevels = RomanLevels{"X", "L", "C", 10}
	HUND    RomanLevels = RomanLevels{"C", "D", "M", 100}
	THOU    string      = "M"
	NUMERIC             = map[string]int{
		"M": 1000,
		"C": 100, "D": 500,
		"X": 10, "L": 50,
		"I": 1, "V": 5,
	}
	CORRECTION = map[string]int{
		"CM": -200, "CD": -200,
		"XC": -20, "XL": -20,
		"IV": -2, "IX": -2,
	}
)

type Roman struct {
	Decodable bool
}

func (ro Roman) Decode(s string) (r int, err error) {
	t := 0
	s = strings.ToUpper(s)
	for _, c := range s {
		if v, err := NUMERIC[string(c)]; err == true {
			t += v
		} else {
			return 0, errors.New("Invalid charater: " + string(c))
		}
	}
	for k, v := range CORRECTION {
		if strings.Contains(s, k) {
			t -= v
		}
	}
	return t, nil
}

func conv(v int, roi RomanLevels) (r string, rem int) {

	b, x := v/roi.Value, v%roi.Value
	switch b {
	case 0, 1, 2, 3:
		return strings.Repeat(roi.Unit, b), x
	case 4:
		return roi.Unit + roi.Five, x
	case 5, 6, 7, 8:
		return roi.Five + strings.Repeat(roi.Unit, (b-5)), x
	case 9:
		return roi.Unit + roi.Ten, x
	}
	return r, rem
}

func (ro Roman) Encode(n int) (r string, err error) {
	if 0 > n || n > 3999 {
		return strconv.FormatInt(int64(n), 10), errors.New("Value should be 1~3999")
	}

	r, en_s, v, rem := "", "", n/1000, n%1000
	r += strings.Repeat(THOU, v)
	for _, roi := range []RomanLevels{HUND, TENS, UNIT} {
		en_s, rem = conv(rem, roi)
		r += en_s
	}
	return r, nil
}

func (ro *Roman) SetDecodable(b bool) {
	ro.Decodable = b
}
