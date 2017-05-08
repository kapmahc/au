package main

import "fmt"

func showHelp(args ...Command) {
	for _, arg := range args {
		fmt.Printf("%s\t\t%s\n", arg.Name(), arg.Usage())
	}
	fmt.Println("help\t\tDisplay available options.")
}
