/*
 * gosofunky - demos 3 'functional programming' ways to use functions in go
 */

package main

import "fmt"

/**
 * Way #1: passing functions to other functions as parameters
 *
 * To demo this, let's define a go [:type:] to act as a passed function. 
 * A [:PassFunc:] takes no arguments and returns a string.
 */
type PassFunc func() string

/**
 * Way #1 (cont'd)
 * 
 * Here's what it looks like when we instantiate a [:PassFunc:] type:
 * [:foo:] is a PassFunc.
 */
func foo() string {
	return "I am the return value of PassFunc foo"
}

/**
 * Way #1 (cont'd)
 *
 * Here's the function we're going to pass [:foo:] to later on:
 * [:takesPassFunc:] takes one PassFunc as a parameter.
 */
func takesPassFunc(passfunc PassFunc) {
	fmt.Printf("takesPassFunc: %v\n", passfunc())
}

/**
 * Way #2: functions returning functions
 *
 * Here's a function that returns a function.
 * [:returnsPassFunc:] returns a PassFunc.
 */
func returnsPassFunc() PassFunc {
	return func() string {
		fmt.Printf("I am inside the returnsPassFunc function.\n")
		return "I am the returnsPassFunc return value."
	}
}

func main() {
	fmt.Printf("[1] takesPassFunc(foo)\n")
	// Must not pass [:foo:] as [:foo():] -- that would be a string.
	takesPassFunc(foo)

	fmt.Printf("[2] returnsPassFunc()\n")
	var f PassFunc = returnsPassFunc()
	// Must not pass [:f:] as [:f():] -- that would be a string.
	takesPassFunc(f)

	/**
	 * Way #3 storing an anonymous function in a variable
	 */
	fmt.Printf("[3] stored anonymous PassFunc\n")
	var storeAnon PassFunc = func() string {
		return "I am the anonymous PassFunc\n"
	}
	// Must not pass [:storeAnon():] as [:storeAnon:] -- that would be a PassFunc!
	fmt.Printf(storeAnon())
}
