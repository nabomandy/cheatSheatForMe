package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/buger/jsonparser"
	"github.com/djimenez/iconv-go"
	"github.com/everystreet/go-shapefile"
	geojson "github.com/paulmach/go.geojson"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

func main() {
	//file, err := os.Open("C:\\Users\\USER\\Desktop\\T1\\AL_11110_D198_20220118.zip")
	file, err := os.Open("C:\\Users\\zeus\\Desktop\\오늘\\용도별건물정보\\AL_11500_D198_20220118.zip")
	fmt.Println(err)
	stat, err := file.Stat()
	fmt.Println(err)

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
		record := scanner.Record()
		if record == nil { // 종료 조건
			break
		}
		feature := record.GeoJSONFeature()
		testT1, _ := json.Marshal(feature)
		collection, _ := geojson.UnmarshalFeature(testT1)
		p := collection.Properties
		V1 := fmt.Sprintf("%v", p)
		fmt.Println(V1)
		V2 := fmt.Sprintf("%v", p["A6"])
		fmt.Println(V2)
		read, _ := iconv.ConvertString(p["A6"].(string), "utf-8", "utf-8")
		fmt.Println(read)

		polygon := collection.Geometry.Polygon
		polygonResult := ""
		for _, multiPolygon := range polygon[0] {
			polygonResult += fmt.Sprintf(", %v %v", multiPolygon[0], multiPolygon[1])
			// Each record contains a shape (from .shp file) and attributes (from .dbf file)
		}

		//for _, feature := range collection.Features {
		//	fmt.Println(feature.Geometry)
		//	fmt.Println(feature.Properties)
		//}
		geometry, err := json.Marshal(feature.Geometry)
		properties, err := json.Marshal(feature.Properties)

		fmt.Println(err)
		//fmt.Println(string(jsonData))
		coordinates, _, _, _ := jsonparser.Get(geometry, "coordinates")
		property, _, _, _ := jsonparser.Get(properties)
		fmt.Println(property)
		//get2, _, _, _ := jsonparser.Get(jsonData, "geometry")
		//fmt.Println(get2)
		var t1 [][][]float64
		_ = json.Unmarshal(coordinates, &t1)
		//polygonResult := ""
		for _, multiPolygon := range t1[0] {
			polygonResult += fmt.Sprintf(", %v %v", multiPolygon[0], multiPolygon[1])
			// Each record contains a shape (from .shp file) and attributes (from .dbf file)
			fmt.Println(record)
		}
		/*
			e.g
			, 197119.187 454372.508, 197103.397 454350.578, 197091.126 454359.418, 197106.927 454381.348, 197119.187 454372.508
		*/

		// Each record contains a shape (from .shp file) and attributes (from .dbf file)
		fmt.Println(record)
	}

	// Err() returns the first error encountered during calls to Record()
	err = scanner.Err()
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