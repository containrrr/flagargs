package cobra_cmd_test

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

func ExampleKnownCommandArgs() {
	// Override arguments for example, normally this uses os.Args[1:]
	rootCmd.SetArgs([]string{"fish", "fiddle", "sprockets", "shoestrings"})
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
	// Output: The animal "fiddle" is named "fish". Here is the extra args: [sprockets shoestrings]
}
