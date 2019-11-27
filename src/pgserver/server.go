package pgmock

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
)

var rowDescs = func() map[string]*_RowDescriptionField {
	return map[string]*_RowDescriptionField{
		"text": &_RowDescriptionField{
			TableOID:     0,
			ColumnAttr:   0,
			DataTypeOID:  25,
			DataTypeSize: -1,
			TypeModifier: 0,
			FormatCode:   0,
		},
		"int4": &_RowDescriptionField{
			TableOID:     0,
			ColumnAttr:   0,
			DataTypeOID:  23,
			DataTypeSize: -1,
			TypeModifier: 0,
			FormatCode:   0,
		},
	}
}()

// ---------------------------------------------------------------------------------------------------------------------

// Server wrapping interface around _Server
type Server interface {
	ListenAndServe(bindAddr string) error
	InjectQueryResponse(queryHash string, columns []string, rows [][]interface{}) error
}

// ---------------------------------------------------------------------------------------------------------------------

// _SessionKey used internall by the Server to key session instances
type _SessionKey struct {
	ProcessID int32
	SecretKey int32
}

// ---------------------------------------------------------------------------------------------------------------------

// _Server handles the incoming connections and session loops
type _Server struct {
	sync.Mutex
	Responder *_Responder
	Sessions  map[_SessionKey]*_Session
}

type _QueryResponse struct {
	Columns *_RowDescription
	Rows    []*_DataRow
}

type _Responder struct {
	sync.Mutex
	Responses map[string]*_QueryResponse
}

// ---------------------------------------------------------------------------------------------------------------------

// NewServer creates and returns a new Server
func NewServer() Server {
	return &_Server{
		Responder: &_Responder{Responses: map[string]*_QueryResponse{}},
		Sessions:  map[_SessionKey]*_Session{},
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (srv *_Server) InjectQueryResponse(hash string, cols []string, rows [][]interface{}) error {

	response := &_QueryResponse{
		Columns: &_RowDescription{
			Fields: []*_RowDescriptionField{},
		},
		Rows: []*_DataRow{},
	}

	// split out the columns to build the _RowDescription things
	for _, c := range cols {
		bits := strings.Split(c, ":")
		if _, found := rowDescs[bits[1]]; !found {
			log.Errorf("unable to find field type for %s", bits[1])
			return errors.New("unable to find field type")
		}
		field := rowDescs[bits[1]]
		field.Name = bits[0]
		response.Columns.Fields = append(response.Columns.Fields, field)
	}

	for _, r := range rows {
		cols := make([]*_DataRowColumn, 0)
		for _, rData := range r {
			switch v := rData.(type) {
			case string:
				cols = append(cols, &_DataRowColumn{Value: []byte(v)})
			case float64:
				cols = append(cols, &_DataRowColumn{Value: []byte(fmt.Sprintf("%f", v))})
			default:
				cols = append(cols, &_DataRowColumn{Value: v.([]byte)})
			}
		}
		response.Rows = append(response.Rows, &_DataRow{Columns: cols})
	}

	// save the response off
	srv.Responder.Lock()
	srv.Responder.Responses[hash] = response
	srv.Responder.Unlock()

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

func (srv *_Server) ListenAndServe(bindAddr string) error {

	// create a new listener
	ln, err := net.Listen("tcp", bindAddr)
	if err != nil {
		return nil
	}

	// close the listener when the service quits
	defer ln.Close()

	// infinite loop
	for {

		// accept connection on port
		conn, err := ln.Accept()
		if err != nil {
			log.Errorf("unable to accept connection, err: %s", err)
		}
		log.Infof("accepted connection from: %s", conn.RemoteAddr())

		// create a new session instance
		session := &_Session{
			Key:            _SessionKey{rand.Int31(), rand.Int31()},
			Messenger:      &_Messenger{Stream: conn},
			CancelCallback: srv.issueCancelRequest,
			Handler:        &_BaseHandler{srv.Responder},
		}

		// add it to the active server list
		srv.Lock()
		srv.Sessions[session.Key] = session
		srv.Unlock()

		// spin off the goroutine to handle the session message processing
		go func(session *_Session) {

			// do some stuff on return
			defer func() {

				// make sure the connection closes
				log.Infof("closed connection to %s", conn.RemoteAddr())
				conn.Close()

				// remove the session from the server
				srv.Lock()
				delete(srv.Sessions, session.Key)
				srv.Unlock()

				// catch panics, don't want to crash the server
				if err := recover(); err != nil {
					log.Errorf("handleConnection panicked, err %s", err)
				}
			}()

			// start the event loop on the session
			if err := session.processMessages(); err != nil {
				log.Errorf("session.processMessages returned an error: %s", err)
				return
			}

		}(session)
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (srv *_Server) issueCancelRequest(req *_CancelRequest) {

	// maintain concurrency
	srv.Lock()
	defer srv.Unlock()

	log.Warnf("Server.issueCancelRequest %v", req)
	if v, found := srv.Sessions[_SessionKey{req.ProcessID, req.SecretKey}]; found {
		v.handleCancel()
	}
}
