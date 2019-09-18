/*
**
**
**
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	userinput := bufio.NewScanner(f)

	for userinput.Scan() {
		fmt.Println(">", userinput.Text())
	}

}
