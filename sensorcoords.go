package main

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"time"
)

//SensorCoords данные датчика.
type SensorCoords struct {
	Time time.Time
	X    float64
	Y    float64
	Z    float64
}

//FromString из строки, приходящей с устройства.
func (sc *SensorCoords) FromString(t time.Time, str string) (err error) {
	if nil == sc {
		return
	}
	sc.Time = t
	str = strings.Replace(str, "X", "", -1)
	str = strings.Replace(str, "Y", "", -1)
	str = strings.Replace(str, "Z", "", -1)
	str = strings.Replace(str, ";", "", -1)
	v := strings.Split(str, ",")
	if len(v) < 3 {
		err = errors.New("split failed")
		return
	}
	sc.X, err = strconv.ParseFloat(v[0], 64)
	if err != nil {
		return
	}
	sc.Y, err = strconv.ParseFloat(v[1], 64)
	if err != nil {
		return
	}
	sc.Z, err = strconv.ParseFloat(v[2], 64)

	rng := 2.
	sc.X = (sc.X / math.MaxInt16) * rng * 9.80664999999998
	sc.Y = (sc.Y / math.MaxInt16) * rng * 9.80664999999998
	sc.Z = (sc.Z / math.MaxInt16) * rng * 9.80664999999998

	return
}
