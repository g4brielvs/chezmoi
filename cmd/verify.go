package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/twpayne/chezmoi/lib/chezmoi"
	vfs "github.com/twpayne/go-vfs"
)

var verifyCommand = &cobra.Command{
	Use:   "verify",
	Short: "Exit with success if the actual state matches the target state, fail otherwise",
	RunE:  makeRunE(config.runVerifyCommand),
}

func init() {
	rootCommand.AddCommand(verifyCommand)
}

func (c *Config) runVerifyCommand(fs vfs.FS, command *cobra.Command, args []string) error {
	actuator := chezmoi.NewAnyActuator(chezmoi.NullActuator)
	if err := c.applyArgs(fs, args, actuator); err != nil {
		return err
	}
	if actuator.Actuated() {
		os.Exit(1)
	}
	return nil
}
