package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
)

const (
	// TASKS tasks dir
	TASKS = "tasks"
	// CONFIG config dir
	CONFIG = "config"
	// STAGE stage env
	STAGE = "stage"
	// EXT json ext
	EXT = ".json"
)

// Init init
type Init struct {
}

// Name name
func (p *Init) Name() string { return "init" }

// Usage usage
func (p *Init) Usage() string { return "Generate sample files." }

// Handle handle
func (p *Init) Handle() error {
	err := p.generateConfig()
	if err == nil {
		err = p.generateTasks()
	}
	return err
}

func (p *Init) generateTasks() error {
	fn := path.Join(TASKS, "init.tpl")
	log.Println("generate file", fn)
	fd, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = fd.WriteString(`
{{define "status"}}
uname -a
{{end}}

`)
	return err
}

func (p *Init) generateConfig() error {
	for _, n := range []string{CONFIG, TASKS} {
		if _, err := os.Stat(n); err == nil {
			return fmt.Errorf("%s already exists", n)
		}
		if err := os.Mkdir(n, 0700); err != nil {
			return err
		}
	}
	// --------
	host, err := os.Hostname()
	if err != nil {
		return err
	}
	// --------
	usr, err := user.Current()
	if err != nil {
		return err
	}
	// --------
	fn := path.Join(CONFIG, STAGE+EXT)
	log.Println("generate file", fn)
	fd, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer fd.Close()
	// --------

	enc := json.NewEncoder(fd)
	enc.SetIndent("", "  ")
	return enc.Encode(map[string]interface{}{
		"key": path.Join(usr.HomeDir, ".ssh", "id_rsa"),
		"roles": map[string]interface{}{
			"www": []interface{}{
				"deploy@www.change-me.com",
			},
			"db": []interface{}{
				"deploy@db.change-me.com",
			},
			"app": []interface{}{
				"deploy@app.change-me.com",
			},
		},
		"args": map[string]interface{}{
			"env":  STAGE,
			"host": host,
			"user": usr.Username,
			"home": usr.HomeDir,
		},
	})
}
