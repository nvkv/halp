package output

import "github.com/nvkv/halp/pkg/types/data/v1"

type Output interface {
	Send(schedule []data.Day) error
}
