package main

import "fmt"

// Version init
type Version struct {
}

// Name name
func (p *Version) Name() string {
	return "version"
}

// Usage usage
func (p *Version) Usage() string { return "Display the program version." }

// Handle handle
func (p *Version) Handle() error {
	fmt.Println("2017.05.04")
	return nil
}
