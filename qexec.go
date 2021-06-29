package qexec

import (
	"fmt"
	"strings"

	"github.com/nickgerace/qexec/internal/command"
)

func PrepareInput(v interface{}) (string, []string) {
	s, ok := v.(string)
	if !ok {
		s = fmt.Sprint(v)
	}
	cmd := strings.Fields(s)
	if len(cmd) < 1 {
		return "", nil
	} else if len(cmd) == 1 {
		return cmd[0], nil
	}
	return cmd[0], cmd[1:]
}

func Exec(cmd string, args ...string) error {
	return command.Exec(cmd, args...)
}
