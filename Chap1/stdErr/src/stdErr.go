/*
** AUTHOR	: Robert James Patterson
** DATE		: 10/08/2019
** SYNOPSIS	:
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
		myString = "Please provide one argument!"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, "This is standard output\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")	
 }