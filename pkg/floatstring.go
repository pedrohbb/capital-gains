package pkg

import "strconv"

type Float64 float64

func (fs Float64) MarshalJSON() ([]byte, error) {
	vs := strconv.FormatFloat(float64(fs), 'f', 2, 64)
	return []byte(vs), nil
}
