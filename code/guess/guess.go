package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	low := 1.0
	high := 100.0
	tries := 0

	fmt.Printf(
		"Pense num numero entre %d e %d que eu vou tentar adivinhar\n\n",
		int(low),
		int(high),
	)

	for {
		guess := (low + high) / 2.0

		fmt.Printf("O número que você pensou foi %d?\n", int(guess))

		fmt.Println("Certo?")
		fmt.Println("(a) Não. Ele é maior")
		fmt.Println("(b) Não. Ele é menor")
		fmt.Println("(c) Isso mesmo!")
		scanner.Scan()

		input := scanner.Text()
		if input == "a" {
			low = math.Floor(guess + 1)
			tries++
		} else if input == "b" {
			high = math.Ceil(guess - 1)
			tries++
		} else if input == "c" {
			fmt.Printf("Hehehe sou foda. Acertei em %d tentativas :b\n", tries)
			break
		} else {
			fmt.Println("Entrada inválida")
		}

		if low > high {
			fmt.Println("Acho que você está tentando me engabelar.")
			fmt.Println("Não brinco mais com você >,<")
			break
		}
	}
}
