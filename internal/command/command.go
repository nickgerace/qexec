package command

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/nickgerace/qexec/internal/platform"
)

func Exec(command string, args ...string) error {
	if command == "" {
		return fmt.Errorf("must use non-empty command")
	}
	cmd := exec.Command(command, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		stdoutScanner := bufio.NewScanner(stdout)
		for stdoutScanner.Scan() {
			fmt.Fprintf(os.Stdout, "%s%s", stdoutScanner.Text(), platform.NEWLINE)
		}
	}()

	foundStderr := false

	wg.Add(1)
	stderrScanner := bufio.NewScanner(stderr)
	go func() {
		defer wg.Done()
		for stderrScanner.Scan() {
			foundStderr = true
			fmt.Fprintf(os.Stderr, "%s%s", stderrScanner.Text(), platform.NEWLINE)
		}
	}()

	if err := cmd.Wait(); err != nil {
		return err
	}
	wg.Wait()

	if foundStderr {
		return fmt.Errorf("found message(s) in stderr for command: %q", cmd.Args)
	}
	return nil
}
