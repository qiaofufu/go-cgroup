package rescource

import (
	"errors"
	"fmt"
	"os"
	"path"
	"reflect"
)

type IResource interface {
	Create() error
	AddPid(pid int) error
	Delete() error
}

func writeToCGroupFile(t reflect.Type, v reflect.Value, dirPath string, flag int) error {
	for i := 0; i < t.NumField(); i++ {
		filename := t.Field(i).Tag.Get("name")

		var val []byte
		switch v.Field(i).Kind() {
		case reflect.String:
			val = []byte(v.Field(i).String())
		case reflect.Uint:
			val = []byte(fmt.Sprintf("%d", v.Field(i).Uint()))

		default:
			return errors.New(fmt.Sprintf("%s not is string or uint", v.Field(i).Kind().String()))
		}

		if len(val) != 0 && filename != "" {
			err := writeToFile(path.Join(dirPath, filename), val, flag, 0750)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func writeToFile(name string, data []byte, flag int, perm os.FileMode) error {
	file, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return errors.New(fmt.Sprintf("open file error :%v", err))
	}
	_, err = file.Write(data)
	if err != nil {
		return errors.New(fmt.Sprintf("write file error :%v", err))
	}
	if err := file.Close(); err != nil {
		return errors.New(fmt.Sprintf("close file eroor: %v", err))
	}
	return nil
}

func addPid(dirPath string, pid int) error {
	filePath := path.Join(dirPath, "cgroup.procs")

	err := writeToFile(filePath, []byte(fmt.Sprintf("%d", pid)), os.O_APPEND|os.O_WRONLY, 0750)
	if err != nil {
		return err
	}
	return nil
}
