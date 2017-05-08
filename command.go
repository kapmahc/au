package main

// Command command
type Command interface {
	Name() string
	Usage() string
	Handle() error
}
