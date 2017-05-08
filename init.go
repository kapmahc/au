package main

// Init init
type Init struct {
}

// Name name
func (p *Init) Name() string { return "init" }

// Usage usage
func (p *Init) Usage() string { return "Generate sample files." }

// Handle handle
func (p *Init) Handle() error { return nil }
