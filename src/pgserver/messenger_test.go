package pgmock

import (
	"bytes"
	"testing"

	. "github.com/onsi/gomega"
)

// ---------------------------------------------------------------------------------------------------------------------

func createBufMesPair() (*bytes.Buffer, *_Messenger) {
	b := &bytes.Buffer{}
	return b, newMessenger(b)
}

// ---------------------------------------------------------------------------------------------------------------------

func TestWriteByte(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write an int16 to the messenger
	m.writeByte(127)

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{127}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestWriteByteArray(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write an int16 to the messenger
	m.writeByteArray(10, 20, 30)

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{10, 20, 30}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestWriteByte4(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write an int16 to the messenger
	m.writeByte4([4]byte{10, 20, 30, 40})

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{10, 20, 30, 40}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestWriteInt8(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write an int16 to the messenger
	m.writeInt8(127)

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{127}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestWriteInt8Array(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write an int16 to the messenger
	m.writeInt8Array(10, 20, 30)

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{10, 20, 30}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestWriteInt16(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write an int16 to the messenger
	m.writeInt16(123)

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{0, 123}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestWriteInt16Array(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write an int16 to the messenger
	m.writeInt16Array(10, 20, 30)

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{0, 10, 0, 20, 0, 30}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestWriteInt32(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write an int16 to the messenger
	m.writeInt32(234)

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{0, 0, 0, 234}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestWriteInt32Array(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write an int16 to the messenger
	m.writeInt32Array(10, 20, 30)

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{0, 0, 0, 10, 0, 0, 0, 20, 0, 0, 0, 30}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestWriteInt64(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write an int16 to the messenger
	m.writeInt64(345)

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{0, 0, 0, 0, 0, 0, 1, 89}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestWriteInt64Array(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write an int16 to the messenger
	m.writeInt64Array(10, 20, 30)

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{0, 0, 0, 0, 0, 0, 0, 10, 0, 0, 0, 0, 0, 0, 0, 20, 0, 0, 0, 0, 0, 0, 0, 30}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestWriteString(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write an int16 to the messenger
	m.writeString("test")

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte("test\x00")))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestWriteStringArray(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write an int16 to the messenger
	m.writeStringArray("test1", "test2")

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte("test1\x00test2\x00")))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestReadByte(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b := bytes.NewBufferString("A")
	m := newMessenger(b)

	// read a byte
	fb := m.readByte()

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(fb).To(Equal(byte('A')))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestReadBytes(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b := bytes.NewBufferString("--testdata--")
	m := newMessenger(b)

	// read 4 bytes
	fb := m.readBytes(4)

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(fb).To(Equal([]byte("--te")))

	// read the next 8 bytes
	fb = m.readBytes(8)

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(fb).To(Equal([]byte("stdata--")))

	// read the next 0 bytes
	fb = m.readBytes(0)

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(fb).To(Equal([]byte{}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestReadInt16(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b := bytes.NewBuffer([]byte{0, 123})
	m := newMessenger(b)

	// read a byte
	res := m.readInt16()

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(res).To(Equal(int16(123)))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestReadInt32(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b := bytes.NewBuffer([]byte{0, 0, 0, 123})
	m := newMessenger(b)

	// read a byte
	res := m.readInt32()

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(res).To(Equal(int32(123)))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestReadString(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b := bytes.NewBuffer([]byte("test1\x00test2\x00"))
	m := newMessenger(b)

	// read a string
	res := m.readString()

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(res).To(Equal("test1"))

	// read another string
	res = m.readString()

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(res).To(Equal("test2"))

	// read another string
	res = m.readString()

	// check it's correct in the buffer
	Expect(m.Error).To(BeNil())
	Expect(res).To(Equal(""))
}
