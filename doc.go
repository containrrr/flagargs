/*
Package flagargs is a helper for stdlib flag (and github.com/spf13/pflag) to handle known positional arguments

Example:
	import (
		"containrrr.dev/flagargs"
		"flag"
		"fmt"
	)

	var parser flagargs.Parser
	var name string

	func init() {
		flag.StringVar(&name, "name", "", "the name of the thing")
		parser = flagargs.NewParser("name")
	}

	func main() {
		flag.Parse()
		extra, err := parser.ParseAndUpdateFlags(flag.CommandLine, flag.CommandLine.Args())
		if err != nil {
			_ = fmt.Errorf("error setting known arg flag: %v", err)
			return
		}
		fmt.Printf("name was set to %q\n", name)
		fmt.Printf("got %v extra argument(s): %v\n", len(extra), extra)
	}

*/
package flagargs
