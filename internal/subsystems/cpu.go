package subsystems

import "github.com/liuliqiang/laudocker/internal/g"

type CpuSubsystem struct {
}

func (*CpuSubsystem) Name() string {
	return "cpu"
}

func (*CpuSubsystem) Set(path string, res *g.ResourceConfig) error {
	panic("implement me")
}

func (*CpuSubsystem) Apply(path string, pid int) error {
	panic("implement me")
}

func (*CpuSubsystem) Remove(path string) error {
	panic("implement me")
}
