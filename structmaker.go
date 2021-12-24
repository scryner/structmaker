package main

import (
	"fmt"
	"github.com/scryner/structmaker/convert"
	"io/ioutil"
	"os"
)

func main() {
	// get argument for base struct name
	args := os.Args
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "invalid usage: insufficient argument")
		os.Exit(10)
	}

	baseStructName := args[1]

	// stat stdin and check within shell pipes
	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to stat stdin:", err)
		os.Exit(10)
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Fprintln(os.Stderr, "invalid usage: must be used by shell pipe")
		os.Exit(11)
	}

	// read stdin content
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read from pipe:", err)
		os.Exit(20)
	}

	structContent, err := convert.FromJson(b, baseStructName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to conver from json:", err)
		os.Exit(30)
	}

	// prnt
	fmt.Println(structContent)
}
