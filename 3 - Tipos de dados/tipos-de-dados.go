package main

import (
	"fmt"
	"errors"
)

func main() {
	var numero1 int = 2222222
	fmt.Println(numero1)

	// inferencia de tipos
	numero := 10000000
	fmt.Println(numero)

	var numero3 uint32 = 333333
	fmt.Println(numero3)

	// não suportavar numero4 uint32 = -444444
	//fmt.Println(numero4)

	//alias
	//INT32 = RUNE
	var numero5 rune = 55555555
	fmt.Println(numero5)

	//BYTE = UINT8  >>>. aguenta só dois caracteres
	var numero6 byte = 66
	fmt.Println(numero6)


	// float >> quebrados

	var numeroReal1 float32 = 4444.4444
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1234321.321213  // não pode por apenas float ... exlicitamente
	fmt.Println(numeroReal2)

	//inferencia de tipos >> o sistema já designa o tipo de acordo com a arquitetura do seu windows
	numeroReal3 := 14.45
	fmt.Println(numeroReal3)


	// FIM NUMEROS REAIS


	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)


	char := '7' // vsai dar resultado da tabela ASK desse caracter | // rune = int do "char :=""
	fmt.Println(char)



	//STRINGS

	var texto string // retornar vazio pois não foi declaro valor
	fmt.Println(texto)

	var texto2 int16 // retornar 0 pois não foi declaro valor, valor inicial que é 0
	fmt.Println(texto2)


	// retonar nill que é erro se houver

	var casa string = "pow"
	fmt.Println(casa)


	pai := 5
	fmt.Println(pai) // aqui tem que declarar valor de qualquer jeito


	// FIM DE STRINGS


	// BOOLEANS


	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool
	fmt.Println(booleano2) // sem declarar valor, automaticamente é false



	// Tipo ERRO
	
	var erro error //erro nome da variavel // error nome do tipo da variavel 
	fmt.Println(erro) // nil de 0, de erro


	var errado error = errors.New("Erro interno")  // errors nome do pacote dessa variavel
	fmt.Println(errado)
}