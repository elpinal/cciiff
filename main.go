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
	if len(args) < 1 {
		return errors.New("no argument, but want 1 argument")
	}
	if len(args) > 1 {
		return errors.New("too many argument, but want just 1 argument")
	}
	src := args[0]
	return clang(src)
}

const compiler = "clang"

// Executes clang and outputs the file name of the result.
func clang(src string) error {
	file, err := ioutil.TempFile("", "cciiff")
	if err != nil {
		return errors.Wrap(err, "create a temporary file")
	}
	_, err = exec.LookPath(compiler)
	if err != nil {
		return fmt.Errorf("no %s installed", compiler)
	}
	cmd := exec.Command(compiler, "-o", file.Name(), src)
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		os.Remove(file.Name())
		return errors.Wrapf(err, "execute %s", compiler)
	}
	fmt.Printf("file name: %s\n", file.Name())
	return file.Close()
}
