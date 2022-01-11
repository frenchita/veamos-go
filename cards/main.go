package main

import (
	"fmt"
)

type cards []card

func main() {

	my_cards := cards{}

	for {
		action := StringPrompt("Ingrese el número de la acción que desea realizar\n1 - Listar Cards\n2 - Agregar Card\n3 - Modificar Card\n4 - Eliminar Card\n5 - Salir\n")
		switch action {
		case "1":
			my_cards.show()
		case "2":
			my_cards = my_cards.add()
		case "3":
			fmt.Println("3")
		case "4":
			fmt.Println("4")
		case "5":
			fmt.Println("Saliendo ...")
		default:
			fmt.Printf("default")
		}

		if action == "5" {
			break
		}
	}

}
