package schema

import (
	"github.com/volatiletech/null"
)

// NullBool is a nullable bool.
type NullBool = null.Bool

var (
	// NewNullBool creates a new Bool
	NewNullBool = null.NewBool

	// NullBoolFrom creates a new Bool that will always be valid.
	NullBoolFrom = null.BoolFrom

	// NullBoolFromPtr creates a new Bool that will be null if f is nil.
	NullBoolFromPtr = null.BoolFromPtr
)

// NullByte is an nullable int.
type NullByte = null.Byte

var (
	// NewNullByte creates a new Byte
	NewNullByte = null.NewByte

	// NullByteFrom creates a new Byte that will always be valid.
	NullByteFrom = null.ByteFrom

	// NullByteFromPtr creates a new Byte that be null if i is nil.
	NullByteFromPtr = null.ByteFromPtr
)

// // NullBytes is a global byte slice of JSON null
// var NullBytes = null.NullBytes

// NullBytes is a nullable []byte.
type NullBytes = null.Bytes

var (
	// NewNullBytes creates a new Bytes
	NewNullBytes = null.NewBytes

	// NullBytesFrom creates a new Bytes that will be invalid if nil.
	NullBytesFrom = null.BytesFrom

	// NullBytesFromPtr creates a new Bytes that will be invalid if nil.
	NullBytesFromPtr = null.BytesFrom
)

// NullFloat32 is a nullable float32.
type NullFloat32 = null.Float32

var (
	// NewNullFloat32 creates a new Float32
	NewNullFloat32 = null.NewFloat32

	// NullFloat32From creates a new Float32 that will always be valid.
	NullFloat32From = null.Float32From

	// NullFloat32FromPtr creates a new Float32 that be null if f is nil.
	NullFloat32FromPtr = null.Float32FromPtr
)

// NullFloat64 is a nullable float64.
type NullFloat64 = null.Float64

var (
	// NewNullFloat64 creates a new Float64
	NewNullFloat64 = null.NewFloat64

	// NullFloat64From creates a new Float64 that will always be valid.
	NullFloat64From = null.Float64From

	// NullFloat64FromPtr creates a new Float64 that be null if f is nil.
	NullFloat64FromPtr = null.Float64FromPtr
)

// NullInt is a nullable Int.
type NullInt = null.Int

var (
	// NewNullInt creates a new Int
	NewNullInt = null.NewInt

	// NullIntFrom creates a new Int that will always be valid.
	NullIntFrom = null.IntFrom

	// NullIntFromPtr creates a new Int that be null if f is nil.
	NullIntFromPtr = null.IntFromPtr
)

// NullInt16 is a nullable Int16.
type NullInt16 = null.Int16

var (
	// NewNullInt16 creates a new Int
	NewNullInt16 = null.NewInt16

	// NullInt16From creates a new Int that will always be valid.
	NullInt16From = null.Int16From

	// NullInt16FromPtr creates a new Int that be null if f is nil.
	NullInt16FromPtr = null.Int16FromPtr
)

// NullInt32 is a nullable Int32.
type NullInt32 = null.Int32

var (
	// NewNullInt32 creates a new Int
	NewNullInt32 = null.NewInt32

	// NullInt32From creates a new Int that will always be valid.
	NullInt32From = null.Int32From

	// NullInt32FromPtr creates a new Int that be null if f is nil.
	NullInt32FromPtr = null.Int32FromPtr
)

// NullInt64 is a nullable Int64.
type NullInt64 = null.Int64

var (
	// NewNullInt64 creates a new Int
	NewNullInt64 = null.NewInt64

	// NullInt64From creates a new Int that will always be valid.
	NullInt64From = null.Int64From

	// NullInt64FromPtr creates a new Int that be null if f is nil.
	NullInt64FromPtr = null.Int64FromPtr
)

// NullInt8 is a nullable Int8.
type NullInt8 = null.Int8

var (
	// NewNullInt8 creates a new Int
	NewNullInt8 = null.NewInt8

	// NullInt8From creates a new Int that will always be valid.
	NullInt8From = null.Int8From

	// NullInt8FromPtr creates a new Int that be null if f is nil.
	NullInt8FromPtr = null.Int8FromPtr
)

// NullJSON is a nullable JSON.
type NullJSON = null.JSON

var (
	// NewNullJSON creates a new Int
	NewNullJSON = null.NewJSON

	// NullJSONFrom creates a new Int that will always be valid.
	NullJSONFrom = null.JSONFrom

	// NullJSONFromPtr creates a new Int that be null if f is nil.
	NullJSONFromPtr = null.JSONFromPtr
)

// NullString is a nullable String.
type NullString = null.String

var (
	// NewNullString creates a new Int
	NewNullString = null.NewString

	// NullStringFrom creates a new Int that will always be valid.
	NullStringFrom = null.StringFrom

	// NullStringFromPtr creates a new Int that be null if f is nil.
	NullStringFromPtr = null.StringFromPtr
)

// NullUint is a nullable Uint.
type NullUint = null.Uint

var (
	// NewNullUint creates a new Int
	NewNullUint = null.NewUint

	// NullUintFrom creates a new Int that will always be valid.
	NullUintFrom = null.UintFrom

	// NullUintFromPtr creates a new Int that be null if f is nil.
	NullUintFromPtr = null.UintFromPtr
)

// NullUint16 is a nullable Uint16.
type NullUint16 = null.Uint16

var (
	// NewNullUint16 creates a new Int
	NewNullUint16 = null.NewUint16

	// NullUint16From creates a new Int that will always be valid.
	NullUint16From = null.Uint16From

	// NullUint16FromPtr creates a new Int that be null if f is nil.
	NullUint16FromPtr = null.Uint16FromPtr
)

// NullUint32 is a nullable Uint32.
type NullUint32 = null.Uint32

var (
	// NewNullUint32 creates a new Int
	NewNullUint32 = null.NewUint32

	// NullUint32From creates a new Int that will always be valid.
	NullUint32From = null.Uint32From

	// NullUint32FromPtr creates a new Int that be null if f is nil.
	NullUint32FromPtr = null.Uint32FromPtr
)

// NullUint64 is a nullable Uint64.
type NullUint64 = null.Uint64

var (
	// NewNullUint64 creates a new Int
	NewNullUint64 = null.NewUint64

	// NullUint64From creates a new Int that will always be valid.
	NullUint64From = null.Uint64From

	// NullUint64FromPtr creates a new Int that be null if f is nil.
	NullUint64FromPtr = null.Uint64FromPtr
)

// NullUint8 is a nullable Uint8.
type NullUint8 = null.Uint8

var (
	// NewNullUint8 creates a new Int
	NewNullUint8 = null.NewUint8

	// NullUint8From creates a new Int that will always be valid.
	NullUint8From = null.Uint8From

	// NullUint8FromPtr creates a new Int that be null if f is nil.
	NullUint8FromPtr = null.Uint8FromPtr
)

// NullTime is a nullable Time.
type NullTime = null.Time

var (
	// NewNullTime creates a new Int
	NewNullTime = null.NewTime

	// NullTimeFrom creates a new Int that will always be valid.
	NullTimeFrom = null.TimeFrom

	// NullTimeFromPtr creates a new Int that be null if f is nil.
	NullTimeFromPtr = null.TimeFromPtr
)
