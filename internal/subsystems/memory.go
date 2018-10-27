package subsystems

import "github.com/liuliqiang/laudocker/internal/g"

type MemorySubSystem struct {
}

func (s *MemorySubSystem) Name() string {
	return "memory"
}

func (*MemorySubSystem) Set(path string, res *g.ResourceConfig) error {

}

func (*MemorySubSystem) Apply(path string, pid int) error {
	panic("implement me")
}

func (*MemorySubSystem) Remove(path string) error {
	panic("implement me")
}
