package main

import (
	"fmt"
	"math"
	"github.com/jonssonsofia/go_crash_course/03_packages/strutil"
)

func main() {
		fmt.Println(math.Floor(2.7))
		fmt.Println(math.Ceil(2.7))
		fmt.Println(math.Sqrt(16))
		//can't get github package to work
		fmt.Println(strutil.Reverse("olleh"))
}
