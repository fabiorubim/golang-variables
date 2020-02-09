package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Caso uma variável não seja utilizada mas foi declarada, o código não irá compilar
	var nome string = "Fabio Rubim"
	var versao float32 = 1.1
	var idade int = 32                                  //Se não for atribuído algum valor, o valor será zero.
	var cidade = "Sorocaba"                             //Go utiliza também inferência de tipos.
	fmt.Println("Olá, Sr.", nome, "sua idade é", idade) //concatenação
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("Sua cidade é", cidade)

	var altura = 1.79 //Utilizará a inferência, mas no caso de um float, pode ser float32 ou float64, para descobri qual é o tipo utilize reflect.typeOf
	fmt.Println("Sua altura é", altura)
	fmt.Println("O tipo da variável altura é", reflect.TypeOf(altura)) //no caso ele retorna que é um tipo float64. Por padrão ele coloca o maior tipo

	//Encurtando a declaração de variáveis. Operador de declaração curta de variáveis
	//Como no Pascal pode-se utilizar o :=, ele irá inferir e atribuir ao mesmo. Assim não é necessário utilizar a palavra var
	endereco := "Rua XYZ, Nº569"
	fmt.Println("Seu endereço é", endereco)
}
