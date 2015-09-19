// Package leetego provides moorse convertion related funtions
package leetego

import (
	"errors"
	"strings"
)

var ENCODE = map[string]string{
	"A":  ".-",
	"B":  "-...",
	"C":  "-.-.",
	"D":  "-..",
	"E":  ".",
	"F":  "..-.",
	"G":  "--.",
	"H":  "....",
	"I":  "..",
	"J":  ".---",
	"K":  "-.-",
	"L":  ".-..",
	"M":  "--",
	"N":  "-.",
	"O":  "---",
	"P":  ".--.",
	"Q":  "--.-",
	"R":  ".-.",
	"S":  "...",
	"T":  "-",
	"U":  "..-",
	"V":  "...-",
	"W":  ".--",
	"X":  "-..-",
	"Y":  "-.--",
	"Z":  "--..",
	" ":  "/",
	"1":  ".----",
	"2":  "..---",
	"3":  "...--",
	"4":  "....-",
	"5":  ".....",
	"6":  "-....",
	"7":  "--...",
	"8":  "---..",
	"9":  "----.",
	"0":  "-----",
	".":  ".-.-.-",
	":":  "---...",
	",":  "--..--",
	";":  "-.-.-.",
	"?":  "..--..",
	"=":  "-...-",
	"\"": ".----.",
	"'":  ".-..-.",
	"/":  "-..-.",
	"!":  "-.-.--",
	"-":  "-....-",
	"_":  "..--.-",
	"(":  "-.--.",
	")":  "-.--.-",
	"@":  ".--.-.",
	"$":  "...-..-",
	"&":  ".-...",
	"+":  ".-.-.",
}

func reverseEncode() map[string]string {
	n := make(map[string]string)
	for k, v := range ENCODE {
		n[v] = k
	}
	return n
}

var (
	DECODE = reverseEncode()
	SPLIT  = " "
)

type Morse struct {
	Decodable bool
}

func (mo *Morse) SetDecodable(b bool) {
	mo.Decodable = b
}

func (mo Morse) Encode(s string) (r string, err error) {
	s = strings.ToUpper(s)
	l := len(s) - 1
	for index, value := range s {
		mvalue, ok := ENCODE[string(value)]
		if ok {
			r += mvalue
			if index != l {
				r += string(SPLIT)
			}
		} else {
			return "", errors.New("Invalid charater: " + string(value))
		}
	}
	return r, nil
}

func (mo Morse) Decode(s string) (r string, err error) {
	set := strings.Split(s, SPLIT)
	for _, value := range set {
		mvalue, ok := DECODE[value]
		if ok {
			r += mvalue
		} else {
			return "", errors.New("Invalid charater: " + value)
		}
	}
	return r, nil
}
