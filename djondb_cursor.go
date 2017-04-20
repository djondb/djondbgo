package djondbgo

import (
	"fmt"
)
const (
	CS_LOADING = 1
	CS_RECORDS_LOADED = 2
	CS_CLOSED = 3
)

type DjondbCursor struct {
	net *Network
	cursorId *string
	rows []map[string]interface{}
	position int32
	_current *map[string]interface{}
	count int32
	status int
}

func CreateCursor(net *Network, cursorId *string, firstPage []map[string]interface{}) *DjondbCursor {
	result := &DjondbCursor{}
	result.net = net
	result.cursorId = cursorId
	result.rows = firstPage
	result.position = 0
	result._current = nil
	result.count = int32(len(result.rows))

	if cursorId == nil {
		result.status = CS_RECORDS_LOADED
	} else {
		result.status = CS_LOADING
	}
	return result
}

func (cursor *DjondbCursor) Next() (res bool, err error) {
	res = false
	if cursor.status == CS_CLOSED {
		err = fmt.Errorf("Cursor is closed")
		return
	}
	res = true
	if cursor.count > cursor.position {
		row := cursor.rows[cursor.position]
		cursor._current = &row
		cursor.position += 1
	} else {
		if cursor.status == CS_LOADING {
			cmd := CreateCommand()
			var page []map[string]interface{}
			page, err = cmd.fetchRecords(cursor.net, *cursor.cursorId)
			if len(page) == 0 {
				cursor.status = CS_RECORDS_LOADED
				res = false
			} else {
				cursor.rows = append(cursor.rows, page...)
				cursor.count = int32(len(cursor.rows))
				res, err = cursor.Next()
			}
		} else {
			res = false
		}
	}

	return
}

func (cursor *DjondbCursor) Previous() (bool, *string) {
	if cursor.status == CS_CLOSED {
		e := "Cursor is closed"
		return false, &e
	}
	result := true
	if cursor.count > 0 && cursor.position > 0 {
		cursor.position -= 1
		cursor._current = &cursor.rows[cursor.position]
	} else {
		result = false
	}
	return result, nil
}

func (cursor *DjondbCursor) Current() *map[string]interface{} {
	return cursor._current
}
