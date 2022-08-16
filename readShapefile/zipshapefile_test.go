package main

import (
	"bytes"
	"fmt"
	"github.com/djimenez/iconv-go"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"testing"
)

func TestAnsiToUni(t *testing.T) {
	origin := "\ufffd\ufffd\ufffdհ\ufffd\ufffd๰"
	fmt.Println(origin)

	toUniString := AnsiToUniString(origin)
	fmt.Println(toUniString)

	iconvConvert, _ := iconv.ConvertString(origin, "euc-kr", "utf-8")
	fmt.Println(iconvConvert)

	iconvConvertV2, _ := iconv.ConvertString(origin, "cp949", "utf-8")
	fmt.Println(iconvConvertV2)

	iconvConvertV4, _ := iconv.ConvertString(origin, "ISO-8859-1", "utf-8")
	fmt.Println(iconvConvertV4)

	iconvConvertV3, _ := iconv.ConvertString(origin, "utf-8", "utf-8")
	fmt.Println(iconvConvertV3)

	var bufs bytes.Buffer
	wr := transform.NewWriter(&bufs, korean.EUCKR.NewEncoder())
	wr.Write([]byte(origin))
	wr.Close()

	fmt.Println(bufs.String())

	charSet := []string{"utf-8", "euc-kr", "ksc5601", "iso-8859-1", "x-windows-949"}
	for _, i := range charSet {
		out, _ := iconv.ConvertString(origin, i, "utf-8")
		fmt.Println(i, " : ", out)
	}

}
