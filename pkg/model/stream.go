package model

type Stream struct {
	Id    string `bson:"_id"`
	Start Offset `bson:"start"`
	End   Offset `bson:"end"`
}
