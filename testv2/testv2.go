package main

import "fmt"

func main() {

	urlList := []string{"testv3", "abc", "def", "ghi"}
	remove := []string{"abc", "testv3"}

loop:
	for i, url := range urlList {
		for _, rem := range remove {
			if url == rem {
				urlList = append(urlList[:i], urlList[i+1:]...)
				continue loop
			}
		}
	}

	for _, v := range urlList {
		fmt.Println(v)
	}
}
