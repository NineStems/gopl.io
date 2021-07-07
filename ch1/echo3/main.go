// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

//!+
func main() {
	fmt.Println(EchoEx1(os.Args))
	fmt.Println(EchoEx2(os.Args))
}

//!-

/*
	ex from book
*/
func EchoEx1(args []string) string {
	return strings.Join(args, " ")
}

func EchoEx2(args []string) string {
	res := ""
	for i := 0; i < len(args); i++ {
		res += fmt.Sprintf("%v %s \n", i, args[i])
	}
	return  res
}

func echoEx3Concat(args []string) {

}

func echoEx3Join(args []string) {

}
