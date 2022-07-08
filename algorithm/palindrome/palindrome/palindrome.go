package palindrome

import "fmt"

func main() {
	var str string = "kasur ini rusak"
	data := false
	for n := 0; n < len(str)/2; n++ {
		if str[n] != str[len(str)-n-1] {
			data = false
		} else {
			data = true
		}
	}
	fmt.Println(data)
}
