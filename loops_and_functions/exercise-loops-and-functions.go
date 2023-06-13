package main

import (
	"fmt"
)

/** Exercício: Laços e Funções
*	Como forma de brincar com funções e loops, vamos implementar uma função de raiz quadrada: dado um número x, queremos encontrar o número z para o qual z² é quase o x.
*	Os computadores normalmente calculam a raiz quadrada de x usando um loop. Começando com algumas suposições z, podemos ajustar z com base em quão perto z² é x, produzindo uma melhor aproximação:
*
*	z - = (z * z - x) / (2 * z)
*
*	Repetir este ajuste faz com que a aproximação seja cada vez melhor até chegarmos a uma resposta tão próxima da raiz quadrada quanto possível.
*
*	Implemente isto na func Sqrt fornecida. Um palpite inicial aceitável para z é 1, não importa qual seja a entrada. Para começar, repita o cálculo 10 vezes e imprima cada z ao longo do caminho.
*	Veja como você chega perto da resposta para vários valores de x (1, 2, 3, ...) e com que rapidez a aproximação melhora.
*
*	Dica: para declarar e inicializar um valor de ponto flutuante, dê a ela uma sintaxe de ponto flutuante ou use uma conversão:
*
*	z: = 1.0
*	z: = float64 (1)
*
*	Em seguida, altere a condição de loop para parar quando o valor for interrompido mudando (ou apenas muda em uma quantidade muito pequena).
*	Veja se isso é mais ou menos de 10 iterações. Tente outras estimativas iniciais para z, como x ou x/2. Quão perto estão os resultados da sua função para o math.Sqrt da biblioteca padrão?
*
*	(Nota: Se você está interessado nos detalhes do algoritmo, o z² - x acima é o quão longe z² é de onde ele precisa ser (x), e a divisão por 2z é o derivado de z², para dimensionar o quanto ajustamos z pela rapidez com que z² está mudando.
*	Essa abordagem geral é chamada de método de Newton. Ele funciona bem para muitas funções, mas especialmente bem para a raiz quadrada.)
**/

func Sqrt(x float64) float64 {
	z := float64(1)
	for n := 0; n < 10; n++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(9))
}
