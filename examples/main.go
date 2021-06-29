package main

import (
	"fmt"
	"os"

	"github.com/nickgerace/qexec"
)

func run() error {
	var cmd string
	var args []string

	cmd, args = qexec.PrepareInput("")
	if err := qexec.Exec(cmd, args...); err == nil {
		return fmt.Errorf("this should have failed")
	}
	cmd, args = qexec.PrepareInput("go help build")
	if err := qexec.Exec(cmd, args...); err != nil {
		return err
	}
	if err := qexec.Exec("echo", "hello world"); err != nil {
		return err
	}

	// This can be simplified by returning the last function call, but is separated for readability.
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
