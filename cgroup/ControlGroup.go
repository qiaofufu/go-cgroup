package cgroup

import "github.com/qiaofufu/go-cgroup/rescource"

type ControlGroup struct {
	r []rescource.IResource
}

func (c *ControlGroup) AddPid(pid int) error {
	for _, v := range c.r {
		err := v.AddPid(pid)
		if err != nil {
			return err
		}
	}
}

func (c *ControlGroup) Delete() error {
	for _, v := range c.r {
		err := v.Delete()
		if err != nil {
			return err
		}
	}
}

func New(resource []rescource.IResource) (*ControlGroup, error) {
	cg := &ControlGroup{r: resource}
	for _, v := range cg.r {
		err := v.Create()
		if err != nil {
			return nil, err
		}

	}
	return cg, nil
}


