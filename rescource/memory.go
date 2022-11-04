package rescource

import (
	"os"
	"path"
	"reflect"
)

const (
	baseDirMemory = "/sys/fs/cgroup/memory"
	MB            = KB * 1024
	KB            = 1024
)

type Memory struct {
	Strategy      string
	LimitInMemory uint `filename:"memory.limit_in_bytes"`
	Swappiness    uint `filename:"memory.swappiness"`
}

func (m *Memory) Create() error {
	dirPath := path.Join(baseDirMemory, m.Strategy)
	if err := os.Mkdir(dirPath, 0750); err != nil && !os.IsExist(err) {
		return err
	}
	t := reflect.TypeOf(*m)
	v := reflect.ValueOf(*m)
	err := writeToCGroupFile(t, v, dirPath, os.O_TRUNC|os.O_WRONLY)
	if err != nil {
		return err
	}
	return nil
}

func (m *Memory) AddPid(pid int) error {
	dirPath := path.Join(baseDirMemory, m.Strategy)
	err := addPid(dirPath, pid)
	if err != nil {
		return err
	}
	return nil
}
