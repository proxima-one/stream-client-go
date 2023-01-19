package proximaclient

import "errors"

var offsetNotFoundError = errors.New("offset not found")

func OffsetNotFoundError() error {
	return offsetNotFoundError
}
