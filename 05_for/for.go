// `for` é a única estrutura de looping do Go. Aqui estão
// três tipos básicos de `for` loops.

package main

import "fmt"

func main() {

	// O tipo mais comum, com uma única condição.
	i := 0
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	// Uma clássica inicial/condição/final `for` loop.
	for j := 6; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` sem uma condição se repetirá várias vezes
	// até que você dê um `break` no loop ou `return` da
	// função de fechamento.
	for {
		fmt.Println("loop")
		break
	}
}
