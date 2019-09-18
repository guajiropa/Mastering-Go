/*
** AUTHOR	: Robert James Patterson
** DATE		: 09/15/2019
** SYNOPSIS	: Go offers excellent support for string formatting in the printf tradition. Here are some
** examples of common string formatting tasks.
 */
package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func main() {

	// Create a variable of type 'point'.
	p := point{5, 9}

	myint := 123
	myfloat := 123.456
	myArg := ""



	// This prints an instance of our point struct, using the %v verb.
	fmt.Printf("The point is at : %v\n", p)

	// If the value is a struct, the %+v variant will include the struct’s field names.
	fmt.Printf("The point is at : %+v\n", p)

	// The %#v variant prints a Go syntax representation of the value, i.e. the source code snippet
	// that would produce that value.
	fmt.Printf("The point is at : %#v\n", p)

	// To print the type of a value, use %T.
	fmt.Printf("The point of type : %T\n", p)

	// Formatting booleans is straight-forward.
	fmt.Printf("P is of type point: %t\n", true)

	// Use %d for standard, base-10 formatting.
	fmt.Printf("This is an integer value using the %%d verb: %d\n", myint)

	// This prints the binary representation of the same value.
	fmt.Printf("The %%b verb is the same number as a binary : %b\n", myint)

	// This prints the character corresponding to the given integer.
	fmt.Printf("Use the %%c verb to get the Unicode chacter for 'myint' : %c\n", myint)

	// %x provides hex encoding.
	fmt.Printf("You can use %%x to get the same number in hex encoding : %x\n", myint)

	// There are also several formatting options for floats. For basic decimal formatting use %f.
	fmt.Printf("A basic decimal number using the %%f verb : %f\n", myfloat)

	// You can specify the number of percision points in a float using %.2f
	fmt.Printf("You can set the persicion using the %%.2f verb : %.2f\n", myfloat)

	// You can specify the whole width of the float using the the %<n>.<n>f verb
	fmt.Printf("You can specify the width with the verb notation %%8.2f : %8.2f\n", myfloat)

	myBigFloat := 123.456789101112131415116171809
	// Use %e and %E format the float in (slightly different versions of) scientific notation.
	fmt.Printf("You can use %%e for large floats to get scientific notation %e\n", myBigFloat)
	fmt.Printf("You can use %%E for large floats to get scientific notation %E\n", myBigFloat)

	// 

}
