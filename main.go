package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	nombre :=
		`
Bienvenido al juego de cara y sello 

Introduce tu nombre:

`
	fmt.Print(nombre)
	fmt.Scanln(&nombre)

	fmt.Println("Tu nombre es", nombre)

	menu :=
		`
	Elije
	[ 0 ] Cara
	[ 1 ] Sello
	¿Qué prefieres?
`
	fmt.Print(menu)

	var eleccion int
	fmt.Scanln(&eleccion)

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 1
	resultado := rand.Intn(max-min+1) + min

	if resultado == eleccion {
		fmt.Println("Has ganado")
	} else {
		fmt.Println("Has Perdido")
	}

	i := 0
	fmt.Print("Seleccione [0] para salir o [1] para seguir jugando  \n")
	fmt.Scanln(&i)
	switch i {
	case 0:
		fmt.Println("Fin del Juego")
	case 1:
		fmt.Println("Nueva Partida")

		nombre :=
			`
Bienvenido al juego de cara y sello 

Introduce tu nombre:

`
		fmt.Print(nombre)
		fmt.Scanln(&nombre)

		fmt.Println("Tu nombre es", nombre)

		menu :=
			`
	Elije
	[ 0 ] Cara
	[ 1 ] Sello
	¿Qué prefieres?
`
		fmt.Print(menu)

		var eleccion int
		fmt.Scanln(&eleccion)

		rand.Seed(time.Now().UnixNano())
		min := 0
		max := 1
		resultado := rand.Intn(max-min+1) + min

		if resultado == eleccion {
			fmt.Println("Has ganado")
		} else {
			fmt.Println("Has Perdido")
		}

		i := 0
		fmt.Print("Seleccione [0] para salir o [1] para seguir jugando  \n")
		fmt.Scanln(&i)
		switch i {
		case 0:
			fmt.Println("Fin del Juego")
		case 1:
			fmt.Println("Nueva Partida")

		}

	}

}
