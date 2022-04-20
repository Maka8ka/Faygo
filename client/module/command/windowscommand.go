package commandimport (
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)
type WindowsCommand struct {
}
func NewWindowsCommand() *WindowsCommand {
	return &WindowsCommand{}
}
func (lc *WindowsCommand) Exec(args ...string) (int, string, error) {
	args = append([]string{"/c"}, args...)
	cmd := exec.Command("cmd", args...)	cmd.SysProcAttr = &syscall.SysProcAttr{}	outpip, err := cmd.StdoutPipe()
	defer outpip.Close()	if err != nil {
		return 0, "", err
	}	err = cmd.Start()
	if err != nil {
		return 0, "", err
	}	out, err := ioutil.ReadAll(outpip)
	if err != nil {
		return 0, "", err
	}
	cmdout := ConvertByte2String(out, "GB18030")
	return cmd.Process.Pid, string(cmdout), nil
}
func (lc *WindowsCommand) ExecAsync(stdout chan string, args ...string) int {
	var pidChan = make(chan int, 1)	go func() {
		args = append([]string{"/c"}, args...)
		cmd := exec.Command("cmd", args...)		cmd.SysProcAttr = &syscall.SysProcAttr{}		outpip, err := cmd.StdoutPipe()
		defer outpip.Close()		if err != nil {
			panic(err)
		}		err = cmd.Start()
		if err != nil {
			panic(err)
		}		pidChan <- cmd.Process.Pid		out, err := ioutil.ReadAll(outpip)
		if err != nil {
			panic(err)
		}		stdout <- string(out)
	}()	return <-pidChan
}
func (lc *WindowsCommand) ExecIgnoreResult(args ...string) error {
	args = append([]string{"/c"}, args...)
	cmd := exec.Command("cmd", args...)	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{}	err := cmd.Run()	return err
}
