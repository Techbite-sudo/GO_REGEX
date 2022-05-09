package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	str1 := "#Golang#Python$Php&Kotlin@@"
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	newStr := reg.ReplaceAllString(str1,"-")
	fmt.Println(newStr)
}