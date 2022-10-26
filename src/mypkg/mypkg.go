package mypkg

import "fmt"

type CarPublic struct {
	Brand string
	Year  int
}

// private (inicia con minuscula) - no lo puedes acceder desde otro archivo
type carPrivate struct {
	brand string
	year  int
}

func PrintMessage(texto string) {
	fmt.Println(texto)
}

// private
func printMessage001(texto string) {
	fmt.Println(texto)
}
