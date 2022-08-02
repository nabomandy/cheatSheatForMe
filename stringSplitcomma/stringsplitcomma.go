package main

import (
	"fmt"
	"strings"
)

func main() {
	// string split with comma
	te := "dms-sync-asset-debug, dms-sync-asset-debug, dms-sync-asset-debug, dms-sync-asset-debug, dms-sync-asset-debug"
	fmt.Println(te)
	slice := strings.Split(te, ",")
	fmt.Println(slice)
	// string split with comma

}
