package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//for {
	for {
		fmt.Println("Example 3")
		time.Sleep(1 * time.Second)

		f, _ := os.OpenFile("example3.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		f.WriteString("Example3\n")
		f.Close()

	}
}
