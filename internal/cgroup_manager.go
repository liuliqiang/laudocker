package internal

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/liuliqiang/laudocker/internal/g"
)

var SubsystemsIns = []g.SubSystem{}

type CgroupManager struct {
	Path     string
	Resource *g.ResourceConfig
}

func NewCgroupManager(path string) *CgroupManager {
	return &CgroupManager{
		Path: path,
	}
}

func (c *CgroupManager) Apply(pid int) error {
	for _, subSysIns := range SubsystemsIns {
		subSysIns.Apply(c.Path, pid)
	}
	return nil
}

func (c *CgroupManager) Set(res *g.ResourceConfig) error {
	for _, subSysIns := range SubsystemsIns {
		subSysIns.Set(c.Path, res)
	}
	return nil
}

func (c *CgroupManager) Destroy() error {
	for _, subSysIns := range SubsystemsIns {
		if err := subSysIns.Remove(c.Path); err != nil {
			logrus.Warnf("remove cgroup fail %v", err)
		}
	}
	return nil
}

func FindCgroupMountpoint(subsystem string) string {
	f, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return ""
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		fields := strings.Split(txt, " ")
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subsystem {
				return fields[4]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return ""
	}
	return ""
}

func GetCgroupPath(subsystem string, cgroupPath string, autoCreate bool) (string, error) {
	cgroupRoot := FindCgroupMountpoint(subsystem)
	if _, err := os.Stat(path.Join(cgroupRoot, cgroupPath)); err == nil || (autoCreate && os.IsNotExist(err)) {
		if os.IsNotExist(err) {
			if err := os.Mkdir(path.Join(cgroupRoot, cgroupPath), 0755); err != nil {
				return "", fmt.Errorf("error create cgroup %v", err)
			}
		}
		return path.Join(cgroupRoot, cgroupPath), nil
	} else {
		return "", fmt.Errorf("cgroup path error %v", err)
	}
}
