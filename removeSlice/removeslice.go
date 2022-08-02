package main

import "fmt"

func main() {
	urlList := []string{"testv3", "abc", "def", "ghi"}
	remove := []string{"abc", "testv3"}

loop:
	for i := 0; i < len(urlList); i++ {
		url := urlList[i]
		for _, rem := range remove {
			if url == rem {
				urlList = append(urlList[:i], urlList[i+1:]...)
				i-- // Important: decrease index
				continue loop
			}
		}
	}

	fmt.Println(urlList)

}
