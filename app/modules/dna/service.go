package dna

import (
	"errors"
	"fmt"
	"regexp"
	"x-men/util/constant"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return service{repository: repository}
}

func (s service) IsMutant(dna []string) (isMutant bool, err error) {
	if len(dna) < constant.MutantSequenceSize {
		err = errors.New(constant.ErrorSizeDNA)
		return
	}

	newDna := DNA{}
	err = newDna.LoadDNA(dna)
	if err != nil {
		return
	}

	count := 0

	for i := 0; i < len(dna); i++ {
		count += s.validateSequencesByPosition(i, newDna)
		if count > 1 {
			isMutant = true
			break
		}
	}

	s.repository.SaveDNA(dna, isMutant)

	return
}

func (s service) validateSequencesByPosition(index int, newDna DNA) (count int) {
	channel := make(chan string, 2)
	if index != newDna.size-1 {
		channel = make(chan string, 6)
		newDna.ReadCrossDiagonalDnaInvertAsync(index, channel)
		newDna.ReadCrossDiagonalDnaAsync(index, channel)
		newDna.ReadDiagonalDnaInvertAsync(index, channel)
		newDna.ReadDiagonalDnaAsync(index, channel)
	}
	newDna.ReadColumnDnaAsync(index, channel)
	newDna.ReadRowDnaAsync(index, channel)

	for i := 0; i < cap(channel); i++ {
		count += len(s.findSequences(constant.MutantSequenceSize, <-channel))
	}
	return
}

func (s service) findSequences(size int, nucleobaseList string) (sequences []string) {
	if len(nucleobaseList) < size {
		return
	}
	regexMutant := regexp.MustCompile(fmt.Sprintf(`[A]{%v}|[T]{%v}|[C]{%v}|[G]{%v}`, size, size, size, size))
	sequences = regexMutant.FindAllString(nucleobaseList, -1)
	return
}
