package ref

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RefSuite struct {
	suite.Suite
}

func (r *RefSuite) TestBool() {
	assert.NotNil(r.T(), Bool(true))
}

func (r *RefSuite) TestByte() {
	assert.NotNil(r.T(), Byte('a'))
}

func (r *RefSuite) TestComplex64() {
	assert.NotNil(r.T(), Complex64(1))
}

func (r *RefSuite) TestComplex128() {
	assert.NotNil(r.T(), Complex128(1))
}

func (r *RefSuite) TestFloat32() {
	assert.NotNil(r.T(), Float32(0.1))
}

func (r *RefSuite) TestFloat64() {
	assert.NotNil(r.T(), Float64(0.1))
}

func (r *RefSuite) TestInt() {
	assert.NotNil(r.T(), Int(1))
}

func (r *RefSuite) TestInt8() {
	assert.NotNil(r.T(), Int8(1))
}

func (r *RefSuite) TestInt16() {
	assert.NotNil(r.T(), Int16(1))
}

func (r *RefSuite) TestInt32() {
	assert.NotNil(r.T(), Int32(1))
}

func (r *RefSuite) TestInt64() {
	assert.NotNil(r.T(), Int64(1))
}

func (r *RefSuite) TestRune() {
	assert.NotNil(r.T(), Rune('a'))
}

func (r *RefSuite) TestString() {
	assert.NotNil(r.T(), String("abc"))
}

func (r *RefSuite) TestUint() {
	assert.NotNil(r.T(), Uint(1))
}

func (r *RefSuite) TestUint8() {
	assert.NotNil(r.T(), Uint8(1))
}

func (r *RefSuite) TestUint16() {
	assert.NotNil(r.T(), Uint16(1))
}

func (r *RefSuite) TestUint32() {
	assert.NotNil(r.T(), Uint32(1))
}

func (r *RefSuite) TestUint64() {
	assert.NotNil(r.T(), Uint64(1))
}

func (r *RefSuite) TestDuration() {
	assert.NotNil(r.T(), Duration(1000*time.Millisecond))
}

func TestRefSuite(t *testing.T) {
	suite.Run(t, new(RefSuite))
}
