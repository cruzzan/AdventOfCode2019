package InputConverter

import (
	"github.com/sirupsen/logrus"
	"strconv"
)

func StringToInt64(i []string) []int64 {
	res := make([]int64, 0)

	for _, val := range i{
		conv, err := strconv.Atoi(val)

		if err != nil {
			logrus.WithError(err).Error("Could not convert string to int")
		}

		res = append(res, int64(conv))
	}

	return res
}

func StringToInt(i []string) []int {
	res := make([]int, 0)

	for _, val := range i{
		conv, err := strconv.Atoi(val)

		if err != nil {
			logrus.WithError(err).Error("Could not convert string to int")
		}

		res = append(res, int(conv))
	}

	return res
}
