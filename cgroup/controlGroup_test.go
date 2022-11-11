package cgroup

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	cg, err := New("test")
	if err != nil || cg == nil {
		t.Fatal(errors.New("new cgroup error"))
	}
}
