package main

import (
	"fmt"
	"github.com/spf13/cast"
	"os"
	"reflect"
	"strconv"

	"github.com/everystreet/go-shapefile"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

func main() {
	//file, err := os.Open("C:\\Users\\USER\\Desktop\\T1\\AL_11110_D198_20220118.zip")
	file, err := os.Open("C:\\Users\\zeus\\Desktop\\오늘\\용도별건물정보\\AL_11110_D198_20220118.zip")
	fmt.Println(err)
	stat, err := file.Stat()
	fmt.Println(err)
	index := 0

	// Create new ZipScanner
	// The filename can be replaced with an empty string if you don't want to check filenames inside the zip file
	//scanner := shapefile.NewZipScanner(file, stat.Size(), "ne_110m_admin_0_sovereignty.zip")
	scanner, err := shapefile.NewZipScanner(file, stat.Size(), stat.Name())

	// Optionally get file info: shape type, number of records, bounding box, etc.
	info, err := scanner.Info()
	fmt.Println(info)

	// Start the scanner
	err = scanner.Scan()

	// Call Record() to get each record in turn, until either the end of the file, or an error occurs
	for {
		if index == 561 {
			fmt.Println("561")
		}
		record := scanner.Record()
		if record == nil {
			break
		}
		feature := record.GeoJSONFeature()
		for idx := 0; idx < len(feature.Properties); idx++ {
			a := reflect.TypeOf(feature.Properties[idx].Value)
			if fmt.Sprint(a) == "string" {
				value := feature.Properties[idx].Value.(string)
				feature.Properties[idx].Value = ansiToUniString(value)
			}
		}
		//fmt.Println(record.Field("A20"))
		//feature := record.GeoJSONFeature()

		//_, _ := json.Marshal(feature)
		//fmt.Println(len(jsonData))

		// Each record contains a shape (from .shp file) and attributes (from .dbf file)
		//fmt.Println(record)
		index++
		//fmt.Println(index)
	}

	fmt.Println(index)

	// Err() returns the first err
	//r encountered during calls to Record()
	err = scanner.Err()
	fmt.Println(err)
}

func UniToAnsi(src []byte) []byte {
	got, _, _ := transform.String(korean.EUCKR.NewEncoder(), string(src))

	return []byte(got)
}

func AnsiToUni(src []byte) string {
	got, _, _ := transform.String(korean.EUCKR.NewDecoder(), string(src))

	return got
}

func AnsiToUniString(src string) string {
	got, _, _ := transform.String(korean.EUCKR.NewDecoder(), src)
	return got
}

func AnsiToUniStringV2(src interface{}) string {
	got, _, _ := transform.String(korean.EUCKR.NewDecoder(), cast.ToString(src))
	strconv.ParseFloat(got, 64)
	return got
}

func ansiToUniString(src string) string {
	got, _, _ := transform.String(korean.EUCKR.NewDecoder(), src)

	return got

}
