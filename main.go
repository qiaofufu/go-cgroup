package main

import (
	"go-cgroup/cgroup"
	"go-cgroup/rescource"
)

func main() {

	cg := cgroup.New([]rescource.IResource{&rescource.CPU{
		Strategy: "test_cpu",
		Quota:    10 * rescource.MS,
		Period:   1 * rescource.MS,
	}})

	cg.AddPid(1001)
}
