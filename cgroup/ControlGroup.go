package cgroup

import "go-cgroup/rescource"

type ControlGroup struct {
	r []rescource.IResource
}

func (c ControlGroup) AddPid(pid int) {
	for _, v := range c.r {
		v.AddPid(pid)
	}
}

func New(resource []rescource.IResource) *ControlGroup {
	cg := &ControlGroup{r: resource}
	for _, v := range cg.r {
		v.Create()
	}
	return cg
}
