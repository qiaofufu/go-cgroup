package rescource

import "testing"

func TestCPU_Create(t *testing.T) {
	c := CPU{
		Strategy: "test_cpu",
		Quota:    10 * MS,
		Period:   1 * MS,
	}

	if err := c.Create(""); err != nil {
		t.Fatal("create cpu cgroup fail.", err.Error())
	}
}
