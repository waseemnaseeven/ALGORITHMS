package main
	
import (
    "fmt"
    "os"
    "path/filepath"
)

/*
    3 regles d'or du defer:
        1. les defer s'exec en LIFO : le dernier defer rencontre s'exec en premier
        2. les arguments sont evalues tout de suite : c’est “capturer maintenant, exécuter plus tard”.
        3. Ça s’exécute même si la fonction sort par return ou par panic. En cas de panic, Go “déroule la pile” (stack unwinding) et exécute les defers au passage.

    Exception: (voir exit.go) dans le cas du Os.Exit() 
*/

func process(i int) {
	defer fmt.Println("Deffered i value:", i)
	defer fmt.Println("First deferred statement executed")
	defer fmt.Println("Second deferred statement executed")
	defer fmt.Println("Third deferred statement executed")
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i:", i)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}
func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close()
    if err != nil {
        panic(err)
    }
}

func main() {

	process(10)
	path := filepath.Join(os.TempDir(), "defer.txt")
	f := createFile(path)
	defer closeFile(f)
	writeFile(f)
}


