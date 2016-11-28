package v2

import (
	"os"

	"code.cloudfoundry.org/cli/cf/cmd"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/flags"
)

type QuotaCommand struct {
	RequiredArgs    flags.Quota `positional-args:"yes"`
	usage           interface{} `usage:"CF_NAME quota QUOTA"`
	relatedCommands interface{} `related_commands:"org, quotas"`
}

func (_ QuotaCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (_ QuotaCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}