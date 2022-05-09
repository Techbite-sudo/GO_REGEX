package main

import (
	"fmt"
	"regexp"
)
func main(){
	str1 := "We @@@Love@@@@ #Go!$! ****Programming****Language^^^"
	re :=regexp.MustCompile(`[^a-zA-Z0-9]+`)
	fmt.Println("Pattern:",re.String())
	fmt.Println("Matched:",re.MatchString(str1))

	submatchall := re.FindAllString(str1,-1)
	for _,element := range submatchall{
		fmt.Println(element)
	}
}