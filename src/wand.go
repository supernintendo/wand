package main

import (
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"log"
	"mime"
	"os"
	"path/filepath"
)

func main() {
	app := cli.NewApp()
	app.Name = "wand"
	app.Version = "1.0.0"
	app.Usage = "serve a single file over http"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "port, p",
			Usage: "set server port",
			Value: "11002",
		},
		cli.BoolFlag{
			Name:  "script, s",
			Usage: "execute file on request",
		},
	}
	app.Action = func(c *cli.Context) error {
		config := called(c)
		WandWWW.Host(config)

		return nil
	}
	app.Run(os.Args)
}

func called(c *cli.Context) WandConfig {
	// Called with a filename, check to see if that
	// file exists. If it does not, exit.
	config := WandConfig{}

	if c.NArg() > 0 {
		filenameArg := c.Args().Get(0)

		if _, err := os.Stat(filenameArg); err == nil {
			config = WandConfig{
				ContentType: detectContentType(filenameArg),
				Filename:    filenameArg,
				Filepath:    detectFilePath(filenameArg),
				Port:        c.String("port"),
				Script:      c.Bool("script"),
			}
		} else {
			color.Red("Wand: The file '%s' does not exist. Exiting...", filenameArg)
			os.Exit(1)
		}
	} else {
		color.Red("Wand: No file name provided. Exiting...")
		os.Exit(1)
	}
	return config
}

func detectContentType(filename string) string {
	return mime.TypeByExtension(filepath.Ext(filename))
}

func detectFilePath(filename string) string {
	path, err := filepath.Abs(filename)

	if err != nil {
		log.Fatal(err)
	}
	return path
}
