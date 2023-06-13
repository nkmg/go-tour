package main

import "golang.org/x/tour/pic"

/**Exercício: Slices
*	Implemente Pic. Ela deve retornar uma slice de comprimento dy, cada elemento do qual é uma slice de dx 8-bit inteiros sem sinal.
*	Quando você executa o programa, ele irá exibir a sua imagem, interpretando os números inteiros como escala dos valores de cinza (bem, bluescale).
*
*	A escolha da imagem é com você. Funções interessantes incluem x^y, (x+y)/2, e x*y.
*
*	(Você precisa usar um loop para alocar cada []uint8 dentro do [][]uint8.)
*
*	(Utilize uint8(intValue) para converter entre os tipos.)
**/

func Pic(dx, dy int) [][]uint8 {
	the_slice_figure := make([][]uint8, dy) // crie uma slice com dy de comprimento e que cada elemento e uma slice []uint8

	for i := 0; i < dy; i++ { //loop para alocar cada []uint8
		the_slice_aux := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			the_slice_aux[j] = uint8(i * j)
		}
		the_slice_figure[i] = the_slice_aux
	}
	return the_slice_figure
}

func main() {
	pic.Show(Pic)
}
