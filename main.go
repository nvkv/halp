package main

import (
	"time"

	"github.com/nvkv/halp/pkg/config/v1"
	"github.com/nvkv/halp/pkg/datasources/googlesheets/v1"
	"github.com/nvkv/halp/pkg/outputs/console/v1"
	"github.com/nvkv/halp/pkg/schedule/v1"
)

func main() {
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		panic(err)
	}

	halpSheet := googlesheets.Spreadsheet{
		Credentials: cfg.Datasource.GoogleSheets.CredentialsFilePath,
		Tokenfile:   cfg.Datasource.GoogleSheets.TokenFilePath,
		SheetID:     cfg.Datasource.GoogleSheets.SheetID,
		Range:       cfg.Datasource.GoogleSheets.Range,
	}

	schedule, err := schedule.ScheduleWeek(time.Now(), halpSheet)
	if err != nil {
		panic(err)
	}

	outp := console.ConsoleOutput{}
	oerr := outp.Send(schedule)
	if oerr != nil {
		panic(oerr)
	}
}
