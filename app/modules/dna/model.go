package dna

import (
	"errors"
	"regexp"
	"x-men/util/constant"
)

var regexpDna = `^[ATCG]*$`

type DNA struct {
	dna  [][]rune
	size int
}

func (d *DNA) LoadDNA(dnaMatrix []string) (err error) {
	d.dna = [][]rune{}
	d.size = len(dnaMatrix)
	for _, row := range dnaMatrix {
		err = d.validateDna(row)
		if err != nil {
			return
		}
		d.dna = append(d.dna, []rune(row))
	}
	return
}

func (d *DNA) ReadCrossDiagonalDnaInvert(diagonal int) (nucleobaseList string) {
	maxPosition := d.size - 1
	for i := 0; i < d.size-diagonal; i++ {
		nucleobaseList += string(d.dna[maxPosition-i][i+diagonal])
	}
	return
}

func (d *DNA) ReadCrossDiagonalDna(diagonal int) (nucleobaseList string) {
	for i := 0; i <= diagonal; i++ {
		nucleobaseList += string(d.dna[i][diagonal-i])
	}
	return
}

func (d *DNA) ReadDiagonalDnaInvert(diagonal int) (nucleobaseList string) {
	maxPosition := d.size - 1
	for i := 0; i <= diagonal; i++ {
		nucleobaseList += string(d.dna[maxPosition-i][diagonal-i])
	}
	return
}

func (d *DNA) ReadDiagonalDna(diagonal int) (nucleobaseList string) {
	for i := 0; i < d.size-diagonal; i++ {
		nucleobaseList += string(d.dna[i][diagonal+i])
	}
	return
}

func (d *DNA) ReadColumnDna(column int) (nucleobaseList string) {
	for i := 0; i < d.size; i++ {
		nucleobaseList += string(d.dna[i][column])
	}
	return
}

func (d *DNA) ReadRowDna(row int) (nucleobaseList string) {
	for i := 0; i < d.size; i++ {
		nucleobaseList += string(d.dna[row][i])
	}
	return
}

func (d *DNA) validateDna(dnaRow string) (err error) {
	if len(dnaRow) != d.size {
		err = errors.New(constant.ErrorSquareDNA)
		return
	}

	val, _ := regexp.MatchString(regexpDna, dnaRow)
	if !val {
		err = errors.New(constant.ErrorRegexpDNA)
		return
	}
	return
}

func (d *DNA) ReadCrossDiagonalDnaInvertAsync(diagonal int, c chan string) {
	c <- d.ReadCrossDiagonalDnaInvert(diagonal)
}

func (d *DNA) ReadCrossDiagonalDnaAsync(diagonal int, c chan string) {
	c <- d.ReadCrossDiagonalDna(diagonal)
}

func (d *DNA) ReadDiagonalDnaInvertAsync(diagonal int, c chan string) {
	c <- d.ReadDiagonalDnaInvert(diagonal)
}

func (d *DNA) ReadDiagonalDnaAsync(diagonal int, c chan string) {
	c <- d.ReadDiagonalDna(diagonal)
}

func (d *DNA) ReadColumnDnaAsync(diagonal int, c chan string) {
	c <- d.ReadColumnDna(diagonal)
}

func (d *DNA) ReadRowDnaAsync(diagonal int, c chan string) {
	c <- d.ReadRowDna(diagonal)
}
