package pgmock

import (
	"bytes"
	"testing"

	. "github.com/onsi/gomega"
)

// ---------------------------------------------------------------------------------------------------------------------

func TestAuthenticationOK(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write the message to the messenger
	err := (&_AuthenticationOk{}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		AuthenticationOkMessageID,
		0, 0, 0, 8, // int32(8)
		0, 0, 0, 0, // int32(0)
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestAuthenticationKerberosV5(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write the message to the messenger
	err := (&_AuthenticationKerberosV5{}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		AuthenticationKerberosV5MessageID,
		0, 0, 0, 8, // int32(8)
		0, 0, 0, 2, // int32(2)
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestAuthenticationCleartextPassword(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write the message to the messenger
	err := (&_AuthenticationCleartextPassword{}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		AuthenticationCleartextPasswordMessageID,
		0, 0, 0, 8, // int32(8)
		0, 0, 0, 3, // int32(3)
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestAuthenticationMD5Password(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write the message to the messenger
	err := (&_AuthenticationMD5Password{}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		AuthenticationMD5PasswordMessageID,
		0, 0, 0, 12, // int32(12)
		0, 0, 0, 5, // int32(5)
		0, 0, 0, 0, // zero value of [4]byte
	}))

	// clear the buffer
	b.Reset()

	// write the message to the messenger with a test salt
	err = (&_AuthenticationMD5Password{Salt: [4]byte{1, 2, 3, 4}}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		AuthenticationMD5PasswordMessageID,
		0, 0, 0, 12, // int32(12)
		0, 0, 0, 5, // int32(5)
		1, 2, 3, 4, // [4]byte{1, 2, 3, 4}
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestAuthenticationSCMCredential(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write the message to the messenger
	err := (&_AuthenticationSCMCredential{}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		AuthenticationSCMCredentialMessageID,
		0, 0, 0, 8, // int32(8)
		0, 0, 0, 6, // int32(6)
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestAuthenticationGSS(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write the message to the messenger
	err := (&_AuthenticationGSS{}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		AuthenticationGSSMessageID,
		0, 0, 0, 8, // int32(8)
		0, 0, 0, 7, // int32(7)
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestAuthenticationSSPI(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write the message to the messenger
	err := (&_AuthenticationSSPI{}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		AuthenticationSSPIMessageID,
		0, 0, 0, 8, // int32(8)
		0, 0, 0, 9, // int32(9)
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestAuthenticationGSSContinue(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write the message to the messenger
	err := (&_AuthenticationGSSContinue{}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		AuthenticationGSSContinueMessageID,
		0, 0, 0, 4, // int32 - length of message including self
		0, 0, 0, 8, // int32(8)
		// no auth data
	}))

	// reset the buffer
	b.Reset()

	// write the message to the messenger
	err = (&_AuthenticationGSSContinue{AuthData: []byte("--bogus-auth-data--")}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		AuthenticationGSSContinueMessageID,
		0, 0, 0, 23, // int32 - length of message including self
		0, 0, 0, 8, // int32(8)
		45, 45, 98, 111, 103, 117, 115, 45, 97, 117, 116, 104, 45, 100, 97, 116, 97, 45, 45, // --bogus-auth-data--
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestAuthenticationSASL(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write the message to the messenger
	err := (&_AuthenticationSASL{}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		AuthenticationSASLMessageID,
		0, 0, 0, 4, // int32 length of message including self
		0, 0, 0, 10, // int32(10)
		// no mechanisms
	}))

	// reset the buffer
	b.Reset()

	// write the message to the messenger
	err = (&_AuthenticationSASL{Mechanisms: []string{"test1", "test2", "test3"}}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		AuthenticationSASLMessageID,
		0, 0, 0, 22, // int32 length of message including self
		0, 0, 0, 10, // int32(10)
		116, 101, 115, 116, 49, 0, // test1\x00
		116, 101, 115, 116, 50, 0, // test2\x00
		116, 101, 115, 116, 51, 0, // test3\x00
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestAuthenticationSASLContinue(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write the message to the messenger
	err := (&_AuthenticationSASLContinue{}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		AuthenticationSASLContinueMessageID,
		0, 0, 0, 4, // int32 length of message including self
		0, 0, 0, 11, // int32(11)
		// no SASLData
	}))

	// reset the buffer
	b.Reset()

	// write the message to the messenger
	err = (&_AuthenticationSASLContinue{SASLData: []byte("test")}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		AuthenticationSASLContinueMessageID,
		0, 0, 0, 8, // int32 length of message including self
		0, 0, 0, 11, // int32(11)
		116, 101, 115, 116, // test
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestAuthenticationSASLFinal(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write the message to the messenger
	err := (&_AuthenticationSASLFinal{}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		AuthenticationSASLFinalMessageID,
		0, 0, 0, 4, // int32 length of message including self
		0, 0, 0, 11, // int32(11)
		// no SASLData
	}))

	// reset the buffer
	b.Reset()

	// write the message to the messenger
	err = (&_AuthenticationSASLFinal{SASLData: []byte("test")}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		AuthenticationSASLFinalMessageID,
		0, 0, 0, 8, // int32 length of message including self
		0, 0, 0, 11, // int32(11)
		116, 101, 115, 116, // test
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestBackendKeyData(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write the message to the messenger
	err := (&_BackendKeyData{}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		BackendKeyDataMessageID,
		0, 0, 0, 12, // int32 length of message including self
		0, 0, 0, 0, // int32 process id
		0, 0, 0, 0, // int32 secret key
	}))

	// reset the buffer
	b.Reset()

	// write the message to the messenger
	err = (&_BackendKeyData{ProcessID: 1234, SecretKey: 4321}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		BackendKeyDataMessageID,
		0, 0, 0, 12, // int32 length of message including self
		0, 0, 4, 210, // int32 process id
		0, 0, 16, 225, // int32 secret key
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestBind(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// create a dummy message to use
	b, m := createBufMesPair()
	m.writeInt32(0). // msg len - ignored
				writeString("--portal--").
				writeString("--statement--").
				writeInt16(2).   // n format codes
				writeInt16(123). // first format code
				writeInt16(234). // second format code
				writeInt16(2).   // n parameters
				writeInt32(9).   // first param length
				writeByteArray([]byte("--valu1--")...).
				writeInt32(9). // second param length
				writeByteArray([]byte("--valu2--")...).
				writeInt16(2).   // n result codes
				writeInt16(123). // first result format code
				writeInt16(234)  // second result format code

	// create a new reader to read the message
	m = newMessenger(bytes.NewBuffer(b.Bytes()))

	// create the message and attempt to read the data into it
	msg := &_Bind{}
	err := msg.read(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(msg.DestinationPortal).To(Equal("--portal--"))
	Expect(msg.PreparedStatement).To(Equal("--statement--"))
	Expect(len(msg.Parameters)).To(Equal(2))
	Expect(msg.Parameters[0].FormatCode).To(Equal(int16(123)))
	Expect(msg.Parameters[0].ResultFormatCode).To(Equal(int16(123)))
	Expect(msg.Parameters[0].Value).To(Equal([]byte("--valu1--")))
	Expect(msg.Parameters[1].FormatCode).To(Equal(int16(234)))
	Expect(msg.Parameters[1].ResultFormatCode).To(Equal(int16(234)))
	Expect(msg.Parameters[1].Value).To(Equal([]byte("--valu2--")))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestBindComplete(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write the message to the messenger
	err := (&_BindComplete{}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		BindCompleteMessageID,
		0, 0, 0, 4, // int32(4)
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestCancelRequest(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// create a dummy message to use
	b, m := createBufMesPair()
	m.writeInt32(12345).
		writeInt32(54321)

	// create a new reader to read the message
	m = newMessenger(bytes.NewBuffer(b.Bytes()))

	// create the message and attempt to read the data into it
	msg := &_CancelRequest{}
	err := msg.read(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(msg.ProcessID).To(Equal(int32(12345)))
	Expect(msg.SecretKey).To(Equal(int32(54321)))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestClose(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// create a dummy message to use
	b, m := createBufMesPair()
	m.writeInt32(0).
		writeByte(CloseStatement).
		writeString("--statement--")

	// create a new reader to read the message
	m = newMessenger(bytes.NewBuffer(b.Bytes()))

	// create the message and attempt to read the data into it
	msg := &_Close{}
	err := msg.read(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(msg.CloseType).To(Equal(byte(CloseStatement)))
	Expect(msg.Name).To(Equal("--statement--"))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestCloseComplete(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write the message to the messenger
	err := (&_CloseComplete{}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		CloseCompleteMessageID,
		0, 0, 0, 4, // int32(4)
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestCommandComplete(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// create the message
	msg := &_CommandComplete{}

	// write the message to the messenger
	err := msg.write(m)

	// assert the results - should fail as no tag is present
	Expect(err).ToNot(BeNil())

	// write the message to the messenger
	b.Reset()
	msg.insert(1, 2)
	err = msg.write(m)
	Expect(err).To(BeNil())
	Expect(msg.Tag).To(Equal("INSERT 1 2"))

	// write the message to the messenger
	b.Reset()
	msg.delete(1)
	err = msg.write(m)
	Expect(err).To(BeNil())
	Expect(msg.Tag).To(Equal("DELETE 1"))

	// write the message to the messenger
	b.Reset()
	msg.update(1)
	err = msg.write(m)
	Expect(err).To(BeNil())
	Expect(msg.Tag).To(Equal("UPDATE 1"))

	// write the message to the messenger
	b.Reset()
	msg.selectOrCreate(1)
	err = msg.write(m)
	Expect(err).To(BeNil())
	Expect(msg.Tag).To(Equal("SELECT 1"))

	// write the message to the messenger
	b.Reset()
	msg.move(1)
	err = msg.write(m)
	Expect(err).To(BeNil())
	Expect(msg.Tag).To(Equal("MOVE 1"))

	// write the message to the messenger
	b.Reset()
	msg.fetch(1)
	err = msg.write(m)
	Expect(err).To(BeNil())
	Expect(msg.Tag).To(Equal("FETCH 1"))

	// write the message to the messenger
	b.Reset()
	msg.copy(1)
	err = msg.write(m)
	Expect(err).To(BeNil())
	Expect(msg.Tag).To(Equal("COPY 1"))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestCopyData(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestCopyDone(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestCopyFail(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestCopyInResponse(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestCopyOutResponse(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestBothResponse(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestDataRow(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// create a simple data rowdescription
	msg := &_DataRow{Columns: []*_DataRowColumn{
		&_DataRowColumn{
			Value: []byte("test"),
		},
	}}

	// write the message to the messenger
	err := msg.write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		DataRowMessageID,
		0, 0, 0, 14, // Length
		0, 1, // Column Count
		0, 0, 0, 4, // Column Value Length
		116, 101, 115, 116, // Column Value 'test'
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestDescribe(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// create a dummy message to use
	b, m := createBufMesPair()
	m.writeByte(CloseStatement).
		writeString("--statement--")

	// create a new reader to read the message
	m = newMessenger(bytes.NewBuffer(b.Bytes()))

	// create the message and attempt to read the data into it
	msg := &_Describe{}
	err := msg.read(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(msg.Target).To(Equal(byte(CloseStatement)))
	Expect(msg.TargetName).To(Equal("--statement--"))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestEmptyQueryResponse(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestErrorResponse(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// create the message
	msg := &_ErrorResponse{}
	msg.AddErrorField('S', "ERROR").
		AddErrorField('V', "ERROR").
		AddErrorField('C', "57014").
		AddErrorField('M', "--error-string--")

	// write the message to the messenger
	err := msg.write(m)

	// assert the results - should fail as no tag is present
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		ErrorResponseMessageID,
		0, 0, 0, 44, // int32 length of message including self
		'S',                   // byte1 severity
		69, 82, 82, 79, 82, 0, // ERROR\x00
		'V',                   // byte1 9.6+ severity
		69, 82, 82, 79, 82, 0, // ERROR\x00
		'C',                   // Byte1 SQLState code
		53, 55, 48, 49, 52, 0, // 57014 - cancel request
		'M',                                                                          // Byte1 Message
		45, 45, 101, 114, 114, 111, 114, 45, 115, 116, 114, 105, 110, 103, 45, 45, 0, // --error--string--\x00
		0, // message termination \x00
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestExecute(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestFlush(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestFunctionCall(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestFunctionCallResponse(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestGSSResponse(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestNegotiateProtocolVersion(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestNoData(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestNoticeResponse(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestNotificationResponse(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestParameterDescription(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// create a simple data rowdescription
	msg := &_ParameterDescription{}

	// write the message to the messenger
	err := msg.write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		ParameterDescriptionMessageID,
		0, 0, 0, 6, // Length
		0, 0, //param oid Count
	}))

	// reset
	b.Reset()

	// add some parameter oids
	msg.ParameterOIDS = []int32{1, 2, 3, 4, 5}

	// write the message to the messenger
	err = msg.write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		ParameterDescriptionMessageID,
		0, 0, 0, 26, // Length
		0, 5, // param oid Count
		0, 0, 0, 1, // int32(1)
		0, 0, 0, 2, // int32(2)
		0, 0, 0, 3, // int32(3)
		0, 0, 0, 4, // int32(4)
		0, 0, 0, 5, // int32(5)
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestParameterStatus(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestParse(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// create a dummy message to use
	b, m := createBufMesPair()
	m.writeString("--statement--").
		writeString("SELECT * FROM AWESOME").
		writeInt16(1).
		writeInt32(123)

	// create a new reader to read the message
	m = newMessenger(bytes.NewBuffer(b.Bytes()))

	// create the message and attempt to read the data into it
	msg := &_Parse{}
	err := msg.read(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(msg.Statement).To(Equal("--statement--"))
	Expect(msg.SQL).To(Equal("SELECT * FROM AWESOME"))
	Expect(len(msg.ParameterOIDs)).To(Equal(1))
	Expect(msg.ParameterOIDs[0]).To(Equal(int32(123)))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestParseComplete(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write the message to the messenger
	err := (&_ParseComplete{}).write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		ParseCompleteMessageID,
		0, 0, 0, 4, // int32(4)
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestPasswordMessage(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestPortalSuspended(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestQuery(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestReadyForQuery(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// write the message to the messenger
	err := (&_ReadyForQuery{}).write(m)

	// assert the results
	Expect(err).ToNot(BeNil())

	// reset the buffer
	b.Reset()

	// write the message to the messenger
	err = (&_ReadyForQuery{Indicator: 'Z'}).write(m) // Z is not known as an indicator
	Expect(err).ToNot(BeNil())

	// reset the buffer
	b.Reset()

	// write the message to the messenger
	err = (&_ReadyForQuery{Indicator: ReadyForQueryIdle}).write(m)
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		ReadyForQueryMessageID,
		0, 0, 0, 5, // int32(5)
		ReadyForQueryIdle,
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestRowDescription(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// grab the test objects
	b, m := createBufMesPair()

	// create a simple Row description
	msg := &_RowDescription{Fields: []*_RowDescriptionField{
		&_RowDescriptionField{
			Name:         "test",
			TableOID:     0,
			ColumnAttr:   0,
			DataTypeOID:  25,
			DataTypeSize: -1,
			TypeModifier: 0,
			FormatCode:   0,
		},
	}}

	// write the message to the messenger
	err := msg.write(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(b.Bytes()).To(Equal([]byte{
		RowDescriptionMessageID,
		0, 0, 0, 29, // Length
		0, 1, // Field Count
		116, 101, 115, 116, 0, // Field Name 'test\x00'
		0, 0, 0, 0, // TableOID
		0, 0, // ColumnAttr
		0, 0, 0, 25, // DataTypeOID
		255, 255, // DataTypeSize
		0, 0, 0, 0, // TypeModifier
		0, 0, // FormatCode
	}))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestSASLInitialResponse(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestSASLResponse(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestStartupMessage(t *testing.T) {

	// gomega requirement
	RegisterTestingT(t)

	// create a dummy message to use
	b, m := createBufMesPair()
	m.writeString("key1").writeString("val1").
		writeString("key2").writeString("val2")

	// create a new reader to read the message
	m = newMessenger(bytes.NewBuffer(b.Bytes()))

	// create the message and attempt to read the data into it
	msg := &_StartupMessage{}
	err := msg.read(m)

	// assert the results
	Expect(err).To(BeNil())
	Expect(len(msg.Parameters)).To(Equal(2))
	Expect(msg.Parameters["key1"]).To(Equal("val1"))
	Expect(msg.Parameters["key2"]).To(Equal("val2"))
}

// ---------------------------------------------------------------------------------------------------------------------

func TestSync(t *testing.T) {}

// ---------------------------------------------------------------------------------------------------------------------

func TestTerminate(t *testing.T) {}
