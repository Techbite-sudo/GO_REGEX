package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "Input the html content here containing the images "	 

	re := regexp.MustCompile(`]+\bsrc=["']([^"']+)["']`)

	submatchall := re.FindAllStringSubmatch(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element[1])
	}
}