package main

import (
	"flag"
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
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: cciiff source_file_name")
	}
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		return errors.New("no argument, but want 1 argument")
	}
	src := args[0]
	return clang(src)
}

func clang(src string) error {
	file, err := ioutil.TempFile("", "cciiff")
	if err != nil {
		return errors.Wrap(err, "create a temporary file")
	}
	cmd := exec.Command("clang", "-o", file.Name(), src)
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		os.Remove(file.Name())
		return errors.Wrap(err, "execute clang")
	}
	fmt.Printf("file name: %s\n", file.Name())
	return file.Close()
}
