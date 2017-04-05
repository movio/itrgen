package itrgen

import (
	"io/ioutil"
	"os"
)

func write(file string, code []byte) error {
	if err := deleteIfExists(file); err != nil {
		return err
	}

	return ioutil.WriteFile(file, code, 0644)
}

func deleteIfExists(file string) error {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return nil
	}

	return os.Remove(file)
}
