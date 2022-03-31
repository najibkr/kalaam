package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/najibkr/kalaam/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Zabaan programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(user.Username, os.Stdin, os.Stdout)
}
