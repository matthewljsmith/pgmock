package pgmock

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// ---------------------------------------------------------------------------------------------------------------------

type _SessionCancelRequestCallback func(*_CancelRequest)

type _Session struct {
	Key                 _SessionKey
	Messenger           *_Messenger
	CancelCallback      _SessionCancelRequestCallback
	IsHandshakeComplete bool
	Handler             MessageHandler
}

// ---------------------------------------------------------------------------------------------------------------------

func (session *_Session) handleCancel() {

	log.Warnf("Session.handleCancel")

	// write an error message
	errRes := &_ErrorResponse{Fields: []*_ErrorResponseField{
		&_ErrorResponseField{Indicator: ErrorSeverity, Message: "ERROR"},
		&_ErrorResponseField{Indicator: ErrorSQLStateCode, Message: SQLStateCodeQueryCanceled},
		&_ErrorResponseField{Indicator: ErrorMessage, Message: "Request was cancelled"},
	}}
	errRes.Fields = append(errRes.Fields)
	errRes.write(session.Messenger)

	// write a ReadyForQuery message
	err := (&_ReadyForQuery{Indicator: ReadyForQueryIdle}).write(session.Messenger)
	if err != nil {
		log.Fatalf("failed to write ReadyForQuery message, err: %s", err)
	}
	log.Infof("succesfully wrote ReadyForQuery message")
}

// ---------------------------------------------------------------------------------------------------------------------

func (session *_Session) processMessages() error {

	// loop until things break or get closed
	for {

		// if there's still no startup message we're in handshake mode
		if !session.IsHandshakeComplete {

			// if hsndshake returns true it was a cancel request
			if err := session.doHandshake(); err != nil {
				return err // kill connection
			}

			continue
		}

		// if we're past handshake, process the next message instead
		log.Infof("handshake is complete, processing next message")
		if err := session.processNextMessage(); err != nil {
			return err
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (session *_Session) doHandshake() error {

	log.Infof("handshakeComplete is false, attempting handshake")

	// shortcut
	m := session.Messenger

	// read and discard the message length
	msgLen := m.readInt32()
	if m.Error != nil {
		return fmt.Errorf("unable to read message length, err: %s", m.Error)
	}
	log.Infof("read message length: %d", msgLen)

	protocolVersion := m.readInt32()
	if m.Error != nil {
		return fmt.Errorf("unable to protocol version, err: %s", m.Error)
	}
	log.Infof("read protocol version: %d", protocolVersion)

	// work out the LSB/MSB
	lsb, msb := protocolVersion&0xFFFF, (protocolVersion>>16)&0xFFFF
	log.Infof("lsb(%d) msb(%d)", lsb, msb)

	// // SSLRequest uses the code 80877103 which yields 1234/5679 respectively ( 52.2.9 SSL Session Encryption )
	if msb == 1234 && lsb == 5679 {
		m.writeByte('N')
		return nil
	}

	// CancelRequest uses the code 80877102 which yields 1234/5678 respectively
	if msb == 1234 && lsb == 5678 {

		// read in the cancel request message
		cancelRequest := &_CancelRequest{}
		err := cancelRequest.read(m)
		if err != nil {
			return fmt.Errorf("unable to read cancel request, err: %s", err)
		}
		log.Infof("read cancel request, %v", cancelRequest)

		// kill the connection that sent the message ( postgres uses a separate connection to kill others )
		session.CancelCallback(cancelRequest)

		// quit the connection
		return fmt.Errorf("cancel request")
	}

	// StartupMessage is the only other option..
	if msb == 3 && lsb == 0 {

		log.Infof("have msb/lsb matching protocol 3.0 - accept startup message")

		// read the startup message
		msg := &_StartupMessage{}
		err := msg.read(m)
		if err != nil {
			return fmt.Errorf("unable to read startup message, err: %s", err)
		}

		log.Infof("read startup message: %v", msg)

		// we've done the handshake
		session.IsHandshakeComplete = true

		// write an AuthenticationOK message
		err = (&_AuthenticationOk{}).write(m)
		if err != nil {
			return fmt.Errorf("failed to write AuthenticationOK message, err: %s", err)
		}
		log.Infof("succesfully wrote AuthenticationOK message")

		// write the BackendKeyData message
		err = (&_BackendKeyData{
			ProcessID: session.Key.ProcessID,
			SecretKey: session.Key.SecretKey,
		}).write(m)
		if err != nil {
			return fmt.Errorf("failed to write BackendKeyData message, err: %s", err)
		}
		log.Infof("succesfully wrote BackendKeyData message")

		// write a ReadyForQuery message
		err = (&_ReadyForQuery{Indicator: 'I'}).write(m)
		if err != nil {
			return fmt.Errorf("failed to write ReadyForQuery message, err: %s", err)
		}
		log.Infof("succesfully wrote ReadyForQuery message")
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

func (session *_Session) processNextMessage() error {

	// shortcut
	m := session.Messenger

	// grab the messageID
	msgID := m.readByte()
	if m.Error != nil {
		return fmt.Errorf("unable to read messageID, err: %s", m.Error)
	}
	log.Infof("found message ID: %s", string(msgID))

	// get the message length and discard
	msgLen := m.readInt32()
	if m.Error != nil {
		return fmt.Errorf("unable to read message length, err: %s", m.Error)
	}
	log.Infof("found message length: %d", msgLen)

	// yep, big ol switch case for the message IDs, message registry would be nicer
	switch msgID {

	// terminate the connection straight away
	case TerminateMessageID:
		return fmt.Errorf("Terminate recieved")

	// pass Query on to Handler
	case QueryMessageID:
		err := session.Handler.HandleQuery(m)
		if err != nil {
			return fmt.Errorf("handling of Query message failed, err: %s", err)
		}

	case ParseMessageID:
		err := session.Handler.HandleParse(m)
		if err != nil {
			return fmt.Errorf("handling of Parse message failed, err: %s", err)
		}

	case DescribeMessageID:
		err := session.Handler.HandleDescribe(m)
		if err != nil {
			return fmt.Errorf("handling of Describe message failed, err: %s", err)
		}

	case SyncMessageID:
		err := (&_ReadyForQuery{Indicator: ReadyForQueryIdle}).write(m)
		if err != nil {
			return fmt.Errorf("handling of Sync message failed, err: %s", err)
		}
		log.Infof("wrote ReadyForQuery message")
	}

	return nil
}
