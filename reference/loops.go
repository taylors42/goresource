package main

import (
	"fmt"
)

func main() {
	nomes := []string{"a", "b", "c"}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < len(nomes); i++ {
		fmt.Println(nomes[i])
	}

	for everynameofkeywhoyouwant, nome := range nomes {
		fmt.Println(everynameofkeywhoyouwant, nome)
		fmt.Println(nome)
	}

	for {
		fmt.Println("loop infinito")
	}
}
