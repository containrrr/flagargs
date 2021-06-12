
# Flagargs

A helper for flag-derived command line parsers for mapping non-flag arguments to a set of ordered flags and 
separating all unknown flags (as well as anyhthing after `--`)  to a separate slice.

The `containrrr.dev/flagargs` is dependency-free, but `containrrr.dev/flagargs/cobra_cmd` depends 
on `github.com/spf13/cobra`.

## Install
```shell
go get containrrr.dev/flagargs/cobra_cmd
```
or for the core library (for usage with pflag or flag):
```shell
go get containrrr.dev/flagargs
```

## Usage

### Cobra
```go
package main

import (
	"containrrr.dev/flagargs/cobra_cmd"
	"fmt"
	"github.com/spf13/cobra"
)

var extraArgs []string
var rootCmd = &cobra.Command{
	Short: "hello",
	Args:  cobra_cmd.KnownCommandArgs(&extraArgs, "name", "animal"),
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		animal, _ := cmd.Flags().GetString("animal")
		fmt.Printf("The animal %q is named %q. Here is the extra args: %v", animal, name, extraArgs)
	},
}

func init() {
	rootCmd.PersistentFlags().String("name", "NAME", "the name")
	rootCmd.PersistentFlags().String("animal", "ANIMAL", "the animal")
}

func main() {
	rootCmd.SetArgs([]string{"fish", "fiddle", "sprockets", "shoestrings"})
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
	// Output: The animal "fiddle" is named "fish". Here is the extra args: [sprockets shoestrings]
}
```

### pflag
```go
package main

import (
	"containrrr.dev/flagargs"
	"fmt"
	"github.com/spf13/pflag"
)

var (
	flags *pflag.FlagSet
	parser *flagargs.Parser
)

func init() {
	flags = pflag.NewFlagSet("example", pflag.ExitOnError)
	flags.String("name", "", "the name of the thing")
	parser = flagargs.NewParser("name")
}

func main() {
	extra, err := parser.ParseAndUpdateFlags(flags, flags.Args())
	if err != nil {
		_ = fmt.Errorf("error setting known arg flag: %v", err)
		return
	}
	name, _ := flags.GetString("name")
	fmt.Printf("name was set to %q, ", name)
	fmt.Printf("got %v extra argument(s): %v\n", len(extra), extra)
}
```

### stdlib flag
```go
package main

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

```