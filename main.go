package main

import (
	"fmt"
	"log"
	"os"

	fs "fs/ascii"
)

func main() {
	if len(os.Args) != 3 {
		log.Println("Argument Error !")
	}
	fmt.Println(fs.FinalPrint(os.Args[1], os.Args[2]))
}
