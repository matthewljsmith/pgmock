package pgmock

import (
	"math/rand"
	"net"
	"sync"

	log "github.com/sirupsen/logrus"
)

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
	Sessions map[_SessionKey]*_Session
}

// ---------------------------------------------------------------------------------------------------------------------

func NewServer() *_Server {
	return &_Server{
		Sessions: map[_SessionKey]*_Session{},
	}
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
			Handler:        &_BaseHandler{&_ResponseLoader{}},
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
