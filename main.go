package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/dassongh/custom-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hey %s! Starting the interpreter for you\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
