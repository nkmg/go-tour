package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

/**Exercício: Maps
*	Implementar WordCount. Ele deve retornar um map das contagens de cada "word" na string s.
*	A função wc.Test executa um conjunto de testes contra a função fornecida e imprime o sucesso ou falha.
**/

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")

	map_solve := make(map[string]int)

	for i := 0; i < len(words); i++ {
		if elem, ok := map_solve[words[i]]; ok {
			map_solve[words[i]] = elem + 1
		} else {
			map_solve[words[i]] = 1
		}
	}

	return map_solve
}

func main() {
	wc.Test(WordCount)
}
