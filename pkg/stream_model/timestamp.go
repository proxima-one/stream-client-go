package stream_model

import (
	"encoding/base64"
	"fmt"
	"github.com/proxima-one/streamdb-client-go/internal"
	"strconv"
	"strings"
	"time"
)

type Timestamp struct {
	EpochMs int64    `bson:"epochMs"`
	Parts   []string `bson:"parts"`
}

func ZeroTimestamp() *Timestamp {
	return NewTimestamp(0, nil)
}

func NewTimestamp(epochMs int64, parts []string) *Timestamp {
	return &Timestamp{epochMs, parts}
}

func (this *Timestamp) ToDebugString() string {
	return fmt.Sprintf("%v, %s", this.EpochMs, strings.Join(this.Parts, ","))
}

func (this *Timestamp) ToTime() time.Time {
	return time.Unix(this.EpochMs/1000, this.EpochMs%1000*1e6)
}

func (this *Timestamp) ToString() string {
	if len(this.Parts) == 0 {
		return fmt.Sprint(this.EpochMs)
	}

	partsStr := strings.Join(this.Parts, ",")
	partsBase64 := base64.StdEncoding.EncodeToString([]byte(partsStr))

	return fmt.Sprintf("%d~%s", this.EpochMs, partsBase64)
}

func NewTimestampFromString(s string) (*Timestamp, error) {
	timestampParts := strings.Split(s, "~")
	if len(timestampParts) != 1 && len(timestampParts) != 2 {
		return nil, fmt.Errorf("string represetation of timestamp can't have more than 2 parts or less than 1")
	}

	epochMs, err := internal.StringToInt64(timestampParts[0])
	if err != nil {
		return nil, err
	}

	var parts []string
	if len(timestampParts) == 2 {
		partsStr, err := base64.StdEncoding.DecodeString(timestampParts[1])
		if err != nil {
			return nil, err
		}
		parts = strings.Split(string(partsStr), ",")
	}
	return NewTimestamp(epochMs, parts), nil
}

func (this *Timestamp) Compare(timestamp Timestamp) int {
	if this.EpochMs < timestamp.EpochMs {
		return -1
	}

	if this.EpochMs > timestamp.EpochMs {
		return 1
	}

	minLength := internal.Min(len(this.Parts), len(timestamp.Parts))

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
			continue
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
