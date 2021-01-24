package command

import (
	"runtime"
)


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
