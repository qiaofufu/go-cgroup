package rescource

import (
	"os"
	"path"
	"reflect"
)

const (
	BaseDirCPU = "/sys/fs/cgroup/cpu"
	MS         = 1000
)

type CPU struct {
	dirPath string
	Quota   uint `filename:"cpu.cfs_quota_us"`
	Period  uint `filename:"cpu.cfs_period_us"`
}

// Create dir
func (c *CPU) Create(name string) error {
	// make dir
	c.dirPath = path.Join(BaseDirCPU, name)
	if err := os.Mkdir(c.dirPath, 0750); err != nil && !os.IsExist(err) {
		return err
	}
	t := reflect.TypeOf(*c)
	v := reflect.ValueOf(*c)
	err := writeToCGroupFile(t, v, c.dirPath, os.O_TRUNC|os.O_WRONLY)
	if err != nil {
		return err
	}
	return nil
}

func (c *CPU) AddPid(pid int) error {
	err := addPid(c.dirPath, pid)
	if err != nil {
		return err
	}
	return nil
}

func (c *CPU) Delete() error {
	if err := os.RemoveAll(c.dirPath); err != nil {
		return err
	}
	return nil
}
