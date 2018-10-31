package container

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/liuliqiang/laudocker/internal/g"

	"github.com/liuliqiang/laudocker/internal"

	"github.com/sirupsen/logrus"
)

func Run(tty bool, command string, res *g.ResourceConfig) {
	parent := NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		logrus.Error(err)
	}

	cgroupManager := internal.NewCgroupManager("laudocker-cgroup")
	defer cgroupManager.Destroy()
	cgroupManager.Set(res)
	cgroupManager.Apply(parent.Process.Pid)

	parent.Wait()
	os.Exit(-1)
}

func NewParentProcess(tty bool, command string) *exec.Cmd {
	args := []string{"init", command}
	logrus.Debugf("new parent process: init %s", command)
	cmd := exec.Command("/proc/self/exe", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}

	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd
}

func RunContainerInitProcess(command string, args []string) (err error) {
	logrus.Infof("container command: %s", command)

	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	argv := []string{command}
	if err = syscall.Exec(command, argv, os.Environ()); err != nil {
		logrus.Error(err.Error())
	}

	return nil
}
