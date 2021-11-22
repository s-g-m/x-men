package adapter

import (
	"x-men/app/modules/dna"
	"x-men/app/modules/statistic"
)

type HttpAdapter interface {
	statistic.HttpAdapter
	dna.HttpAdapter
}
