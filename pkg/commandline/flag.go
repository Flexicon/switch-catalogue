package commandline

import "github.com/urfave/cli"

func (cmd *Cmd) RegisterFlags() {
	cmd.app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "dry-run, d",
			Usage: "Run a command without persisting changes",
		},
	}
}
