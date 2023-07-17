package proximaclient

import (
	"encoding/base64"
	"fmt"
	"github.com/proxima-one/stream-client-go/v2/internal"
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

func (timestamp *Timestamp) DebugString() string {
	return fmt.Sprintf("%v, %s", timestamp.EpochMs, strings.Join(timestamp.Parts, ","))
}

func (timestamp *Timestamp) Time() time.Time {
	return time.Unix(timestamp.EpochMs/1000, timestamp.EpochMs%1000*1e6)
}

func (timestamp *Timestamp) String() string {
	if len(timestamp.Parts) == 0 {
		return fmt.Sprint(timestamp.EpochMs)
	}

	partsStr := strings.Join(timestamp.Parts, ",")
	partsBase64 := base64.StdEncoding.EncodeToString([]byte(partsStr))

	return fmt.Sprintf("%d~%s", timestamp.EpochMs, partsBase64)
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

func (timestamp *Timestamp) Compare(another Timestamp) int {
	switch {
	case timestamp.EpochMs < another.EpochMs:
		return -1
	case timestamp.EpochMs > another.EpochMs:
		return 1
	}

	minLength := internal.Min(len(timestamp.Parts), len(another.Parts))
	for i := 0; i < minLength; i++ {
		if a, err := strconv.Atoi(timestamp.Parts[i]); err == nil {
			if b, err := strconv.Atoi(another.Parts[i]); err == nil {
				switch {
				case a > b:
					return 1
				case a < b:
					return -1
				}
				continue
			}
		}

		switch {
		case timestamp.Parts[i] > another.Parts[i]:
			return 1
		case timestamp.Parts[i] < another.Parts[i]:
			return -1
		}
	}

	// shorter length means less
	return len(timestamp.Parts) - len(another.Parts)
}

func (timestamp *Timestamp) Equals(another Timestamp) bool {
	return timestamp.Compare(another) == 0
}

func (timestamp *Timestamp) GreaterThan(another Timestamp) bool {
	return timestamp.Compare(another) > 0
}

func (timestamp *Timestamp) LessThan(another Timestamp) bool {
	return timestamp.Compare(another) < 0
}
