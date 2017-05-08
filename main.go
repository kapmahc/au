package main

import (
	"log"
	"os"
)

func main() {
	commands := []Command{
		&Init{},
		&Version{},
	}
	args := os.Args
	switch len(args) {
	case 2:
		for _, cmd := range commands {
			if cmd.Name() == args[1] {
				if err := cmd.Handle(); err != nil {
					log.Fatal(err)
				}
				return
			}
		}
	case 3:
		env := args[1]
		cmd := args[2]
		log.Printf("run %s[%s]", cmd, env)
		return
	}

	showHelp(commands...)
}
