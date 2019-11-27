package pgmock

const (
	AuthenticationOkMessageID                byte = 'R'
	AuthenticationKerberosV5MessageID        byte = 'R'
	AuthenticationCleartextPasswordMessageID byte = 'R'
	AuthenticationMD5PasswordMessageID       byte = 'R'
	AuthenticationSCMCredentialMessageID     byte = 'R'
	AuthenticationGSSMessageID               byte = 'R'
	AuthenticationSSPIMessageID              byte = 'R'
	AuthenticationGSSContinueMessageID       byte = 'R'
	AuthenticationSASLMessageID              byte = 'R'
	AuthenticationSASLContinueMessageID      byte = 'R'
	AuthenticationSASLFinalMessageID         byte = 'R'
	BackendKeyDataMessageID                  byte = 'K'
	BindMessageID                            byte = 'B'
	BindCompleteMessageID                    byte = '2'
	CloseMessageID                           byte = 'C'
	CloseCompleteMessageID                   byte = '3'
	CommandCompleteMessageID                 byte = 'C'
	CopyDataMessageID                        byte = 'd'
	CopyDoneMessageID                        byte = 'c'
	CopyFailMessageID                        byte = 'f'
	CopyInResponseMessageID                  byte = 'G'
	CopyOutResponseMessageID                 byte = 'H'
	CopyBothResponseMessageID                byte = 'W'
	DataRowMessageID                         byte = 'D'
	DescribeMessageID                        byte = 'D'
	EmptyQueryResponseMessageID              byte = 'I'
	ErrorResponseMessageID                   byte = 'E'
	ExecuteMessageID                         byte = 'E'
	FlushMessageID                           byte = 'H'
	FunctionCallMessageID                    byte = 'F'
	FunctionCallResponseMessageID            byte = 'V'
	GSSResponseMessageID                     byte = 'p'
	NegotiateProtocolVersionMessageID        byte = 'v'
	NoDataMessageID                          byte = 'n'
	NoticeResponseMessageID                  byte = 'N'
	NotificationResponseMessageID            byte = 'A'
	ParameterDescriptionMessageID            byte = 't'
	ParameterStatusMessageID                 byte = 'S'
	ParseMessageID                           byte = 'P'
	ParseCompleteMessageID                   byte = '1'
	PasswordMessageMessageID                 byte = 'p'
	PortalSuspendedMessageID                 byte = 's'
	QueryMessageID                           byte = 'Q'
	ReadyForQueryMessageID                   byte = 'Z'
	RowDescriptionMessageID                  byte = 'T'
	SASLInitialResponseMessageID             byte = 'p'
	SASLResponseMessageID                    byte = 'p'
	SyncMessageID                            byte = 'S'
	TerminateMessageID                       byte = 'X'
)

const (
	ReadyForQueryIdle        byte = 'I'
	ReadyForQueryTransaction byte = 'T'
	ReadyForQueryError       byte = 'E'
)

const (
	CloseStatement byte = 'S'
	ClosePortal    byte = 'P'
)

const (
	ErrorSeverity         byte = 'S'
	ErrorSeverity9P       byte = 'V'
	ErrorSQLStateCode     byte = 'C'
	ErrorMessage          byte = 'M'
	ErrorDetail           byte = 'M'
	ErrorHint             byte = 'H'
	ErrorPosition         byte = 'P'
	ErrorInternalPosition byte = 'p'
	ErrorInternalQuery    byte = 'q'
	ErrorWhere            byte = 'W'
	ErrorSchema           byte = 's'
	ErrorTable            byte = 't'
	ErrorColumn           byte = 'c'
	ErrorDataType         byte = 'd'
	ErrorConstraint       byte = 'n'
	ErrorFile             byte = 'F'
	ErrorLine             byte = 'L'
	ErrorRoutine          byte = 'R'
)

const (
	SQLStateCodeQueryCanceled string = "57014"
)

const (
	OIDInt4 int32 = 23
)
