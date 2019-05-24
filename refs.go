// Package ref implements primitive type pointer de-referencing function.
package ref

import (
	"time"
)

// Bool returns a pointer to a bool
func Bool(value bool) *bool {
	return &value
}

// Byte returns a pointer to a byte
func Byte(value byte) *byte {
	return &value
}

// Complex64 returns a pointer to a complex64
func Complex64(value complex64) *complex64 {
	return &value
}

// Complex128 returns a pointer to a complex128
func Complex128(value complex128) *complex128 {
	return &value
}

// Float32 returns a pointer to a float32
func Float32(value float32) *float32 {
	return &value
}

// Float64 returns a pointer to a float64
func Float64(value float64) *float64 {
	return &value
}

// Int returns a pointer to a int
func Int(value int) *int {
	return &value
}

// Int8 returns a pointer to a int8
func Int8(value int8) *int8 {
	return &value
}

// Int16 returns a pointer to a int16
func Int16(value int16) *int16 {
	return &value
}

// Int32 returns a pointer to a int32
func Int32(value int32) *int32 {
	return &value
}

// Int64 returns a pointer to a int64
func Int64(value int64) *int64 {
	return &value
}

// Rune returns a pointer to a rune
func Rune(value rune) *rune {
	return &value
}

// String returns a pointer to a string
func String(value string) *string {
	return &value
}

// Uint returns a pointer to a uint
func Uint(value uint) *uint {
	return &value
}

// Uint8 returns a pointer to a uint8
func Uint8(value uint8) *uint8 {
	return &value
}

// Uint16 returns a pointer to a uint16
func Uint16(value uint16) *uint16 {
	return &value
}

// Uint32 returns a pointer to a uint32
func Uint32(value uint32) *uint32 {
	return &value
}

// Uint64 returns a pointer to a uint64
func Uint64(value uint64) *uint64 {
	return &value
}

// Duration returns a pointer to a duration
func Duration(value time.Duration) *time.Duration {
	return &value
}
