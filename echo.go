package main

import (
	"os"
	"fmt"
	"time"
	"strings"
)

func main() {
	measure(a)
	measure(b)
}

func a() {
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func b()  {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

func measure(f func()) {
	start := time.Now()
	f()
	end := time.Now()
	fmt.Println(end.Sub(start))
}