package variables //el package es el nombre de la carpeta

import "fmt"

func MuestroEnteros() {
	var intcomun int     // aca intcomun vale cero
	intde32 := int32(10) // asignacion
	intde64 := int64(100)

	fmt.Println("intcomun = ", intcomun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)
}
