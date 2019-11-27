package pgmock

import (
	"fmt"
)

// Message formats from PG 10.0
// https://www.postgresql.org/docs/current/protocol-message-formats.html

// ---------------------------------------------------------------------------------------------------------------------

// AuthenticationOk (B)

// Byte1('R') 	Identifies the message as an authentication request.
// Int32(8)		Length of message contents in bytes, including self.
// Int32(0)		Specifies that the authentication was successful.

type _AuthenticationOk struct{}

func (pgm *_AuthenticationOk) write(m *_Messenger) error {
	m.writeByte(AuthenticationOkMessageID).writeInt32(8).writeInt32(0)
	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// AuthenticationKerberosV5 (B)

// Byte1('R')	Identifies the message as an authentication request.
// Int32(8)	Length of message contents in bytes, including self.
// Int32(2)	Specifies that Kerberos V5 authentication is required.

type _AuthenticationKerberosV5 struct{}

func (pgm *_AuthenticationKerberosV5) write(m *_Messenger) error {
	m.writeByte(AuthenticationKerberosV5MessageID).writeInt32(8).writeInt32(2)
	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// AuthenticationCleartextPassword (B)

// Byte1('R')	Identifies the message as an authentication request.
// Int32(8)		Length of message contents in bytes, including self.
// Int32(3)		Specifies that a clear-text password is required.

type _AuthenticationCleartextPassword struct{}

func (pgm *_AuthenticationCleartextPassword) write(m *_Messenger) error {
	m.writeByte(AuthenticationCleartextPasswordMessageID).writeInt32(8).writeInt32(3)
	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// AuthenticationMD5Password (B)

// Byte1('R') 	Identifies the message as an authentication request.
// Int32(12) 	Length of message contents in bytes, including self.
// Int32(5) 	Specifies that an MD5-encrypted password is required.
// Byte4 		The salt to use when encrypting the password.

type _AuthenticationMD5Password struct {
	Salt [4]byte
}

func (pgm *_AuthenticationMD5Password) write(m *_Messenger) error {
	m.writeByte(AuthenticationMD5PasswordMessageID).writeInt32(12).writeInt32(5).writeByte4(pgm.Salt)
	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// AuthenticationSCMCredential (B)

// Byte1('R')  Identifies the message as an authentication request.
// Int32(8)    Length of message contents in bytes, including self.
// Int32(6)    Specifies that an SCM credentials message is required.

type _AuthenticationSCMCredential struct{}

func (pgm *_AuthenticationSCMCredential) write(m *_Messenger) error {
	m.writeByte(AuthenticationSCMCredentialMessageID).writeInt32(8).writeInt32(6)
	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// AuthenticationGSS (B)

// Byte1('R')  Identifies the message as an authentication request.
// Int32(8)    Length of message contents in bytes, including self.
// Int32(7)    Specifies that GSSAPI authentication is required.

type _AuthenticationGSS struct{}

func (pgm *_AuthenticationGSS) write(m *_Messenger) error {
	m.writeByte(AuthenticationGSSMessageID).writeInt32(8).writeInt32(7)
	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// AuthenticationSSPI (B)

// Byte1('R')  Identifies the message as an authentication request.
// Int32(8)    Length of message contents in bytes, including self.
// Int32(9)    Specifies that SSPI authentication is required.

type _AuthenticationSSPI struct{}

func (pgm *_AuthenticationSSPI) write(m *_Messenger) error {
	m.writeByte(AuthenticationSSPIMessageID).writeInt32(8).writeInt32(9)
	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// AuthenticationGSSContinue (B)

// Byte1('R')  Identifies the message as an authentication request.
// Int32       Length of message contents in bytes, including self.
// Int32(8)    Specifies that this message contains GSSAPI or SSPI data.
// Byten       GSSAPI or SSPI authentication data.

type _AuthenticationGSSContinue struct {
	AuthData []byte
}

func (pgm *_AuthenticationGSSContinue) write(m *_Messenger) error {
	m.writeByte(AuthenticationGSSContinueMessageID).writeInt32(int32(4 + len(pgm.AuthData))).writeInt32(8).writeByteArray(pgm.AuthData...)
	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// AuthenticationSASL (B)

// Byte1('R')  Identifies the message as an authentication request.
// Int32       Length of message contents in bytes, including self.
// Int32(10)   Specifies that SASL authentication is required.

// The message body is a list of SASL authentication mechanisms, in the server's order of preference. A zero byte is
// required as terminator after the last authentication mechanism name. For each mechanism, there is the following:

// String      Name of a SASL authentication mechanism.

type _AuthenticationSASL struct {
	Mechanisms []string
}

func (pgm *_AuthenticationSASL) write(m *_Messenger) error {

	// work out the message length
	mechanisms := []byte{}
	for _, s := range pgm.Mechanisms {
		mechanisms = append(mechanisms, []byte(s+"\x00")...)
	}

	// write the message out
	m.writeByte(AuthenticationSASLMessageID).writeInt32(int32(4 + len(mechanisms))).writeInt32(10).writeByteArray(mechanisms...)
	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// AuthenticationSASLContinue (B)

// Byte1('R')	Identifies the message as an authentication request.
// Int32		Length of message contents in bytes, including self.
// Int32(11)	Specifies that this message contains a SASL challenge.
// Byten		SASL data, specific to the SASL mechanism being used.

type _AuthenticationSASLContinue struct {
	SASLData []byte
}

func (pgm *_AuthenticationSASLContinue) write(m *_Messenger) error {
	m.writeByte(AuthenticationSASLContinueMessageID).writeInt32(int32(4 + len(pgm.SASLData))).writeInt32(11).writeByteArray(pgm.SASLData...)
	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// AuthenticationSASLFinal (B)

// Byte1('R')	Identifies the message as an authentication request.
// Int32		Length of message contents in bytes, including self.
// Int32(12)	Specifies that SASL authentication has completed.
// Byten		SASL outcome "additional data", specific to the SASL mechanism being used.

type _AuthenticationSASLFinal struct {
	SASLData []byte
}

func (pgm *_AuthenticationSASLFinal) write(m *_Messenger) error {
	m.writeByte(AuthenticationSASLFinalMessageID).writeInt32(int32(4 + len(pgm.SASLData))).writeInt32(11).writeByteArray(pgm.SASLData...)
	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// BackendKeyData (B)

// Byte1('K')	Identifies the message as cancellation key data. The frontend must save these values if it wishes to be able to issue CancelRequest messages later.
// Int32(12)	Length of message contents in bytes, including self.
// Int32		The process ID of this backend.
// Int32		The secret key of this backend.

type _BackendKeyData struct {
	ProcessID int32
	SecretKey int32
}

func (pgm *_BackendKeyData) write(m *_Messenger) error {
	m.writeByte(BackendKeyDataMessageID).writeInt32(12).writeInt32(pgm.ProcessID).writeInt32(pgm.SecretKey)
	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// Bind (F)

// Byte1('B')	Identifies the message as a Bind command.
// Int32		Length of message contents in bytes, including self.
// String		The name of the destination portal (an empty string selects the unnamed portal).
// String		The name of the source prepared statement (an empty string selects the unnamed prepared statement).
// Int16		The number of parameter format codes that follow (denoted C below). This can be zero to indicate that
//				there are no parameters or that the parameters all use the default format (text); or one, in which case
//				the specified format code is applied to all parameters; or it can equal the actual number of parameters.
// Int16[C]	The parameter format codes. Each must presently be zero (text) or one (binary).
// Int16		The number of parameter values that follow (possibly zero). This must match the number of parameters
//				needed by the query.

// Next, the following pair of fields appear for each parameter:

// Int32		The length of the parameter value, in bytes (this count does not include itself). Can be zero. As a
//				special case, -1 indicates a NULL parameter value. No value bytes follow in the NULL case.
// Byten		The value of the parameter, in the format indicated by the associated format code. n is the above length.

// After the last parameter, the following fields appear:

// Int16		The number of result-column format codes that follow (denoted R below). This can be zero to indicate
//				that there are no result columns or that the result columns should all use the default format (text);
//				or one, in which case the specified format code is applied to all result columns (if any); or it can
//				equal the actual number of result columns of the query.
// Int16[R]	The result-column format codes. Each must presently be zero (text) or one (binary).

type _BindParameter struct {
	Value            []byte
	FormatCode       int16
	ResultFormatCode int16
}

type _Bind struct {
	DestinationPortal string
	PreparedStatement string
	Parameters        []*_BindParameter
}

func (pgm *_Bind) read(m *_Messenger) error {

	// ignore the first byte as it's already been fetched to determine message type

	// grab the message length and discard..
	m.readInt32()
	if m.Error != nil {
		return fmt.Errorf("unable to read message length, err: %s", m.Error)
	}

	// grab the destination portal
	pgm.DestinationPortal = m.readString()
	if m.Error != nil {
		return fmt.Errorf("unable to read destination portal, err: %s", m.Error)
	}

	// grab the source prepared statement
	pgm.PreparedStatement = m.readString()
	if m.Error != nil {
		return fmt.Errorf("unable to read prepared statement ID, err: %s", m.Error)
	}

	// get the number of format codes
	nFormatCodes := m.readInt16()
	if m.Error != nil {
		return fmt.Errorf("unable to read nFormatCodes, err: %s", m.Error)
	}

	formatCodes := make([]int16, nFormatCodes)
	for i := int16(0); i < nFormatCodes; i++ {
		formatCodes[i] = m.readInt16()
		if m.Error != nil {
			return fmt.Errorf("unable to read formatCodes %d, err: %s", i, m.Error)
		}
	}

	// helper function to the right format code
	getFormatCode := func(idx int) int16 {
		if nFormatCodes == 0 {
			return -1
		} else if len(formatCodes) < idx {
			return formatCodes[0]
		} else {
			return formatCodes[idx]
		}
	}

	// get the number of query parameters
	nParameters := m.readInt16()
	if m.Error != nil {
		return fmt.Errorf("unable to read nParameters, err: %s", m.Error)
	}

	// iterate through the parameters
	pgm.Parameters = make([]*_BindParameter, nParameters)
	for i := int16(0); i < nParameters; i++ {

		// create a new param
		param := &_BindParameter{
			FormatCode: getFormatCode(int(i)),
		}

		// read the length of the value
		valLen := m.readInt32()
		if m.Error != nil {
			return fmt.Errorf("unable to read parameters %d length, err: %s", i, m.Error)
		}

		// read in the vaue
		param.Value = m.readBytes(valLen)
		if m.Error != nil {
			return fmt.Errorf("unable to read parameters %d value, err: %s", i, m.Error)
		}

		// add it to the param list
		pgm.Parameters[i] = param
	}

	// read the result column format codes
	nResultFormatCodes := m.readInt16()
	if m.Error != nil {
		return fmt.Errorf("unable to read nResultFormatCodes, err: %s", m.Error)
	}

	// read in the format codes
	resultFormatCodes := make([]int16, nResultFormatCodes)
	for i := int16(0); i < nResultFormatCodes; i++ {
		resultFormatCodes[i] = m.readInt16()
		if m.Error != nil {
			return fmt.Errorf("unable to read resultFormatCode %d, err: %s", i, m.Error)
		}
	}

	// helper function to the right result column format code
	getResultFormatCode := func(idx int) int16 {
		if nResultFormatCodes == 0 {
			return -1
		} else if len(resultFormatCodes) < idx {
			return resultFormatCodes[0]
		} else {
			return resultFormatCodes[idx]
		}
	}

	// set the result format codes
	for i := int16(0); i < nParameters; i++ {
		pgm.Parameters[i].ResultFormatCode = getResultFormatCode(int(i))
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

// BindComplete (B)

// Byte1('2')	Identifies the message as a Bind-complete indicator.
// Int32(4)		Length of message contents in bytes, including self.

type _BindComplete struct{}

func (pgm *_BindComplete) write(m *_Messenger) error {
	m.writeByte(BindCompleteMessageID).writeInt32(4)
	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// CancelRequest (F)

// Int32(16)			Length of message contents in bytes, including self.
// Int32(80877102)		The cancel request code. The value is chosen to contain 1234 in the most significant 16 bits,
//						and 5678 in the least significant 16 bits. (To avoid confusion, this code must not be the same
//						as any protocol version number.)
// Int32				The process ID of the target backend.
// Int32				The secret key for the target backend.

type _CancelRequest struct {
	ProcessID int32
	SecretKey int32
}

func (pgm *_CancelRequest) read(m *_Messenger) error {

	// this one is a bit special in that it has no Byte1 identifier
	// the LSB/MSB are calculated in the message loop so all we need
	// to read is the procID and secret key
	pgm.ProcessID = m.readInt32()
	if m.Error != nil {
		return m.Error
	}
	pgm.SecretKey = m.readInt32()
	if m.Error != nil {
		return m.Error
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

// Close (F)

// Byte1('C')	Identifies the message as a Close command.
// Int32		Length of message contents in bytes, including self.
// Byte1		'S' to close a prepared statement; or 'P' to close a portal.
// String		The name of the prepared statement or portal to close (an empty string selects the unnamed prepared statement or portal).

type _Close struct {
	CloseType byte
	Name      string
}

func (pgm *_Close) read(m *_Messenger) error {

	// read and discard the message length
	m.readInt32()
	if m.Error != nil {
		return m.Error
	}

	// read the close type in
	pgm.CloseType = m.readByte()
	if m.Error != nil {
		return m.Error
	}
	pgm.Name = m.readString()
	if m.Error != nil {
		return m.Error
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

// CloseComplete (B)

// Byte1('3')	Identifies the message as a Close-complete indicator.
// Int32(4)		Length of message contents in bytes, including self.

type _CloseComplete struct{}

func (pgm *_CloseComplete) write(m *_Messenger) error {
	m.writeByte(CloseCompleteMessageID).writeInt32(4)
	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// CommandComplete (B)

// Byte1('C')	Identifies the message as a command-completed response.
// Int32		Length of message contents in bytes, including self.
// String		The command tag. This is usually a single word that identifies which SQL command was completed.
// 			For an INSERT command, the tag is INSERT oid rows, where rows is the number of rows inserted. oid is the
// 				object ID of the inserted row if rows is 1 and the target table has OIDs; otherwise oid is 0.
// 			For a DELETE command, the tag is DELETE rows where rows is the number of rows deleted.
// 			For an UPDATE command, the tag is UPDATE rows where rows is the number of rows updated.
// 			For a SELECT or CREATE TABLE AS command, the tag is SELECT rows where rows is the number of rows retrieved.
// 			For a MOVE command, the tag is MOVE rows where rows is the number of rows the cursor's position has been changed by.
// 			For a FETCH command, the tag is FETCH rows where rows is the number of rows that have been retrieved from the cursor.
// 			For a COPY command, the tag is COPY rows where rows is the number of rows copied. (Note: the row count appears only in PostgreSQL 8.2 and later.)

type _CommandComplete struct {
	Tag string
}

func (pgm *_CommandComplete) insert(oid, rows int)    { pgm.Tag = fmt.Sprintf("INSERT %d %d", oid, rows) }
func (pgm *_CommandComplete) delete(rows int)         { pgm.Tag = fmt.Sprintf("DELETE %d", rows) }
func (pgm *_CommandComplete) update(rows int)         { pgm.Tag = fmt.Sprintf("UPDATE %d", rows) }
func (pgm *_CommandComplete) selectOrCreate(rows int) { pgm.Tag = fmt.Sprintf("SELECT %d", rows) }
func (pgm *_CommandComplete) move(rows int)           { pgm.Tag = fmt.Sprintf("MOVE %d", rows) }
func (pgm *_CommandComplete) fetch(rows int)          { pgm.Tag = fmt.Sprintf("FETCH %d", rows) }
func (pgm *_CommandComplete) copy(rows int)           { pgm.Tag = fmt.Sprintf("COPY %d", rows) }

func (pgm *_CommandComplete) write(m *_Messenger) error {

	if pgm.Tag == "" {
		return fmt.Errorf("_CommandComplete.Tag was empty, please call one of the insert/delete/update functions before writing")
	}

	// length is int32 + string + \x00
	m.writeByte(CommandCompleteMessageID).writeInt32(int32(4 + len(pgm.Tag) + 1)).writeString(pgm.Tag)
	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// CopyData (F & B)

// ---------------------------------------------------------------------------------------------------------------------

// CopyDone

// ---------------------------------------------------------------------------------------------------------------------

// CopyFail

// ---------------------------------------------------------------------------------------------------------------------

// CopyInResponse

// ---------------------------------------------------------------------------------------------------------------------

// CopyOutResponse

// ---------------------------------------------------------------------------------------------------------------------

// CopyBothResponse

// ---------------------------------------------------------------------------------------------------------------------

// DataRow (B)

// Byte1('D')	Identifies the message as a data row.
// Int32		Length of message contents in bytes, including self.
// Int16		The number of column values that follow (possibly zero).

// Next, the following pair of fields appear for each column:

// Int32		The length of the column value, in bytes (this count does not include itself). Can be zero. As a special
//				case, -1 indicates a NULL column value. No value bytes follow in the NULL case.
// Byten		The value of the column, in the format indicated by the associated format code. n is the above length.

type _DataRowColumn struct {
	Value []byte
}

type _DataRow struct {
	Columns []*_DataRowColumn
}

func (pgm *_DataRow) write(m *_Messenger) error {

	// write the initial byte
	m.writeByte(DataRowMessageID)
	if m.Error != nil {
		return fmt.Errorf("_DataRow unable to write initial byte, err: %s", m.Error)
	}

	// work out the message length
	msgLen := int32(4 + 2) // msglen + column count = 6 bytes

	for _, c := range pgm.Columns {
		msgLen += int32(4) // column length
		msgLen += int32(len(c.Value))
	}
	m.writeInt32(msgLen)
	if m.Error != nil {
		return fmt.Errorf("_DataRow unable to write message length, err: %s", m.Error)
	}

	// write number of columns
	m.writeInt16(int16(len(pgm.Columns)))
	if m.Error != nil {
		return fmt.Errorf("_DataRow unable to write column count, err: %s", m.Error)
	}

	// write each column
	for _, c := range pgm.Columns {
		if c.Value == nil {
			m.writeInt32(-1)
			continue
		}
		m.writeInt32(int32(len(c.Value)))
		if len(c.Value) > 0 {
			m.writeByteArray(c.Value...)
		}
		if m.Error != nil {
			return fmt.Errorf("_DataRow unable to write column, err: %s", m.Error)
		}
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

// Describe (F)

// Byte1('D')	Identifies the message as a Describe command.
// Int32		Length of message contents in bytes, including self.
// Byte1		'S' to describe a prepared statement; or 'P' to describe a portal.
// String		The name of the prepared statement or portal to describe (an empty string selects the unnamed prepared
//				statement or portal).

type _Describe struct {
	Target     byte
	TargetName string
}

func (pgm *_Describe) read(m *_Messenger) error {

	// messageID / length already read

	// try and read the target byte
	pgm.Target = m.readByte()
	if m.Error != nil {
		return fmt.Errorf("_Describe unable to read target byte, err: %s", m.Error)
	}

	// try and read the target name
	pgm.TargetName = m.readString()
	if m.Error != nil {
		return fmt.Errorf("_Describe unable to read target name, err: %s", m.Error)
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

// EmptyQueryResponse

// ---------------------------------------------------------------------------------------------------------------------

// ErrorResponse (B)

// Byte1('E')	Identifies the message as an error.
// Int32		Length of message contents in bytes, including self.

// The message body consists of one or more identified fields, followed by a zero byte as a terminator. Fields can appear
// in any order. For each field there is the following:

// Byte1		A code identifying the field type; if zero, this is the message terminator and no string follows. The
//				presently defined field types are listed in Section 53.8. Since more field types might be added in
//				future, frontends should silently ignore fields of unrecognized type.
// String		The field value.

type _ErrorResponseField struct {
	Indicator byte
	Message   string
}

type _ErrorResponse struct {
	Fields []*_ErrorResponseField
}

func (pgm *_ErrorResponse) AddErrorField(indicator byte, message string) *_ErrorResponse {
	if pgm.Fields == nil {
		pgm.Fields = []*_ErrorResponseField{}
	}
	pgm.Fields = append(pgm.Fields, &_ErrorResponseField{indicator, message})
	return pgm
}

func (pgm *_ErrorResponse) write(m *_Messenger) error {

	// always starts with 'E'
	m.writeByte(ErrorResponseMessageID)

	// work out the total message length
	msgLen := int32(5)
	for _, f := range pgm.Fields {
		msgLen += int32(1 + len(f.Message) + 1)
	}
	m.writeInt32(msgLen)

	// then write the actual fields
	for _, f := range pgm.Fields {
		m.writeByte(f.Indicator)
		m.writeString(f.Message)
	}
	m.writeByte('\x00')

	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// Execute

// ---------------------------------------------------------------------------------------------------------------------

// Flush

// ---------------------------------------------------------------------------------------------------------------------

// FunctionCall

// ---------------------------------------------------------------------------------------------------------------------

// FunctionCallResponse

// ---------------------------------------------------------------------------------------------------------------------

// GSSResponse

// ---------------------------------------------------------------------------------------------------------------------

// NegotiateProtocolVersion

// ---------------------------------------------------------------------------------------------------------------------

// NoData

// ---------------------------------------------------------------------------------------------------------------------

// NoticeResponse

// ---------------------------------------------------------------------------------------------------------------------

// NotificationResponse

// ---------------------------------------------------------------------------------------------------------------------

// ParameterDescription (B)

// Byte1('t')	Identifies the message as a parameter description.
// Int32		Length of message contents in bytes, including self.
// Int16		The number of parameters used by the statement (can be zero).

// Then, for each parameter, there is the following:

// Int32		Specifies the object ID of the parameter data type.

type _ParameterDescription struct {
	ParameterOIDS []int32
}

func (pgm *_ParameterDescription) write(m *_Messenger) error {

	m.writeByte(ParameterDescriptionMessageID).
		writeInt32(int32(6 + (len(pgm.ParameterOIDS) * 4))).
		writeInt16(int16(len(pgm.ParameterOIDS)))
	if m.Error != nil {
		return fmt.Errorf("_ParameterDescription unable to write initial byte + length, err: %s", m.Error)
	}

	// write the parameter oids
	for _, p := range pgm.ParameterOIDS {
		m.writeInt32(p)
		if m.Error != nil {
			return fmt.Errorf("_ParameterDescription unable to write parameter OID, err: %s", m.Error)
		}
	}

	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// ParameterStatus

// ---------------------------------------------------------------------------------------------------------------------

// Parse (F)

// Byte1('P')	Identifies the message as a Parse command.
// Int32		Length of message contents in bytes, including self.
// String		The name of the destination prepared statement (an empty string selects the unnamed prepared statement).
// String		The query string to be parsed.
// Int16		The number of parameter data types specified (can be zero). Note that this is not an indication of the
//				number of parameters that might appear in the query string, only the number that the frontend wants to
//				prespecify types for.

// Then, for each parameter, there is the following:

// Int32		Specifies the object ID of the parameter data type. Placing a zero here is equivalent to leaving the type unspecified.

type _Parse struct {
	Statement     string
	SQL           string
	ParameterOIDs []int32
}

func (pgm *_Parse) read(m *_Messenger) error {

	// the message id and length are already read in

	// read the destination prepared statement
	pgm.Statement = m.readString()
	if m.Error != nil {
		return fmt.Errorf("_Parse unable to read statement, err: %s", m.Error)
	}

	// read the query
	pgm.SQL = m.readString()
	if m.Error != nil {
		return fmt.Errorf("_Parse unable to read query, err: %s", m.Error)
	}

	// work out n params
	nParameterOIDS := m.readInt16()
	if m.Error != nil {
		return fmt.Errorf("_Parse unable to read nParameterOIDS, err: %s", m.Error)
	}

	// iterate through and fetch them all
	pgm.ParameterOIDs = make([]int32, nParameterOIDS)
	for i := 0; i < int(nParameterOIDS); i++ {
		poid := m.readInt32()
		if m.Error != nil {
			return fmt.Errorf("_Parse unable to read parameterOID, err: %s", m.Error)
		}
		pgm.ParameterOIDs[i] = poid
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

// ParseComplete (B)

// Byte1('1')	Identifies the message as a Parse-complete indicator.
// Int32(4)	Length of message contents in bytes, including self.

type _ParseComplete struct{}

func (pgm *_ParseComplete) write(m *_Messenger) error {
	m.writeByte(ParseCompleteMessageID).writeInt32(4)
	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// PasswordMessage

// ---------------------------------------------------------------------------------------------------------------------

// PortalSuspended

// ---------------------------------------------------------------------------------------------------------------------

// Query (F)

// Byte1('Q')	Identifies the message as a simple query.
// Int32		Length of message contents in bytes, including self.
// String		The query string itself.

type _Query struct {
	SQL string
}

func (pgm *_Query) read(m *_Messenger) error {

	// Byte1 and int32 are already satisfied, just read the string
	pgm.SQL = m.readString()
	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// ReadyForQuery (B)

// Byte1('Z')	Identifies the message type. ReadyForQuery is sent whenever the backend is ready for a new query cycle.
// Int32(5)		Length of message contents in bytes, including self.
// Byte1		Current backend transaction status indicator. Possible values are 'I' if idle (not in a transaction
//				block); 'T' if in a transaction block; or 'E' if in a failed transaction block (queries will be rejected
//				until block is ended).

type _ReadyForQuery struct {
	Indicator byte
}

func (pgm *_ReadyForQuery) write(m *_Messenger) error {
	if pgm.Indicator == 0 {
		return fmt.Errorf("_ReadyForQuery.Indicator was empty, must be 'I', 'T' or 'E'")
	}
	if pgm.Indicator != 'I' && pgm.Indicator != 'T' && pgm.Indicator != 'E' {
		return fmt.Errorf("_ReadyForQuery.Indicator was invalid, found %s but expect 'I', 'T' or 'E'", string(pgm.Indicator))
	}
	m.writeByte(ReadyForQueryMessageID).writeInt32(5).writeByte(pgm.Indicator)
	return m.Error
}

// ---------------------------------------------------------------------------------------------------------------------

// RowDescription (B)

// Byte1('T') Identifies the message as a row description.
// Int32	  Length of message contents in bytes, including self.
// Int16	  Specifies the number of fields in a row (can be zero).

// Then, for each field, there is the following:

// String	The field name.
// Int32	If the field can be identified as a column of a specific table, the object ID of the table; otherwise zero.
// Int16	If the field can be identified as a column of a specific table, the attribute number of the column; otherwise zero.
// Int32	The object ID of the fields data type.
// Int16	The data type size (see pg_type.typlen). Note that negative values denote variable-width types.
// Int32	The type modifier (see pg_attribute.atttypmod). The meaning of the modifier is type-specific.
// Int16	The format code being used for the field. Currently will be zero (text) or one (binary). In a RowDescription
//			returned from the statement variant of Describe, the format code is not yet known and will always be zero.

type _RowDescriptionField struct {
	Name         string
	TableOID     int32
	ColumnAttr   int16
	DataTypeOID  int32
	DataTypeSize int16
	TypeModifier int32
	FormatCode   int16
}

func (f *_RowDescriptionField) Size() int32 {
	// name + \x00 + int32 * 3 + int16 * 3
	return int32(len(f.Name) + 1 + (4 * 3) + (2 * 3))
}

type _RowDescription struct {
	Fields []*_RowDescriptionField
}

func (pgm *_RowDescription) write(m *_Messenger) error {

	// write initial byte
	m.writeByte(RowDescriptionMessageID)
	if m.Error != nil {
		return fmt.Errorf("_RowDescription unable to write initial byte, err: %s", m.Error)
	}

	// work out the length
	msgLen := int32(6) // int32 + int16
	for _, f := range pgm.Fields {
		msgLen += f.Size()
	}
	m.writeInt32(msgLen)
	if m.Error != nil {
		return fmt.Errorf("_RowDescription unable to write message length, err: %s", m.Error)
	}

	// write the field count
	m.writeInt16(int16(len(pgm.Fields)))
	if m.Error != nil {
		return fmt.Errorf("_RowDescription unable to write field count length, err: %s", m.Error)
	}

	// write out the fields
	for _, f := range pgm.Fields {
		m.writeString(f.Name).
			writeInt32(f.TableOID).
			writeInt16(f.ColumnAttr).
			writeInt32(f.DataTypeOID).
			writeInt16(f.DataTypeSize).
			writeInt32(f.TypeModifier).
			writeInt16(f.FormatCode)
		if m.Error != nil {
			return fmt.Errorf("_RowDescription unable to write field, err: %s", m.Error)
		}
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

// SASLInitialResponse

// ---------------------------------------------------------------------------------------------------------------------

// SASLResponse

// ---------------------------------------------------------------------------------------------------------------------

// StartupMessage (F)

// Int32			Length of message contents in bytes, including self.
// Int32(196608)	The protocol version number. The most significant 16 bits are the major version number (3 for the protocol
// 					described here). The least significant 16 bits are the minor version number (0 for the protocol described here).

// The protocol version number is followed by one or more pairs of parameter name and value strings. A zero byte is
// required as a terminator after the last name/value pair. Parameters can appear in any order. user is required, others
// are optional. Each parameter is specified as:

// String		The parameter name. Currently recognized names are:
// user			The database user name to connect as. Required; there is no default.
// database		The database to connect to. Defaults to the user name.
// options		Command-line arguments for the backend. (This is deprecated in favor of setting individual run-time
// 				parameters.) Spaces within this string are considered to separate arguments, unless escaped with a
// 				backslash (\); write \\ to represent a literal backslash.
// replication	Used to connect in streaming replication mode, where a small set of replication commands can be issued
// 				instead of SQL statements. Value can be true, false, or database, and the default is false. See Section
// 				53.4 for details.

// In addition to the above, other parameters may be listed. Parameter names beginning with _pq_. are reserved for use as
// protocol extensions, while others are treated as run-time parameters to be set at backend start time. Such settings will
// be applied during backend start (after parsing the command-line arguments if any) and will act as session defaults.

// String			The parameter value.

type _StartupMessage struct {
	Parameters map[string]string
}

func (pgm *_StartupMessage) read(m *_Messenger) error {

	// this one is a bit special in that it has no Byte1 identifier
	// it's read from the message loop after the len/protocol are read

	// new up the parameter map
	pgm.Parameters = map[string]string{}

	// as len is gone and so is the protocol version, just grab the parameters
	for {
		k := m.readString()
		if m.Error != nil {
			return m.Error
		}
		if k == "" {
			break
		}
		v := m.readString()
		if m.Error != nil {
			return m.Error
		}
		if v == "" {
			break
		}

		pgm.Parameters[k] = v
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

// Sync ( Empty Message )

// ---------------------------------------------------------------------------------------------------------------------

// Terminate ( Empty Message )
