package cmd

import (
	"github.com/urfave/cli/v2"
)

func StringFlag(name, value, usage, aliases string) *cli.StringFlag {
	return &cli.StringFlag{
		Name: name,
		Value: value,
		Usage: usage,
		Aliases: []string{aliases},
	}
}

func BoolFlag(name, usage, aliases string) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:  name,
		Usage: usage,
		Aliases: []string{aliases},
	}
}

func IntFlag(name string, value int, usage, aliases string) *cli.IntFlag {
	return &cli.IntFlag{
		Name:  name,
		Value: value,
		Usage: usage,
		Aliases: []string{aliases},
	}
}
