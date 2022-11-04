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
	Strategy string
	Quota    uint `name:"cpu.cfs_quota_us"`
	Period   uint `name:"cpu.cfs_period_us"`
}

// Create dir
func (c *CPU) Create() error {
	// make dir
	dirPath := path.Join(BaseDirCPU, c.Strategy)
	if err := os.Mkdir(dirPath, 0750); err != nil && !os.IsExist(err) {
		return err
	}
	t := reflect.TypeOf(*c)
	v := reflect.ValueOf(*c)
	err := writeToCGroupFile(t, v, dirPath, os.O_TRUNC|os.O_WRONLY)
	if err != nil {
		return err
	}
	return nil
}

func (c *CPU) AddPid(pid int) error {
	dirPath := path.Join(BaseDirCPU, c.Strategy)
	err := addPid(dirPath, pid)
	if err != nil {
		return err
	}
	return nil
}
