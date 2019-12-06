package inputReader

import (
	"bufio"
	"encoding/csv"
	"os"
)
import "github.com/sirupsen/logrus"

func ReadLines(file string) []string {
	res := make([]string, 0)
	f, err := os.Open(file)

	if err != nil {
		logrus.WithError(err).Error("Could not open file")
		return []string{}
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		res = append(res, s.Text())
	}

	return res
}

func ReadCsv(file string, del rune) [][]string {
	res := make([][]string, 0)
	f, err := os.Open(file)

	if err != nil {
		logrus.WithError(err).Error("Could not open file")
		return [][]string{}
	}

	defer f.Close()

	r := csv.NewReader(f)

	r.Comma = del

	res, err = r.ReadAll()

	if err != nil {
		logrus.WithError(err).Error("Could not read file")
		return [][]string{}
	}

	return res
}
