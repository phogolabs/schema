package schema

import (
	"fmt"

	"github.com/rs/xid"
)

// ID represents a unique request id
type ID = xid.ID

// IDFromBytes convert the byte array representation of `ID` back to `ID`
func IDFromBytes(data []byte) (ID, error) {
	value, err := xid.FromBytes(data)
	if err != nil {
		err = fmt.Errorf("%w %q", err, string(data))
	}

	return value, err
}

// IDFromString reads an ID from its string representation
func IDFromString(text string) (ID, error) {
	value, err := xid.FromString(text)
	if err != nil {
		err = fmt.Errorf("%w %q", err, text)
	}

	return value, err
}

// NewID creates a ID
func NewID() ID {
	return xid.New()
}
