package model

import "fmt"

type Offset struct {
	OffsetId  string    `bson:"offsetId"`
	Height    int       `bson:"height"`
	Timestamp Timestamp `bson:"timestamp"`
}

func NewOffset(id string, height int, timestamp Timestamp) *Offset {
	if len(id) == 0 && height > 0 {
		panic("Offset costructor: invalid args.")
	}

	return &Offset{id, height, timestamp}
}

func (this *Offset) Equals(offset Offset) bool {
	return this.OffsetId == offset.OffsetId
}

func ZeroOffset() *Offset {
	return NewOffset("", 0, *ZeroTimestamp())
}

func (this *Offset) CanSucceed(offset Offset) bool {
	return this.Height-1 == offset.Height && this.Timestamp.GreaterThan(offset.Timestamp)
}

func (this *Offset) CanPrecede(offset Offset) bool {
	if this.Equals(*ZeroOffset()) {
		return true
	}

	return this.Height+1 == offset.Height && this.Timestamp.LessThan(offset.Timestamp)
}

func (this *Offset) Dump() string {
	return fmt.Sprintf("%v-%v@(%v)", this.Height, this.OffsetId, this.Timestamp.Dump())
}
