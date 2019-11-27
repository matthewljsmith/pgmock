package pgmock

import (
	"encoding/binary"
	"io"
)

// ---------------------------------------------------------------------------------------------------------------------

type _Messenger struct {
	Stream io.ReadWriter
	Error  error
}

// ---------------------------------------------------------------------------------------------------------------------

func newMessenger(stream io.ReadWriter) *_Messenger {
	return &_Messenger{
		Stream: stream,
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *_Messenger) writeByte(v byte) *_Messenger {
	m.Error = binary.Write(m.Stream, binary.BigEndian, v)
	return m
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *_Messenger) writeByteArray(vs ...byte) *_Messenger {
	for _, v := range vs {
		m.writeByte(v)
	}
	return m
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *_Messenger) writeByte4(vs [4]byte) *_Messenger {
	for _, v := range vs {
		m.writeByte(v)
	}
	return m
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *_Messenger) writeInt8(v int8) *_Messenger {
	m.Error = binary.Write(m.Stream, binary.BigEndian, v)
	return m
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *_Messenger) writeInt8Array(vs ...int8) *_Messenger {
	for _, v := range vs {
		m.writeInt8(v)
	}
	return m
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *_Messenger) writeInt16(v int16) *_Messenger {
	m.Error = binary.Write(m.Stream, binary.BigEndian, v)
	return m
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *_Messenger) writeInt16Array(vs ...int16) *_Messenger {
	for _, v := range vs {
		m.writeInt16(v)
	}
	return m
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *_Messenger) writeInt32(v int32) *_Messenger {
	m.Error = binary.Write(m.Stream, binary.BigEndian, v)
	return m
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *_Messenger) writeInt32Array(vs ...int32) *_Messenger {
	for _, v := range vs {
		m.writeInt32(v)
	}
	return m
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *_Messenger) writeInt64(v int64) *_Messenger {
	m.Error = binary.Write(m.Stream, binary.BigEndian, v)
	return m
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *_Messenger) writeInt64Array(vs ...int64) *_Messenger {
	for _, v := range vs {
		m.writeInt64(v)
	}
	return m
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *_Messenger) writeString(v string) *_Messenger {
	b := []byte(v)
	b = append(b, 0)
	m.Error = binary.Write(m.Stream, binary.BigEndian, b)
	return m
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *_Messenger) writeStringArray(vs ...string) *_Messenger {
	for _, v := range vs {
		m.writeString(v)
	}
	return m
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *_Messenger) readByte() (res byte) {
	m.Error = binary.Read(m.Stream, binary.BigEndian, &res)
	return res
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *_Messenger) readBytes(len int32) []byte {

	// no-op
	if len == 0 {
		return []byte{}
	}

	// otherwise make a []byte the correct length
	res := make([]byte, len)
	m.Error = binary.Read(m.Stream, binary.BigEndian, &res)
	return res
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *_Messenger) readInt16() (res int16) {
	m.Error = binary.Read(m.Stream, binary.BigEndian, &res)
	return res
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *_Messenger) readInt32() (res int32) {
	m.Error = binary.Read(m.Stream, binary.BigEndian, &res)
	return res
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *_Messenger) readString() string {

	var buf []byte
	var b byte
	for {
		m.Error = binary.Read(m.Stream, binary.BigEndian, &b)
		if m.Error == io.EOF {
			m.Error = nil
			break
		}
		if m.Error != nil {
			return ""
		}
		if b == 0 {
			break
		}
		buf = append(buf, b)
	}
	return string(buf)
}
