package main

import (
	"fmt"
	"os"
)

/*
	os.Exit(code) tue le process immédiatement (au niveau OS), avec un code de sortie.
	Il ne “retourne” pas de main, il ne “laisse pas Go finir proprement” : 
	il demande au système d’exploitation d’arrêter le programme tout de suite.
*/

func main() {

	defer fmt.Println("Deferred statement")

	fmt.Println("Starting the main function")

	// Exit with status code of 1
	os.Exit(1)

	// This will never be executed
	fmt.Println("End of main function")

}