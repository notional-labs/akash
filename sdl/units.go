package sdl

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	errNegativeValue = fmt.Errorf("invalid: negative value not allowed")
)

const (
	unitk  = 1000
	unitKi = 1024

	unitM  = unitk * unitk
	unitMi = unitKi * unitKi

	unitG  = unitM * unitk
	unitGi = unitMi * unitKi

	unitT  = unitG * unitk
	unitTi = unitGi * unitKi

	unitP  = unitT * unitk
	unitPi = unitTi * unitKi

	unitE  = unitP * unitk
	unitEi = unitPi * unitKi
)

var suffixes = []struct {
	symbol string
	unit   uint64
}{
	{"k", unitk},
	{"Ki", unitKi},

	{"M", unitM},
	{"Mi", unitMi},

	{"G", unitG},
	{"Gi", unitGi},

	{"T", unitT},
	{"Ti", unitTi},

	{"P", unitP},
	{"Pi", unitPi},

	{"E", unitE},
	{"Ei", unitEi},
}

// CPU shares.  One CPUQuantity = 1/1000 of a CPU
type cpuQuantity uint32

func (u *cpuQuantity) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var sval string
	if err := unmarshal(&sval); err != nil {
		return err
	}

	if strings.HasSuffix(sval, "m") {
		sval = strings.TrimSuffix(sval, "m")
		val, err := strconv.ParseUint(sval, 10, 32)
		if err != nil {
			return err
		}
		if val < 0 {
			return errNegativeValue
		}
		*u = cpuQuantity(val)
		return nil
	}

	val, err := strconv.ParseFloat(sval, 64)
	if err != nil {
		return err
	}

	val *= 1000

	if val < 0 {
		return errNegativeValue
	}

	*u = cpuQuantity(val)

	return nil
}

// Memory,Disk size in bytes.
type byteQuantity uint64

func (u *byteQuantity) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var sval string
	if err := unmarshal(&sval); err != nil {
		return err
	}
	val, err := parseWithSuffix(sval)
	if err != nil {
		return err
	}
	*u = byteQuantity(val)
	return nil
}

func parseWithSuffix(sval string) (uint64, error) {
	for _, suffix := range suffixes {
		if !strings.HasSuffix(sval, suffix.symbol) {
			continue
		}

		sval := strings.TrimSuffix(sval, suffix.symbol)

		val, err := strconv.ParseFloat(sval, 64)
		if err != nil {
			return 0, err
		}

		val *= float64(suffix.unit)

		if val < 0 {
			return 0, errNegativeValue
		}

		return uint64(val), nil
	}

	val, err := strconv.ParseFloat(sval, 64)
	if err != nil {
		return 0, err
	}

	if val < 0 {
		return 0, errNegativeValue
	}

	return uint64(val), nil
}
