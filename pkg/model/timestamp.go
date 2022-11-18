package model

import (
	"fmt"
	"strconv"
	"strings"
)

type Timestamp struct {
	EpochMs int      `bson:"epochMs"`
	Parts   []string `bson:"parts"`
}

func ZeroTimestamp() *Timestamp {
	return NewTimestamp(0, nil)
}

func NewTimestamp(epochMs int, parts []string) *Timestamp {
	p := parts
	if parts == nil {
		p = []string{}
	}
	return &Timestamp{epochMs, p}
}

func (this *Timestamp) Dump() string {
	return fmt.Sprintf("%v, %s", this.EpochMs, strings.Join(this.Parts, ","))
}

func (this *Timestamp) Compare(timestamp Timestamp) int {
	if this.EpochMs < timestamp.EpochMs {
		return -1
	}

	if this.EpochMs > timestamp.EpochMs {
		return 1
	}

	minLength := Min(len(this.Parts), len(timestamp.Parts))

	for i := 0; i < minLength; i++ {
		a, aerr := strconv.Atoi(this.Parts[i])
		b, berr := strconv.Atoi(timestamp.Parts[i])

		if aerr == nil && berr == nil {
			if a > b {
				return 1
			}

			if a < b {
				return -1
			}
			return 0
		}

		if this.Parts[i] > timestamp.Parts[i] {
			return 1
		}

		if this.Parts[i] < timestamp.Parts[i] {
			return -1
		}
	}

	// shorter length means less
	return len(this.Parts) - len(timestamp.Parts)
}

func (this *Timestamp) Equals(timestamp Timestamp) bool {
	return this.Compare(timestamp) == 0
}

func (this *Timestamp) GreaterThan(timestamp Timestamp) bool {
	return this.Compare(timestamp) > 0
}

func (this *Timestamp) LessThan(timestamp Timestamp) bool {
	return this.Compare(timestamp) < 0
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
