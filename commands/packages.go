package commands

type Command struct {
	Description string
	Invoke      func([]string) error
}

var CMDS = make(map[string]Command)
