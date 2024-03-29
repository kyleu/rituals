package vote

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Results struct {
	Floats     []float64 `json:"floats"`
	Count      int       `json:"count"`
	Min        float64   `json:"min"`
	Max        float64   `json:"max"`
	Range      float64   `json:"range"`
	Sum        float64   `json:"sum"`
	Mean       float64   `json:"mean"`
	Median     float64   `json:"median"`
	Mode       []float64 `json:"mode"`
	ModeString string    `json:"modeString"`
}

func modeString(mode ...float64) string {
	if len(mode) == 0 {
		return "0"
	}
	return strings.Join(lo.Map(mode, func(m float64, _ int) string {
		return fmt.Sprint(m)
	}), ", ")
}

func (v Votes) Results() *Results {
	if len(v) == 0 {
		return &Results{ModeString: "0"}
	}
	ret := &Results{Floats: make([]float64, 0, len(v))}
	lo.ForEach(v, func(x *Vote, _ int) {
		fl, err := strconv.ParseFloat(x.Choice, 64)
		if err == nil {
			if math.IsNaN(fl) || math.IsInf(fl, 0) {
				fl = 0
			}
			ret.Floats = append(ret.Floats, fl)
		}
	})
	ret.Count = len(ret.Floats)
	lo.ForEach(ret.Floats, func(fl float64, idx int) {
		if fl < ret.Min || idx == 0 {
			ret.Min = fl
		}
		if fl > ret.Max {
			ret.Max = fl
		}
		ret.Sum += fl
	})
	ret.Range = ret.Max - ret.Min
	slices.Sort(ret.Floats)
	ret.Mean = ret.Sum / float64(ret.Count)
	if math.IsNaN(ret.Mean) || math.IsInf(ret.Mean, 0) {
		ret.Mean = 0
	}
	ret.Median = median(ret.Floats)
	if math.IsNaN(ret.Median) || math.IsInf(ret.Median, 0) {
		ret.Median = 0
	}
	ret.Mode = mode(ret.Floats)
	ret.ModeString = modeString(ret.Mode...)
	return ret
}

func median(input []float64) float64 {
	c := slices.Clone(input)
	slices.Sort(c)
	l := len(c)
	if l == 0 {
		return math.NaN()
	} else if l%2 == 0 {
		var sum float64
		for _, x := range c[l/2-1 : l/2+1] {
			sum += x
		}
		return sum / 2
	}
	return c[l/2]
}

func mode(input []float64) []float64 {
	l := len(input)
	if l == 1 {
		return input
	} else if l == 0 {
		return nil
	}
	c := slices.Clone(input)
	slices.Sort(c)
	ret := make([]float64, 0, 5)
	cnt, maxCnt := 1, 1
	for i := 1; i < l; i++ {
		switch {
		case c[i] == c[i-1]:
			cnt++
		case cnt == maxCnt && maxCnt != 1:
			ret = append(ret, c[i-1])
			cnt = 1
		case cnt > maxCnt:
			ret = append(ret[:0], c[i-1])
			maxCnt, cnt = cnt, 1
		default:
			cnt = 1
		}
	}
	switch {
	case cnt == maxCnt:
		ret = append(ret, c[l-1])
	case cnt > maxCnt:
		ret = append(ret[:0], c[l-1])
		maxCnt = cnt
	}
	// Since length must be greater than 1,
	// check for slices of distinct values
	if maxCnt == 1 || len(ret)*maxCnt == l && maxCnt != l {
		return []float64{}
	}
	return ret
}
