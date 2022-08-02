package main

import (
	"fmt"
	"sort"
)

func unique(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func main() {

	urlList := []string{"testv3", "abc", "def", "ghi"}
	remove := []string{"abc", "testv3", "testv4"}

	urlList = append(urlList, "remove...") // remove... 추가
	urlList = append(urlList, remove...)   // [testv3 abc def ghi remove... abc testv3 testv4]
	fmt.Println(urlList)
	fmt.Println("=================================")
	fmt.Println(unique(urlList)) // 결과는 [testv3 abc def ghi remove... testv4]
	fmt.Println("=================================")
	fmt.Println(unique([]string{"abc", "cde", "efg", "efg", "abc", "cde"}))

	a_array := []string{"1", "2", "3", "4"}
	b_array := []string{"3", "4"}

	fmt.Println(diff(a_array, b_array))

}

func diff(a []string, b []string) []string {
	// Turn b into a map
	var m map[string]bool
	m = make(map[string]bool, len(b))
	for _, s := range b {
		m[s] = false
	}
	// Append values from the longest slice that don't exist in the map
	var diff []string
	for _, s := range a {
		if _, ok := m[s]; !ok {
			diff = append(diff, s)
			continue
		}
		m[s] = true
	}
	// Sort the resulting slice
	sort.Strings(diff)
	return diff
}
