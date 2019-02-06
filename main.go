package main

import (
	"log"
	"os"

	"github.com/nvkv/halp/pkg/config/v1"
	"github.com/nvkv/halp/pkg/datasources/googlesheets/v1"
	"github.com/nvkv/halp/pkg/srv/telegram/v1"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "v0"
	app.Name = "halp"
	app.Usage = ""
	app.Description = "Well, you know. I'm here if you need halp"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "config.hcl",
			Usage: "Where to look up config file",
		},
	}

	app.Action = func(c *cli.Context) error {
		configpath := c.String("config")

		cfg, err := config.LoadConfig(configpath)
		if err != nil {
			panic(err)
		}

		halpSheet := googlesheets.Spreadsheet{
			Credentials: cfg.Datasource.GoogleSheets.CredentialsFilePath,
			Tokenfile:   cfg.Datasource.GoogleSheets.TokenFilePath,
			SheetID:     cfg.Datasource.GoogleSheets.SheetID,
			Range:       cfg.Datasource.GoogleSheets.Range,
		}

		bot := telegram.Bot{
			Datasource:       halpSheet,
			Token:            cfg.Telegram.Token,
			WhitelistedChats: cfg.Telegram.WhitelistedChats,
		}
		oerr := bot.Start()
		if oerr != nil {
			panic(oerr)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
