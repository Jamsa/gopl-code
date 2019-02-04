// 指针是实现标准库flag包的关键
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit traling newline")
var sep = flag.String("s", " ", "seperator")

// main 
func main()  {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
