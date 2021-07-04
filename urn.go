package schema

import (
	"database/sql/driver"
	"fmt"

	"github.com/leodido/go-urn"
)

// URN is an alias to https://github.com/leodido/go-urn.URN
type URN struct {
	value *urn.URN
}

// NewURN creates a new urn
func NewURN(namespace, key string) *URN {
	input := []byte("urn:" + namespace + ":" + key)
	// parse the value
	value, err := urn.NewMachine().Parse(input)
	if err != nil {
		panic(err)
	}

	return &URN{value: value}
}

// URNFromString parses a URN from string
func URNFromString(text string) (*URN, error) {
	value, err := urn.NewMachine().Parse([]byte(text))
	if err != nil {
		return nil, err
	}

	return &URN{value: value}, nil
}

// URNFromBytes parses a URN from bytes
func URNFromBytes(data []byte) (*URN, error) {
	value, err := urn.NewMachine().Parse(data)
	if err != nil {
		return nil, err
	}

	return &URN{value: value}, nil
}

// String returns the urn as string
func (u *URN) String() string {
	// return the resource
	return u.value.String()
}

// Namespace returns the namespace
func (u *URN) Namespace() string {
	return u.value.ID
}

// Key returns the namespace key
func (u *URN) Key() string {
	return u.value.SS
}

// Value implements the driver.Valuer interface.
func (u *URN) Value() (driver.Value, error) {
	return u.String(), nil
}

// Scan implements the sql.Scanner interface.
// A 16-byte slice will be handled by UnmarshalBinary, while
// a longer byte slice or a string will be handled by UnmarshalText.
func (u *URN) Scan(input interface{}) error {
	switch value := input.(type) {
	case *URN:
		// replace the value
		*u = *value
	case string:
		key, err := URNFromString(value)
		if err != nil {
			return err
		}
		// replace the value
		*u = *key
	default:
		return fmt.Errorf("schema: cannot convert %T to URN", input)
	}

	return nil
}
