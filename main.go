package main

import (
	"fmt"
	"github.com/ob-tech/openbazaar-cli/commands"
	"os"

	_ "github.com/ob-tech/openbazaar-cli/list"
)

func init() {
	commands.CMDS["help"] = commands.Command{"Show how " + os.Args[0] + " works", Help}
}

func Help(args []string) error {
	fmt.Println("Usage: ", os.Args[0], " command options\n")
	fmt.Println("Available commands:")
	for t, v := range commands.CMDS {
		fmt.Println("\t", t, "\t", v.Description)
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		Help(nil)
		return
	}
	if exec, ok := commands.CMDS[os.Args[1]]; ok {
		if err := exec.Invoke(os.Args[2:]); err != nil {
			panic(err)
		}
	}
}
