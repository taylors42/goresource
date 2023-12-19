package main

import "fmt"

type People struct {
	Age        int
	Name       string
	StillAlive bool

	/* if the property name start in lowercase is private, in Upper is public*/
}

func main() {

	taylor := People{
		Age:        18,
		Name:       "taylor",
		StillAlive: true,
	}
	fmt.Println(taylor.Age)
}
