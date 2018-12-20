package schema

import "github.com/gofrs/uuid"

type (
	// UUID is an array type to represent the value of a UUID, as defined in RFC-4122.
	UUID = uuid.UUID

	// NullUUID can be used with the standard sql package to represent a
	// UUID value that can be NULL in the database.
	NullUUID = uuid.NullUUID
)

var (
	// NewUUIDV1 returns a UUID based on the current timestamp and MAC address.
	NewUUIDV1 = uuid.NewV1

	// NewUUIDV2 returns a DCE Security UUID based on the POSIX UID/GID.
	NewUUIDV2 = uuid.NewV2

	// NewUUIDV3 returns a UUID based on the MD5 hash of the namespace UUID and name.
	NewUUIDV3 = uuid.NewV3

	// NewUUIDV4 returns a randomly generated UUID.
	NewUUIDV4 = uuid.NewV4

	// NewUUIDV5 returns a UUID based on SHA-1 hash of the namespace UUID and name.
	NewUUIDV5 = uuid.NewV5
)

var (
	// UUIDNil is the nil UUID, as specified in RFC-4122, that has all 128 bits set to
	// zero.
	UUIDNil = uuid.Nil

	// UUIDFromBytes returns a UUID generated from the raw byte slice input.
	// It will return an error if the slice isn't 16 bytes long.
	UUIDFromBytes = uuid.FromBytes

	// UUIDFromBytesOrNil returns a UUID generated from the raw byte slice input.
	// Same behavior as FromBytes(), but returns uuid.Nil instead of an error.
	UUIDFromBytesOrNil = uuid.FromBytesOrNil

	// UUIDFromString returns a UUID parsed from the input string.
	// Input is expected in a form accepted by UnmarshalText.
	UUIDFromString = uuid.FromString

	// UUIDFromStringOrNil returns a UUID parsed from the input string.
	// Same behavior as FromString(), but returns uuid.Nil instead of an error.
	UUIDFromStringOrNil = uuid.FromStringOrNil
)

// NewNullUUID creates a new nullable uuid
func NewNullUUID(uuid UUID, valid bool) NullUUID {
	return NullUUID{
		UUID:  uuid,
		Valid: valid,
	}
}

// NullUUIDFrom creates a new nullable UUID that will always be valid.
func NullUUIDFrom(uuid UUID) NullUUID {
	return NewNullUUID(uuid, true)
}

// NullUUIDFromPtr creates a new nullable uuid that be null if i is nil.
func NullUUIDFromPtr(uuid *UUID) NullUUID {
	if uuid == nil {
		return NewNullUUID(UUIDNil, false)
	}
	return NewNullUUID(*uuid, true)
}
