package schema

import "github.com/rs/xid"

// ID represents a unique request id
type ID = xid.ID

var (
	// IDFromBytes convert the byte array representation of `ID` back to `ID`
	IDFromBytes = xid.FromBytes

	// IDFromString reads an ID from its string representation
	IDFromString = xid.FromString
)

// NewID creates a ID
func NewID() ID {
	return xid.New()
}
