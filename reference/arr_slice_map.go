package main

import (
	"fmt"
)

func main() {

	/* slice is a dynamic list, in go arrays got fixed size*/
	slice := []string{"a", "b", "c"} /* this slice got the space to 3 items, but can increase */

	/*
		slice_set_size_on_declare := make([]string, 10, 20)
		 he can got 10 items inside he and is prepared to more 10
	*/

	fmt.Println(slice)
	slice = append(slice, "d") /*include the d on arr*/
	fmt.Println("the items inside slice is", slice, "and you length is", len(slice))

	/*
		the hash table is a data structure who implement a collection of keys
		and every key is mapped to give access to a specific data
	*/
	age := make(map[string]int8)
	age["tay"] = 18
	age["joao"] = 20
	/*age[key] = 1 , to show I can use print (age[key] ) */
	fmt.Println(age["tay"]) /* if you write something don´t exist, the return is 0*/
	val, ok := age["tay"]
	fmt.Println(val, ok) /*return: 18 + true*/
	/*
		this function return 2 numbers, the data and a boolean,
		and if don't exist the key the return is 0, false
		but in a simple print we just se the data because
		don´t use another parameter

	*/
}
