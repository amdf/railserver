package main

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

//SensorCoords данные датчика.
type SensorCoords struct {
	Time time.Time
	X    int
	Y    int
	Z    int
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
	sc.X, err = strconv.Atoi(v[0])
	if err != nil {
		return
	}
	sc.Y, err = strconv.Atoi(v[1])
	if err != nil {
		return
	}
	sc.Z, err = strconv.Atoi(v[2])
	return
}
