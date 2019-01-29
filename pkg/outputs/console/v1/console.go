package console

import (
	"fmt"

	"github.com/nvkv/halp/pkg/types/data/v1"
	"github.com/nvkv/halp/pkg/types/output/v1"
)

type ConsoleOutput struct{}

func (o ConsoleOutput) Send(schedule []data.Day) *output.OutputError {
	for _, m := range schedule {
		fmt.Println(m)
	}
	return nil
}
