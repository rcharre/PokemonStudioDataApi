package cli

import (
	"flag"
	"os"
)

type ExecFunc func() error

type Command struct {
	flagSet *flag.FlagSet
	exec    ExecFunc
}

func NewCommand(flagSet *flag.FlagSet, exec ExecFunc) *Command {
	return &Command{
		flagSet,
		exec,
	}
}

func (c Command) Usage() {
	c.flagSet.Usage()
}

func (c Command) Execute() (err error) {
	if err = c.flagSet.Parse(os.Args[1:]); err != nil {
		return err
	}
	return c.exec()
}
