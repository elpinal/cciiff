package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	return nil
}

func clang(src string) error {
	file, err := ioutil.TempFile("", "cciiff")
	if err != nil {
		return errors.Wrap(err, "create a temporary file")
	}
	defer os.Remove(file.Name())
	cmd := exec.Command("clang", "-o", file.Name(), src)
	err = cmd.Run()
	if err != nil {
		return errors.Wrap(err, "execute clang")
	}
	return file.Close()
}
