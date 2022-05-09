package main 
import(
	"fmt"
	"regexp"
)
func main() {
	str1 := "Thats a nice joke ðŸ˜†ðŸ˜†ðŸ˜† ðŸ˜›"
	emojiRx := regexp.MustCompile(`[\x{1F600}-\x{1F6FF}|[\x{2600}-\x{26FF}]`)
	str := emojiRx.ReplaceAllString(str1, `[e]`) 
	fmt.Println(str) 
}