package rescource

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestMemory_Create(t *testing.T) {
	mem := &Memory{
		Strategy:      "test_memory",
		LimitInMemory: 512 * MB,
		Swappiness:    0,
	}

	if err := mem.Create(""); err != nil {
		t.Fatal(fmt.Sprintf("%v create memory cgroup fail.", err))
	}
}

func TestMemory_AddPid(t *testing.T) {
	mem := &Memory{
		Strategy:      "test_memory",
		LimitInMemory: 10 * MB,
		Swappiness:    0,
	}

	if err := mem.Create(""); err != nil {
		t.Fatal(fmt.Sprintf("%v create memory cgroup fail.", err))
	}

	pid := os.Getpid()

	if err := mem.AddPid(pid); err != nil {
		t.Fatal(fmt.Sprintf("%v add memory pid fail.", err))
	}

	time.Sleep(time.Second * 5)
}
