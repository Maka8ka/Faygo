package command

import (
	"runtime"
)

type Commander interface {
	Exec(args ...string) (int, string, error)
	ExecAsync(stdout chan string, args ...string) int
	ExecIgnoreResult(args ...string) error
}

func NewCommand() Commander {
	var cmd Commander

	switch runtime.GOOS {
	case "linux":
		cmd = NewLinuxCommand()
	case "windows":
		cmd = NewWindowsCommand()
	default:
		cmd = NewLinuxCommand()
	}

	return cmd
}
