package main

import "fmt"

func main() {
	urlList := []string{"test", "abc", "def", "ghi"}
	remove := []string{"abc", "test"}

	for i := 0; i < len(urlList); i++ {
		url := urlList[i]
		for _, rem := range remove {
			if url == rem {
				urlList = append(urlList[:i], urlList[i+1:]...)
				i-- // Important: decrease index
			}
		}
	}

	fmt.Println(urlList)

}
