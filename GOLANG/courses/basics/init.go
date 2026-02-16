package main

import "fmt"


/*
	init() en Go, c’est un hook d’initialisation automatique : le runtime l’appelle tout seul, avant main(), après l’initialisation des variables de package.
*/
func init() {
	fmt.Println("Initializing package1...")
}

func init() {
	fmt.Println("Initializing package2...")
}

func init() {
	fmt.Println("Initializing package3...")
}

func main() {

	fmt.Println("Inside the main function")

}