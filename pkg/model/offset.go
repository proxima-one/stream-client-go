package model

import (
	"fmt"
	"github.com/proxima-one/streamdb-client-go/pkg/utils"
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

func (this *Offset) UnmarshalJSON(data []byte) error {
	offset, err := NewOffsetFromString(strings.Trim(string(data), "\""))
	if err != nil {
		return err
	}
	*this = *offset
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
	offset.Height, err = utils.StringToInt64(parts[0])
	if err != nil {
		return nil, err
	}
	offset.OffsetId = parts[1]
	timestamp, err := NewTimestampFromString(parts[2])
	offset.Timestamp = *timestamp
	return &offset, err
}

func (this *Offset) ToString() string {
	if this.Equals(ZeroOffset()) {
		return "0"
	}
	return fmt.Sprintf("%d-%s-%s", this.Height, this.OffsetId, this.Timestamp.ToString())
}

func (this *Offset) Equals(offset *Offset) bool {
	return this.OffsetId == offset.OffsetId
}

func ZeroOffset() *Offset {
	return NewOffset("", 0, ZeroTimestamp())
}

func (this *Offset) CanSucceed(offset Offset) bool {
	return this.Height-1 == offset.Height && this.Timestamp.GreaterThan(offset.Timestamp)
}

func (this *Offset) CanPrecede(offset Offset) bool {
	if this.Equals(ZeroOffset()) {
		return true
	}

	return this.Height+1 == offset.Height && this.Timestamp.LessThan(offset.Timestamp)
}

func (this *Offset) ToDebugString() string {
	return fmt.Sprintf("%v-%v@(%v)", this.Height, this.OffsetId, this.Timestamp.ToDebugString())
}
