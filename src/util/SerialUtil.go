package util

import (
	"time"
	"strconv"
	"math/rand"
)

func GenerateSerial() string {
	t := time.Now()
	timestamp := strconv.FormatInt(t.UTC().UnixNano(),10)
	randInt := rand.Intn(10000)

	serialnum := timestamp + strconv.Itoa(randInt)

	return serialnum
}
