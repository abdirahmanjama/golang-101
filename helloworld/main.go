package main

import "fmt"

func main() {
	fmt.Println(2 * 2)
	fmt.Print(2 * 5)
}

//go run main.go

//go build main.go complies our application into an executable. We can run this executable by running ./[filename]

/*
Package:

A package is a directory/module which contains a collection of Go files.
We typically store cohesive code here. Each file within a package must declare it's package at the top of the file.

Why do call packages main?

Inside of Go there are two type of packages: Executable and Reusable.

Executable generates a file that we can run.

Reusable packages are a great place to keep reusable code. Keeps our code clean and modular.

We called this package "main". This package states that it is an executable.

The name of the package determines if its an executable or reusable file.



*/
