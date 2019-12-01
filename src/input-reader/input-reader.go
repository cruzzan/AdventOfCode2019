package input_reader

import (
	"bufio"
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
