package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		fmt.Println("Example 1")
		time.Sleep(1 * time.Second)

		f, _ := os.OpenFile("example1.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		f.WriteString("Example1\n")
		f.Close()
	}

}
