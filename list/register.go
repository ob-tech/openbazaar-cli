package list

import "github.com/ob-tech/openbazaar-cli/commands"

func init() {
	commands.CMDS["list"] = commands.Command{"Get shop information", nil}
}
