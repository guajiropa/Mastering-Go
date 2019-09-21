/*
** AUTHOR	: Robert James Patterson
** DATE		: 09/17/2019
** SYNOPSIS	: Reading keys from the os.Stdio as a file
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

	keyreader := bufio.NewScanner(f)
	for keyreader.Scan() {
		fmt.Println(">", keyreader.Text())
	}
}
