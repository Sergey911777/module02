package main

import (
	"fmt"
	"strconv"
)

func main() {
	chysl := 97
	str := "104"
	strNew, _ := strconv.Atoi(str)

	fmt.Println(strconv.Itoa(chysl), strNew)
}
