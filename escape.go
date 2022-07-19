package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

func main() {

	origin := "%uAC00%uB098%uB2E4%uB77C%3F%26+aa"

	s := unescape(origin)

	fmt.Printf("unescape('%s') = %s \n", origin, s)

	e := escape(s)

	fmt.Printf("escape('%s') = %s \n", s, e)

	fmt.Printf("%s == %s : %v\n", origin, e, origin == e)

}

func unescape(s string) string {
	hex2 := regexp.MustCompile("(?i)^[\\da-f]{2}$")
	hex4 := regexp.MustCompile("(?i)^[\\da-f]{4}$")

	str := []rune(s)
	result := make([]rune, 0)
	length := len(str)
	index := 0
	var chr rune
	var part string
	for index < length {
		chr = str[index]
		index++
		if chr == '%' {
			if str[index] == 'u' {
				part = string(str[index+1 : index+5])
				if hex4.Match([]byte(part)) {
					data, _ := hex.DecodeString(part)
					r := rune(binary.BigEndian.Uint16(data))
					result = append(result, r)
					index += 5
					continue
				}
			} else {
				part = string(str[index : index+2])
				if hex2.Match([]byte(part)) {
					data, _ := hex.DecodeString(part)
					r := rune(data[0])
					result = append(result, r)
					index += 2
					continue
				}
			}
		}
		result = append(result, chr)
	}

	return string(result)
}

func escape(s string) string {
	raw := regexp.MustCompile("[\\w*+\\-./@]")
	str := []rune(s)
	var result = ""
	var length = len(str)
	var index = 0
	var chr rune
	for index < length {
		chr = str[index]
		index++
		if raw.Match([]byte(string(chr))) {
			result += string(chr)
		} else {
			if int32(chr) < 256 {
				result += "%" + strings.ToUpper(hex.EncodeToString([]byte{byte(chr)}))
			} else {
				l := byte(chr & 0xFF)
				h := byte((chr >> 8) & 0xFF)
				result += "%u" + strings.ToUpper(hex.EncodeToString([]byte{h, l}))
			}
		}
	}
	return result
}
