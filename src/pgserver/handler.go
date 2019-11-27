package pgmock

import (
	"crypto/sha1"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type MessageHandler interface {
	HandleQuery(*_Messenger) error
	HandleParse(*_Messenger) error
	HandleDescribe(*_Messenger) error
}

type _BaseHandler struct{
	ResponseLoader *_ResponseLoader
}

func (bh *_BaseHandler) HandleDescribe(m *_Messenger) error {

	msg := &_Describe{}
	err := msg.read(m)
	if err != nil {
		return err
	}
	log.Infof("HandleDesribe(%s, %s)", string(msg.Target), msg.TargetName)

	// if it's a statement issue ParameterDescription, RowDescription
	err = (&_ParameterDescription{ParameterOIDS: []int32{OIDInt4}}).write(m)
	if err != nil {
		return err
	}
	log.Info("WroteParameterDescription")

	return nil
}

func (bh *_BaseHandler) HandleParse(m *_Messenger) error {

	msg := &_Parse{}
	msg.read(m)
	if m.Error != nil {
		return m.Error
	}
	log.Infof("HandleParse(%s)", msg.SQL)

	err := (&_ParseComplete{}).write(m)
	if err != nil {
		return err
	}

	return nil

}

func (bh *_BaseHandler) HandleQuery(m *_Messenger) error {

	msg := &_Query{}
	msg.read(m)
	if m.Error != nil {
		return m.Error
	}
	log.Infof("HandleQuery(%s)", msg.SQL)

	// hash the query string to work out what data to send
	h := sha1.New()
	h.Write([]byte(msg.SQL))
	hash := fmt.Sprintf("%X", h.Sum(nil))

	response, err := bh.ResponseLoader.GetQueryResponse(hash)
	if err != nil {
		return err
	}

	if err := response.Columns.write(m); err != nil {
		log.Errorf("unable to write columns, err: %s", err)
		return err
	}

	for _, row := range response.Rows {
		if err := row.write(m); err != nil {
			log.Errorf("unable to write rows, err: %s", err)
			return err
		}
	}

	complete := &_CommandComplete{}
	complete.selectOrCreate(1)
	err = complete.write(m)
	if err != nil {
		return err
	}

	log.Infof("Wrote CommandComplete")

	(&_ReadyForQuery{Indicator: 'I'}).write(m)

	return nil
}
