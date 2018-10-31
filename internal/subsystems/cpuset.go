package subsystems

import "github.com/liuliqiang/laudocker/internal/g"

type CpusetSubsystem struct {
}

func (*CpusetSubsystem) Name() string {
	return "cpuset"
}

func (*CpusetSubsystem) Set(path string, res *g.ResourceConfig) error {
	panic("implement me")
}

func (*CpusetSubsystem) Apply(path string, pid int) error {
	panic("implement me")
}

func (*CpusetSubsystem) Remove(path string) error {
	panic("implement me")
}
