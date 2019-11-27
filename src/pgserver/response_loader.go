package pgmock

import (
	"strings"
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/nanobox-io/golang-scribble"
)

type _ResponseLoader struct{}

type QueryResponse struct {
	Columns *_RowDescription
	Rows    []*_DataRow
}

type QueryData struct {
	Columns []string        `json:"columns"`
	Rows    [][]interface{} `json:"rows"`
}

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
	}
}()

func (rl *_ResponseLoader) GetQueryResponse(hash string) (*QueryResponse, error) {

	// a new scribble driver, providing the directory where it will be writing to,
	// and a qualified logger if desired
	db, err := scribble.New(".data", nil)
	if err != nil {
		log.Errorf("unable to open scribble database, err: %s", err)
		return nil, err
	}

	var qData QueryData
	if err := db.Read("responses", hash, &qData); err != nil {
		log.Errorf("unable to read response for query hash %s, err: %s", hash, err)
		return nil, err
	}

	response := &QueryResponse{
		Columns: &_RowDescription{
			Fields: []*_RowDescriptionField{},
		},
		Rows: []*_DataRow{},
	}

	// split out the columns to build the _RowDescription things
	for _, c := range qData.Columns {
		bits := strings.Split(c, ":")
		if _, found := rowDescs[bits[1]]; !found {
			log.Errorf("unable to find field type for %s", bits[1])
			return nil, errors.New("unable to find field type")
		}
		field := rowDescs[bits[1]]
		field.Name = bits[0]
		response.Columns.Fields = append(response.Columns.Fields, field)
	}	

	for _, r := range qData.Rows {
		cols := make([]*_DataRowColumn, 0)
		for  _, rData := range r {
			cols = append(cols, &_DataRowColumn{Value: rData.([]byte)})
		}
		response.Rows = append(response.Rows, &_DataRow{Columns: cols})
	}

	return response, nil
}
