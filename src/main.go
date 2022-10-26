package main

import (
	"curso1/src/mypkg" //se puede crear un alis asi:  pkgalias "curso1/src/mypkg"
	"fmt"
	"log"
	"strconv"
)

type car struct {
	brand string
	year  int
}

type pc struct {
	ram   int
	disk  int
	brand string
}

func (varPC pc) ping() {
	fmt.Println(varPC.brand, "Pong")
}

func (varPC *pc) duplicateRAM() {
	varPC.ram = varPC.ram * 2
}

func normalFunction(numero int, mens, mens2 string) {
	fmt.Println("Epale viejo number:", numero, mens, mens2)
}

func dobleReturn(a, b int) (c, d int) {
	return a + b, a - b
}

func main() {
	//Constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("Hello, World . . . GO GO GO ! ! !")
	fmt.Println(pi, pi2)

	//Variable
	base := 12
	var altura int = 14
	var area int = base * altura

	fmt.Println(base, altura, "- area cuadrado:", area)

	//Paquete fmt
	helloMsj := "Hello"
	worldMsj := "World"
	cant := 500
	fmt.Println(helloMsj, worldMsj)
	fmt.Println(helloMsj + worldMsj)

	fmt.Printf("%s es muchas veces, mas de %d veces mas grande \n", worldMsj, cant)
	//%v si no tengo certeza del dato
	fmt.Printf("%v es muchas veces, mas de %v veces mas grande \n", worldMsj, cant)

	message := fmt.Sprintf("%s es muchas veces, mas de %d veces mas grande", worldMsj, cant)
	fmt.Println(message)

	//Tipo de datos
	fmt.Printf("%T \n", helloMsj)

	//Funciones
	normalFunction(400, "que", "es lo ques")

	x, y := dobleReturn(10, 4)
	fmt.Println(x, y)

	xx, _ := dobleReturn(20, 8) //con el _ puedo obviar uno de los valores
	fmt.Println(xx)

	//CICLOS
	//For condicional
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Printf("\n\n")

	//For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	//For forever (Debes detenerlo con Ctrl+C)
	// counterForever := 0
	// for {
	// 	fmt.Println(counterForever)
	// 	counterForever++
	// }

	//OPERADORES LOGICOS SON && || !
	//defer ejecuta la instruccion de ultimo, antes de morir el programa o funcion
	//continue y break, keywords para el ciclo for

	// IF
	val1 := 1
	// val2 := 2
	if val1 == 1 {
		fmt.Println("IF es 1")

	} else {
		fmt.Println("IF no es 1")
	}

	//Convertir texto a numero
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)

	// ARRAY - los array son inmutables, osea, no les puedo agregar o quitar valores (Tuple en Python?)
	var miarray [4]int
	miarray[2] = 8
	fmt.Println(miarray, len(miarray), cap(miarray))

	// SLICE - le puedo agregar o quitar valores (List en Python)
	mislice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(mislice, len(mislice), cap(mislice))
	fmt.Println(mislice[2], mislice[:3], mislice[2:4], mislice[4:])

	mislice = append(mislice, 9)
	fmt.Println(mislice)

	newslice := []int{10, 11, 12}
	mislice = append(mislice, newslice...) //los 3 puntos desmantela el slice para hacer el append
	fmt.Println(mislice)

	strinSlice := []string{"hola", "que", "hace"}
	for i := range strinSlice {
		fmt.Println(i, strinSlice[i])
	}
	for i, valor := range strinSlice {
		fmt.Println(i, valor)
	}

	//Llave-Valor MAP (dict en Python)
	m := make(map[string]int)
	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m, m["Jose"])

	//otra forma de declarar el map
	temp := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	fmt.Println(temp, temp["Earth"])

	//recorre map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//encontrar valor en map
	value, ok := m["Jose"]
	fmt.Println(value, ok)

	value, ok = m["pepito"]
	fmt.Println(value, ok)

	//Type STRUCT - Va fuera de la funcion
	myCar := car{brand: "ford", year: 2020}
	fmt.Println(myCar)

	var otherCar car
	otherCar.brand = "toyota"
	fmt.Println(otherCar)

	//Tipo de Acceso: si empieza con letra mayuscula es publico, minuscula es privado
	var miOtroCar mypkg.CarPublic
	miOtroCar.Brand, miOtroCar.Year = "Ferrari", 2020
	fmt.Println(miOtroCar)

	mypkg.PrintMessage("Hola cocacola")

	//STRUCTS Y PUNTEROS
	aa := 50
	bb := &aa //puntero - direccion de memoria de aa
	fmt.Println(aa, bb, *bb)

	*bb = 100            //le cambia el valor a todas las variables q esten apuntando a esa direccion
	fmt.Println(aa, *bb) //se usa para funciones customizadas y llevar funciones de un paquete a otro

	//ej de uso, se crea struct pc y funcion del struc
	myPC := pc{ram: 16, disk: 200, brand: "IBM"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC.ram)
	myPC.duplicateRAM()
	fmt.Println(myPC.ram)

	//Listas e Interfaces
	myInterface := []interface{}{1, "abc", true}
	fmt.Println(myInterface...)
}
