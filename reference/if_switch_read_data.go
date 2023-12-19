package main

import (
	"fmt"
	"log"
	"os"
)

/*
	we can create vars inside the if and they
	are just available on that range of function
	(is good to save memory)
*/

func main() {
	a, b := 1, 2

	if a > b {
		fmt.Println("a > b")
	} else if b > a {
		fmt.Println("b > a")
	}
	fmt.Println("b = a")

	/**/
	file, err := os.Open("hello.txt")
	if err != nil { /* nil is a representation of a null value*/
		log.Panic(err)
	}

	data := make([]byte, 100) /* the make create a array of byte with {x} bytes*/

	/*read the file and output inside the array*/

	if _, err := file.Read(data);

	/* the function return the amount of bytes and if return a error, with if _ we
	   ignore the amount of bytes verification
	*/

	err != nil {
		log.Panic(err)
	}

	/*we convert the data to string because he is a array*/
	fmt.Println(string(data))
}
