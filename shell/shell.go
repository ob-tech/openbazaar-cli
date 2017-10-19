package shell

import (
	"fmt"
	"strings"
	"github.com/chzyer/readline"
	"github.com/ob-tech/openbazaar-cli/commands"
)

func init() {
	commands.CMDS["shell"] = commands.Command{"Run commands on a shell.", Shell}
}

func Shell(_ []string) error {
	rl, err := readline.New("> ")
	if err != nil {
		return err
	}
	defer rl.Close()

prompt:
	for {
		line, err := rl.Readline()
		if err != nil {
			return err
		}
		args := strings.Split(line, " ")
		if args[0] == "exit" {
			return nil
		}
		for k, v := range commands.CMDS {
			if args[0] == k {
				v.Invoke(args[1:])
				fmt.Println()
				continue prompt
			}
		}
	}
}
