/*
** AUTHOR	: Robert James Patterson
** DATE		: 09/17/2019
** SYNOPSIS	: Standard output is more or less equivalent to printing on the screen. However, using
** standard output might require the use of functions that do not belong to the fmt package, which
** is why it is presented in its own section.
 */
package main

import (
	"io"
	"os"
)

func main() {

	myString := ""
	arguments := os.Args

	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}

	// Now we are using the WriteString function from the io package. In this case, the
	// io.WriteString() function works in the same way as the fmt.Print() function; however, it takes
	// only two parameters. The first parameter is the file to which you want to write, which in this
	// case is os.Stdout, and the second parameter is a string variable.
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
