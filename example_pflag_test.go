package flagargs_test

import (
	"containrrr.dev/flagargs"
	"fmt"
	"github.com/spf13/pflag"
)

func ExampleParser_ParseAndUpdateFlags_pFlag() {
	flags := pflag.NewFlagSet("example", pflag.ExitOnError)
	flags.String("name", "", "the name of the thing")
	parser := flagargs.NewParser("name")

	// Normally you would get the args from os.Args:
	//	 _ = flags.Parse(os.Args)
	// In this example we'll just use a hard coded slice:
	_ = flags.Parse([]string{"arg1", "arg2", "arg3"})

	extra, err := parser.ParseAndUpdateFlags(flags, flags.Args())
	if err != nil {
		_ = fmt.Errorf("error setting known arg flag: %v", err)
		return
	}
	name, _ := flags.GetString("name")
	fmt.Printf("name was set to %q, ", name)
	fmt.Printf("got %v extra argument(s): %v\n", len(extra), extra)

	// Output: name was set to "arg1", got 2 extra argument(s): [arg2 arg3]
}
