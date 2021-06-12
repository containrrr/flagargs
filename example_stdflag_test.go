package flagargs_test

import (
	"containrrr.dev/flagargs"
	"flag"
	"fmt"
)

var parser *flagargs.Parser
var name string

func init() {
	flag.StringVar(&name, "name", "", "the name of the thing")
	parser = flagargs.NewParser("name")
}

func ExampleParser_ParseAndUpdateFlags_stdFlag() {
	flag.Parse()
	extra, err := parser.ParseAndUpdateFlags(flag.CommandLine, flag.CommandLine.Args())
	if err != nil {
		_ = fmt.Errorf("error setting known arg flag: %v", err)
		return
	}
	fmt.Printf("name was set to %q\n", name)
	fmt.Printf("got %v extra argument(s): %v\n", len(extra), extra)
}
