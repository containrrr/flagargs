package cobra_cmd

import (
	"containrrr.dev/flagargs"
	"errors"
	"github.com/spf13/cobra"
)

func KnownCommandArgs(extraArgs *[]string, knownArgs ...string) cobra.PositionalArgs {
	parser := flagargs.NewParser(knownArgs...)
	return func(cmd *cobra.Command, args []string) error {
		flags := cmd.Flags()
		extra, err := parser.ParseAndUpdateFlags(flags, args)
		if err != nil {
			return err
		}
		if extraArgs == nil {
			return errors.New("extraArgs is nil")
		}

		*extraArgs = extra
		return nil
	}
}
