package cgroup

import "github.com/qiaofufu/go-cgroup/rescource"

func New(path string) (*ControlGroup, error) {
	cg := &ControlGroup{Name: path}

	return cg, nil
}

type ControlGroup struct {
	Name string
	r    []rescource.IResource
}

func (c *ControlGroup) AddPid(pid int) error {
	for _, v := range c.r {
		err := v.AddPid(pid)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *ControlGroup) Delete() error {
	for _, v := range c.r {
		err := v.Delete()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *ControlGroup) AddResource(resources ...rescource.IResource) error {
	c.r = resources
	for _, v := range c.r {
		err := v.Create("")
		if err != nil {
			return err
		}
	}
	return nil
}
