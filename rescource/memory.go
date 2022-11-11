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
	dirPath       string
	LimitInMemory uint `filename:"memory.limit_in_bytes"`
	Swappiness    uint `filename:"memory.swappiness"`
}

func (m *Memory) Create(name string) error {
	m.dirPath = path.Join(baseDirMemory, name)
	if err := os.Mkdir(m.dirPath, 0750); err != nil && !os.IsExist(err) {
		return err
	}
	t := reflect.TypeOf(*m)
	v := reflect.ValueOf(*m)
	err := writeToCGroupFile(t, v, m.dirPath, os.O_TRUNC|os.O_WRONLY)
	if err != nil {
		return err
	}
	return nil
}

func (m *Memory) AddPid(pid int) error {
	err := addPid(m.dirPath, pid)
	if err != nil {
		return err
	}
	return nil
}

func (m *Memory) Delete() error {
	if err := os.RemoveAll(m.dirPath); err != nil {
		return err
	}
	return nil
}
