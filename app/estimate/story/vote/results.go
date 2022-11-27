package vote

import (
	"golang.org/x/exp/slices"
	"strconv"
)

type Results struct {
	Floats []float64 `json:"floats"`
	Count  int       `json:"count"`
	Min    float64   `json:"min"`
	Max    float64   `json:"max"`
	Range  float64   `json:"range"`
	Sum    float64   `json:"sum"`
	Mean   float64   `json:"mean"`
	Median float64   `json:"median"`
	Mode   float64   `json:"mode"`
}

func (v Votes) Results() *Results {
	ret := &Results{Floats: make([]float64, 0, len(v))}
	for _, x := range v {
		fl, err := strconv.ParseFloat(x.Choice, 64)
		if err == nil {
			ret.Floats = append(ret.Floats, fl)
		}
	}
	ret.Count = len(ret.Floats)
	for _, fl := range ret.Floats {
		if fl < ret.Min {
			ret.Min = fl
		}
		if fl > ret.Max {
			ret.Max = fl
		}
		ret.Sum += fl
	}
	ret.Range = ret.Max - ret.Min
	slices.Sort(ret.Floats)
	ret.Mean = ret.Sum / float64(ret.Count)
	return ret
}
