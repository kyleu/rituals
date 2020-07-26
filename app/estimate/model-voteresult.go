package estimate

import (
	"fmt"
	"math"
	"strconv"
)

type VoteResult struct {
	Votes     Votes `json:"votes"`
	FinalVote string
	Count     int
	Min       string
	Max       string
	Sum       string
	Mean      string
	Median    string
	Mode      string
}

func CalculateVoteResult(v Votes) VoteResult {
	choices := make([]float64, 0)
	for _, vote := range v {
		f, err := strconv.ParseFloat(vote.Choice, 64)
		if err == nil {
			choices = append(choices, f)
		}
	}

	sum := float64(0)
	for _, f := range choices {
		sum += f
	}

	var min, max float64
	for i, e := range choices {
		if i == 0 || e < min {
			min = e
		}
		if i == 0 || e > max {
			max = e
		}
	}

	var final, mean, median float64
	if len(choices) > 0 {
		final = sum / float64(len(choices))
		mean = sum / float64(len(choices))
		median = choices[int(math.Floor(float64(len(choices))/2.0))]
	}

	return VoteResult{
		Votes:     v,
		FinalVote: trim(final),
		Count:     len(choices),
		Min:       trim(min),
		Max:       trim(max),
		Sum:       trim(sum),
		Mean:      trim(mean),
		Median:    trim(median),
		Mode:      modeCalc(v),
	}
}

func trim(v float64) string {
	fv := fmt.Sprint(v)
	maxChars := 4
	if len(fv) >= maxChars {
		fv = fv[0:maxChars]
	}
	return fv
}

func modeCalc(x Votes) string {
	var maxValue string
	var maxCount int
	for _, e := range x {
		count := 0
		for _, e2 := range x {
			if e.Choice == e2.Choice {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
			maxValue = e.Choice
		}
	}
	return maxValue
}
