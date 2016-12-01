package telemetry

import (
	"encoding/binary"
	"errors"
	"math"
)

type TMPC struct {
	Temp float32
}

func (temp *TMPC) Parse(bytes []byte, scale *SCAL) error {
	if 4 != len(bytes) {
		return errors.New("Invalid length TMPC packet")
	}

	bits := binary.BigEndian.Uint32(bytes[0:4])

	temp.Temp = math.Float32frombits(bits)

	return nil
}
