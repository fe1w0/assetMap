package main

import (
	"assetMap/cmd"
	"assetMap/utils"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"runtime"
)



func main() {
	app := &cli.App{
		Name: "assetMap",
		Version: "v1.0.0",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "fe1w0",
			},
		},
		Usage: "Simple assets scanner",
		Action: utils.Scan,
		Flags: []cli.Flag{
			cmd.StringFlag("ip", "", "ip list", "i"),
			cmd.StringFlag("port", "22,23,53,80-139", "port list", "p"),
			cmd.StringFlag("mode", "", "scan mode", "m"),
			cmd.IntFlag("timeout", 4, "timeout", "t"),
			cmd.IntFlag("concurrency", 10, "concurrency", "c"),
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
