package main

import "fmt"
import "regexp"
import "time"

func main() {
	t := time.Now()
	s := fmt.Sprintf("%v", t)
	fmt.Println("S:", s)
	fmt.Println("time:", t)

	const pat = ".*[0-9]+"

	if ok, _ := regexp.Match(pat, []byte(s)); ok {
		fmt.Println("Match found");
	}

// replace num to *
	re, err := regexp.Compile("")
	if err != nil {
		fmt.Println(err)
		return
	}

	ns := re.ReplaceAllString(s, "*")
	fmt.Println("Replaced:", ns)
}
