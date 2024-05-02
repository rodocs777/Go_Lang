package main

import "fmt"

/* como se fosse uma receita de bolo, serie de passos, instruoes pelo seu programa, que irao fazer
o que ele manda 

elas podem ter parâmetros e retorno

parametros o que poe para ela funcionar e retorno o que poe para retorno
*/


func somar(n1 int8, n2 int8) int8 {
	return n1 + n2

}

func calculosMatematicos(n1 int8, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)



	////////////////////////////////////////


	var f = func() {
		fmt.Println("funcao f")
	}

	f()

	/////////////////////////////


	var t = func(txt string) {
		fmt.Println(txt)
	}

	t("Texto da funcao T")


	////// abaixo ela pode ter um retorno, ao inves de mostrar texto na tela


	var e = func(txt string) string /*retorno*/ { 
		return txt
	}

	resultado := e("Texto da funcao E")
	fmt.Println(resultado)


	//////////////////// funcoes podem ter mais de um retorno | duas variaveis, dois retornos

	resultadosSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadosSoma, resultadoSubtracao)

	
	// caso eu não queia que um dos resultados apareca, coloco _ no caso para ignorar o segundo

	resultadosMais, _ := calculosMatematicos(7, 7)
	fmt.Println(resultadosMais)


	// caso eu não queia que um dos resultados apareca, coloco _ no caso para ignorar o primeiro

	_ , resultadosMenos := calculosMatematicos(7, 7)
	fmt.Println(resultadosMenos)
}
