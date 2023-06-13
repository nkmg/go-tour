package main

import "fmt"

/**Exercício: Fibonacci closure
*	Vamos nos divertir um pouco com as funções.
*	Implemente uma função fibonacci que retorna uma função (um closure) que retorna números sucessivos de Fibonacci (0, 1, 1, 2, 3, 5, ...).
**/

func fibonacci() func() int {
	fibonacci_aux_2, fibonacci_aux_1 := 0, 1

	return func() int {
		fibonacci_number := fibonacci_aux_2
		fibonacci_aux_2, fibonacci_aux_1 = fibonacci_aux_1, fibonacci_number+fibonacci_aux_1
		return fibonacci_number
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
