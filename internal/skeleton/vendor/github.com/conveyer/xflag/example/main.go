// +build ignore

// Package main demonstrates the use `xflag` and `xflag/cflag`.
package main

import (
	"flag"
	"log"

	"github.com/conveyer/xflag"
	"github.com/conveyer/xflag/cflag"
)

var (
	name = flag.String("name", "John Doe", "Name of the user.")
	age  = flag.Int("age", 16, "Age of the user.")

	path = flag.String("paths:src", "/home/smbd", "Path to the Go sources.")

	// emails uses "cflag" package that provides support of
	// complex flags: slices of strings, ints, floats, etc.
	emails = cflag.Strings("emails[]", []string{"test@test.xx"}, "A list of e-mails.")
)

func main() {
	// Parse flags: use 3 different configuration files.
	err := xflag.Parse("./file1.ini", "./file2.ini", "./file3.ini")
	if err != nil {
		log.Fatal(err)
	}

	// Print the values of flags.
	log.Printf("Name: `%s`.", *name)
	log.Printf("Age: %d.", *age)
	log.Printf("Path: `%s`.", *path)
	log.Printf("Emails: %v.", *emails)
}
