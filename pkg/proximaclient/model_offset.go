package proximaclient

import (
	"fmt"
	"github.com/proxima-one/stream-client-go/v2/internal"
	"strings"
)

type Offset struct {
	OffsetId  string    `bson:"offsetId"`
	Height    int64     `bson:"height"`
	Timestamp Timestamp `bson:"timestamp"`
}

func NewOffset(id string, height int64, timestamp *Timestamp) *Offset {
	if len(id) == 0 && height > 0 {
		panic("Offset constructor: invalid args.")
	}
	return &Offset{id, height, *timestamp}
}

func (offset *Offset) UnmarshalJSON(data []byte) error {
	parsedOffset, err := NewOffsetFromString(strings.Trim(string(data), "\""))
	if err != nil {
		return err
	}
	*offset = *parsedOffset
	return nil
}

func NewOffsetFromString(str string) (*Offset, error) {
	var err error
	var offset Offset

	if str == "0" {
		return ZeroOffset(), nil
	}

	parts := strings.Split(str, "-")
	if len(parts) != 3 {
		return nil, fmt.Errorf("offset parts size must be 3")
	}
	offset.Height, err = internal.StringToInt64(parts[0])
	if err != nil {
		return nil, err
	}
	offset.OffsetId = parts[1]
	timestamp, err := NewTimestampFromString(parts[2])
	offset.Timestamp = *timestamp
	return &offset, err
}

func (offset *Offset) String() string {
	if offset.Equals(ZeroOffset()) {
		return "0"
	}
	return fmt.Sprintf("%d-%s-%s", offset.Height, offset.OffsetId, offset.Timestamp.String())
}

func (offset *Offset) Equals(another *Offset) bool {
	return offset.OffsetId == another.OffsetId
}

func ZeroOffset() *Offset {
	return NewOffset("", 0, ZeroTimestamp())
}

func (offset *Offset) CanBeAfter(another Offset) bool {
	return offset.Height > another.Height && offset.Timestamp.GreaterThan(another.Timestamp)
}

func (offset *Offset) CanSucceed(another Offset) bool {
	return offset.Height-1 == another.Height && offset.Timestamp.GreaterThan(another.Timestamp)
}

func (offset *Offset) CanPrecede(another Offset) bool {
	if offset.Equals(ZeroOffset()) {
		return true
	}

	return offset.Height+1 == another.Height && offset.Timestamp.LessThan(another.Timestamp)
}

func (offset *Offset) DebugString() string {
	return fmt.Sprintf("%v-%v@(%v)", offset.Height, offset.OffsetId, offset.Timestamp.DebugString())
}
