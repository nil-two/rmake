package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func usage() {
	os.Stderr.WriteString(`
Usage: rmake [OPTION]... [MAKE-ARGS]...
Recursively find parent directory's Makefile, and execute make.

Options:
	--help       show this help message
	--version    print the version
`[1:])
}

func version() {
	os.Stderr.WriteString(`
v0.1.0
`[1:])
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func do(args []string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	for d := wd; d != filepath.Dir(d); d = filepath.Dir(d) {
		if exists(filepath.Join(d, "Makefile")) {
			if err = os.Chdir(d); err != nil {
				return err
			}

			c := exec.Command("make", args...)
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			return c.Run()
		}
	}
	return nil
}

func _main() error {
	isHelp := flag.Bool("help", false, "")
	isVersion := flag.Bool("version", false, "")
	flag.Usage = usage
	flag.Parse()
	switch {
	case *isHelp:
		usage()
		return nil
	case *isVersion:
		version()
		return nil
	}

	if flag.NArg() < 1 {
		return do(nil)
	}
	return do(flag.Args()[1:])
}

func main() {
	if err := _main(); err != nil {
		fmt.Fprintln(os.Stderr, "rmake:", "error:", err)
		os.Exit(1)
	}
}
