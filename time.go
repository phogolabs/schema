package schema

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"fmt"
	"time"

	"github.com/volatiletech/null"
)

var (
	// DateTimeFormat represents the default format of the time
	DateTimeFormat = "20060102"
)

// A Date represents an instant in time with nanosecond precision.
type Date struct {
	time.Time
}

// MarshalText implements the encoding.TextMarshaler interface.
// The time is formatted in RFC 3339 format, with sub-second precision added if present.
func (t Date) MarshalText() ([]byte, error) {
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalText: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateTimeFormat))
	return t.AppendFormat(b, DateTimeFormat), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The time is expected to be in RFC 3339 format.
func (t *Date) UnmarshalText(data []byte) error {
	// Fractional seconds are handled implicitly by Parse.
	parsed, err := time.Parse(DateTimeFormat, string(data))
	if err != nil {
		return err
	}

	*t = Date{parsed}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in RFC 3339 format, with sub-second precision added if present.
func (t Date) MarshalJSON() ([]byte, error) {
	if y := t.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateTimeFormat)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, DateTimeFormat)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (t *Date) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == string(null.NullBytes) {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	parsed, err := time.Parse(`"`+DateTimeFormat+`"`, string(data))
	if err != nil {
		return err
	}

	*t = Date{parsed}
	return nil
}

// A NullDate represents an instant in time with nanosecond precision.
type NullDate struct {
	Date  Date
	Valid bool
}

// NewNullDate creates a new Time.
func NewNullDate(t time.Time, valid bool) NullDate {
	return NullDate{
		Date:  Date{t},
		Valid: valid,
	}
}

// NullDateFrom creates a new Time that will always be valid.
func NullDateFrom(t time.Time) NullDate {
	return NewNullDate(t, true)
}

// NullDateFromPtr creates a new Time that will be null if t is nil.
func NullDateFromPtr(t *time.Time) NullDate {
	if t == nil {
		return NewNullDate(time.Time{}, false)
	}
	return NewNullDate(*t, true)
}

// MarshalJSON implements json.Marshaler.
func (t NullDate) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return null.NullBytes, nil
	}
	return t.Date.MarshalJSON()
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *NullDate) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, null.NullBytes) {
		t.Valid = false
		t.Date = Date{}
		return nil
	}

	if err := t.Date.UnmarshalJSON(data); err != nil {
		return err
	}

	t.Valid = true
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (t NullDate) MarshalText() ([]byte, error) {
	if !t.Valid {
		return null.NullBytes, nil
	}
	return t.Date.MarshalText()
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (t *NullDate) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		t.Valid = false
		return nil
	}
	if err := t.Date.UnmarshalText(text); err != nil {
		return err
	}
	t.Valid = true
	return nil
}

// SetValid changes this Time's value and sets it to be non-null.
func (t *NullDate) SetValid(v time.Time) {
	t.Date = Date{v}
	t.Valid = true
}

// Ptr returns a pointer to this Time's value, or a nil pointer if this Time is null.
func (t NullDate) Ptr() *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Date.Time
}

// Scan implements the Scanner interface.
func (t *NullDate) Scan(value interface{}) error {
	var err error
	switch x := value.(type) {
	case time.Time:
		t.Date = Date{x}
	case nil:
		t.Valid = false
		return nil
	default:
		err = fmt.Errorf("null: cannot scan type %T into null.Time: %v", value, value)
	}
	t.Valid = err == nil
	return err
}

// Value implements the driver Valuer interface.
func (t NullDate) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Date, nil
}
