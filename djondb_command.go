package djondbgo

import (
	"fmt"
	"github.com/satori/go.uuid"
	"time"
)

const (
	COMMANDTYPE_INSERT = 0
	COMMANDTYPE_UPDATE = 1
	COMMANDTYPE_FIND = 2
	COMMANDTYPE_CLOSECONNECTION = 3
	COMMANDTYPE_DROPNAMESPACE = 4
	COMMANDTYPE_SHUTDOWN = 5
	COMMANDTYPE_SHOWDBS = 6
	COMMANDTYPE_SHOWNAMESPACES = 7
	COMMANDTYPE_REMOVE = 8
	COMMANDTYPE_COMMIT = 9
	COMMANDTYPE_ROLLBACK = 10
	COMMANDTYPE_FETCHCURSOR = 11
	COMMANDTYPE_FLUSHBUFFER = 12
	COMMANDTYPE_CREATEINDEX = 13
	COMMANDTYPE_BACKUP = 14
	COMMANDTYPE_RCURSOR = 15
	COMMANDTYPE_PERSISTINDEXES = 16
	COMMANDTYPE_EXECUTEQUERY = 17
	COMMANDTYPE_EXECUTEUPDATE = 18
)

type Command struct {
	activeTransactionId string
	resultCode int32
	resultMessage string
}

func CreateCommand() *Command {
	return &Command{}
}

func (cmd *Command) writeHeader(net *Network) {
	version := "3.5.60822"
	net.writeString(version)
}

func (cmd *Command) writeOptions(net *Network) {
	options := make(map[string]interface{})
	if cmd.activeTransactionId != "" {
		options["_transactionId"] = cmd.activeTransactionId;
	}
	net.writeBSON(options)
}

func (cmd *Command) readErrorInformation(net *Network) (err error) {
	cmd.resultCode = net.readInt()
   if cmd.resultCode > 0 {
		cmd.resultMessage = net.readString()
		err = fmt.Errorf("Error: %d. Message: %s", cmd.resultCode, cmd.resultMessage)
	}
	return
}

func (cmd *Command) readResultDropNamespace(net *Network) (res bool, err error) {
	net.readInt()

	err = cmd.readErrorInformation(net)
	res = true

	return
}

func (cmd *Command) dropNamespace(net *Network, db string, ns string) (res bool, err error) {
	net.reset()
	cmd.writeHeader(net)
	net.writeInt(COMMANDTYPE_DROPNAMESPACE)
	cmd.writeOptions(net)
	net.writeString(db)
	net.writeString(ns)
	net.flush()
	res, err = cmd.readResultDropNamespace(net)
	return
}

func (cmd *Command) readResultShowDbs(net *Network) (res []string, err error) {
	results := int(net.readInt())
	dbs := make([]string, results, results)
	for x := 0; x < results; x++ {
		dbs[x] = net.readString()
	}

	err = cmd.readErrorInformation(net)
	res = dbs
	return
}

func (cmd *Command) showDbs(net *Network) (res []string, err error) {
	net.reset()
	cmd.writeHeader(net)
	net.writeInt(COMMANDTYPE_SHOWDBS)
	cmd.writeOptions(net)
	net.flush()
	res, err = cmd.readResultShowDbs(net)
	return
}

func (cmd *Command) readResultShowNamespaces(net *Network) (res []string, err error) {
	results := int(net.readInt())
	nss := make([]string, results, results)
	for x := 0; x < results; x++ {
		nss[x] = net.readString()
	}

	err = cmd.readErrorInformation(net)
	res = nss
	return
}

func (cmd *Command) showNamespaces(net *Network, db string) (res []string, err error) {
	net.reset()
	cmd.writeHeader(net)
	net.writeInt(COMMANDTYPE_SHOWNAMESPACES)
	cmd.writeOptions(net)
	net.writeString(db)
	net.flush()
	res, err = cmd.readResultShowNamespaces(net)
	return
}

func (cmd *Command) readResultInsert(net *Network) (res int32, err error) {
	res = net.readInt()
	err = cmd.readErrorInformation(net)
	return
}

func (cmd *Command) insert(net *Network, db string, ns string, data map[string]interface{}) (res int32, err error) {
	net.reset()
	cmd.writeHeader(net)
	net.writeInt(COMMANDTYPE_INSERT) 
	cmd.writeOptions(net)
	net.writeString(db)
	net.writeString(ns)
	net.writeBSON(data)
	net.flush()
	res, err = cmd.readResultInsert(net)
	return
}

func (cmd *Command) readResultFetchRecords(net *Network) (res []map[string]interface{}, err error) {
	flag := net.readInt();
	if flag == 1 {
		res = net.readBSONArray()
	}
	
	err = cmd.readErrorInformation(net)

	return
}

func (cmd *Command) fetchRecords(net *Network, cursorId string) (res []map[string]interface{}, err error) {
	net.reset()
	cmd.writeHeader(net)
	net.writeInt(COMMANDTYPE_FETCHCURSOR)
	cmd.writeOptions(net)
	net.writeString(cursorId)
	net.flush()
	res, err = cmd.readResultFetchRecords(net)
	return
}

func (cmd *Command) readResultFind(net *Network) (result DjondbCursor, err error) {
	cursorId := net.readString()
	flag := net.readInt()
	var results []map[string]interface{}
	if flag == 1 {
		tmp := net.readBSONArray()
		results = tmp
	}

	result = *CreateCursor(net, &cursorId, results)
	err = cmd.readErrorInformation(net)
	return
}

func (cmd *Command) find(net *Network, db string, ns string, sselect string, filter string) (res DjondbCursor, err error) {
	net.reset()
	cmd.writeHeader(net)
	net.writeInt(COMMANDTYPE_FIND)
	cmd.writeOptions(net)
	net.writeString(db)
	net.writeString(ns)
	net.writeString(filter)
	net.writeString(sselect)
	net.flush()
	res, err = cmd.readResultFind(net)
	return
}

func (cmd *Command) readResultUpdate(net *Network) (result bool, err error) {
	result = net.readBoolean()
	err = cmd.readErrorInformation(net)

	return 
}

func (cmd *Command) update(net *Network, db string, ns string, data map[string]interface{}) (res bool, err error) {
	net.reset()
	cmd.writeHeader(net)
	net.writeInt(COMMANDTYPE_UPDATE)
	cmd.writeOptions(net)
	net.writeString(db)
	net.writeString(ns)
	net.writeBSON(data)
	net.flush()
	res, err = cmd.readResultUpdate(net)
	return
}

func (cmd *Command) readResultRemove(net *Network) (res bool, err error) {
	res = net.readBoolean()
	err = cmd.readErrorInformation(net)
	return 
}


func (cmd *Command) remove(net *Network, db string, ns string, id string, revision string) (res bool, err error) {
	net.reset()
	cmd.writeHeader(net)
	net.writeInt(COMMANDTYPE_REMOVE)
	cmd.writeOptions(net)
	net.writeString(db)
	net.writeString(ns)
	net.writeString(id)
	net.writeString(revision)

	net.flush()
	res, err =  cmd.readResultRemove(net)
	return
}

func (cmd *Command) beginTransaction() (err error) {
	u4 := uuid.NewV4()
	cmd.activeTransactionId = fmt.Sprintf("%s", u4)
	return
}

func (cmd *Command) readResultCommitTransaction(net *Network) (err error) {
	err = cmd.readErrorInformation(net)
	return
}

func (cmd *Command) commitTransaction(net *Network) (err error) {
	if cmd.activeTransactionId != "" {
		net.reset()
		cmd.writeHeader(net)
		net.writeInt(COMMANDTYPE_COMMIT)
		cmd.writeOptions(net)
		net.writeString(cmd.activeTransactionId)
		net.flush()

		err = cmd.readResultCommitTransaction(net)
		cmd.activeTransactionId = ""
	} else {
		err = fmt.Errorf("Nothing to commit, you need beginTransaction before committing or rollback")
	}
	return
}

func (cmd *Command) readResultRollbackTransaction(net *Network) (err error) {
	err = cmd.readErrorInformation(net)
	return
}

func (cmd *Command) rollbackTransaction(net *Network) (err error) {
	if cmd.activeTransactionId != "" {
		net.reset()
		cmd.writeHeader(net)
		net.writeInt(COMMANDTYPE_ROLLBACK)
		cmd.writeOptions(net)
		net.writeString(cmd.activeTransactionId)
		net.flush()

		err = cmd.readResultRollbackTransaction(net)
		cmd.activeTransactionId = ""
	} else {
		err = fmt.Errorf("Nothing to rollback, you need beginTransaction before committing or rollback")
	}	
	return
}

func (cmd *Command) readResultCreateIndex(net *Network) (err error) {
	err = cmd.readErrorInformation(net)
	return
}

func (cmd *Command) createIndex(net *Network, indexDef map[string]interface{}) (err error) {
	net.reset()
	cmd.writeHeader(net)
	net.writeInt(COMMANDTYPE_CREATEINDEX)
	cmd.writeOptions(net)
	net.writeBSON(indexDef)
	net.flush()
	err = cmd.readResultCreateIndex(net)
	return
}

func (cmd *Command) readResultBackup(net *Network) (result int32, err error) {
	result =  net.readInt()
	err = cmd.readErrorInformation(net)
	return 
}

func (cmd *Command) backup(net *Network, db string, destFile string) (res int32, err error) {
	net.reset()
	cmd.writeHeader(net)
	net.writeInt(COMMANDTYPE_BACKUP)
	cmd.writeOptions(net)
	net.writeString(db)
	net.writeString(destFile)
	net.flush()
	res, err = cmd.readResultBackup(net)
	return
}

func (cmd *Command) executeQuery(net *Network, query string) (cursorResult DjondbCursor, err error) {
	net.reset()
	cmd.writeHeader(net)
	net.writeInt(COMMANDTYPE_EXECUTEQUERY)
	cmd.writeOptions(net)
	net.writeString(query)
	net.flush()

	flag := net.readInt()
	var cursorTemp *DjondbCursor

	if flag == 1 {
		commandType := net.readInt();
		switch commandType {
			case COMMANDTYPE_INSERT:
				_, err = cmd.readResultInsert(net)

			case COMMANDTYPE_UPDATE:
				_, err = cmd.readResultUpdate(net)

			case COMMANDTYPE_FIND:
				var c DjondbCursor
				c, err = cmd.readResultFind(net)
				cursorTemp = &c

			case COMMANDTYPE_DROPNAMESPACE:
				_, err = cmd.readResultDropNamespace(net)

			case COMMANDTYPE_SHOWDBS:
				dbs, _ := cmd.readResultShowDbs(net)
				err = cmd.readErrorInformation(net)
				arrDbs := make([]map[string]interface{}, len(dbs), len(dbs))
				for x := 0; x < len(dbs); x++ {
					row := make(map[string]interface{})
					row["db"] = dbs[x]
					arrDbs[x] = row
				}
				cursorTemp = CreateCursor(net, nil, arrDbs)

			case COMMANDTYPE_SHOWNAMESPACES:
				nss, _ := cmd.readResultShowNamespaces(net)
				err = cmd.readErrorInformation(net)
				arrNs := make([]map[string]interface{}, len(nss), len(nss))
				for x := 0; x < len(nss); x++ {
					row := make(map[string]interface{})
					row["ns"] = nss[x]
					arrNs[x] = row
				}
				cursorTemp = CreateCursor(net, nil, arrNs)
			case COMMANDTYPE_REMOVE:
				_, err = cmd.readResultRemove(net)
			case COMMANDTYPE_COMMIT:
				err = cmd.readResultCommitTransaction(net)
				cmd.activeTransactionId = ""

			case COMMANDTYPE_ROLLBACK:
				err = cmd.readResultRollbackTransaction(net)
				cmd.activeTransactionId = ""

			case COMMANDTYPE_CREATEINDEX:
				err = cmd.readResultCreateIndex(net)

			case COMMANDTYPE_BACKUP:
				_, err = cmd.readResultBackup(net)
		}
	} else {
		err = cmd.readErrorInformation(net)
	}

	if cursorTemp == nil {
		arr := make([]map[string]interface{}, 1, 1)
		row := make(map[string]interface{})
		row["date"] = fmt.Sprintf(time.Now().String())
		row["success"] = true
		arr[0] = row
		cursorTemp = CreateCursor(net, nil, arr)
	}
	cursorResult = *cursorTemp

	return
}

func (cmd *Command) executeUpdate(net *Network, query string) (result bool, err error) {
	net.reset()
	cmd.writeHeader(net)
	net.writeInt(COMMANDTYPE_EXECUTEUPDATE)
	cmd.writeOptions(net)
	net.writeString(query)
	net.flush()

	flag := net.readInt()
	if flag == 1 {
		commandType := net.readInt();
		switch commandType {
			case COMMANDTYPE_INSERT:
				_, err = cmd.readResultInsert(net)
				result = (err == nil)

			case COMMANDTYPE_UPDATE:
				result, err = cmd.readResultUpdate(net)
				result = (err == nil)

			case COMMANDTYPE_DROPNAMESPACE:
				result, err = cmd.readResultDropNamespace(net)

			case COMMANDTYPE_REMOVE:
				result, err =  cmd.readResultRemove(net)

			case COMMANDTYPE_COMMIT:
				err = cmd.readResultCommitTransaction(net)
				cmd.activeTransactionId = ""
				result = (err == nil)

			case COMMANDTYPE_ROLLBACK:
				err= cmd.readResultRollbackTransaction(net)
				cmd.activeTransactionId = ""
				result = (err == nil)

			case COMMANDTYPE_CREATEINDEX:
				err = cmd.readResultCreateIndex(net)
				result = (err == nil)

			case COMMANDTYPE_BACKUP:
				_, err = cmd.readResultBackup(net)
				result = (err == nil)
		}
	} else {
		err = cmd.readErrorInformation(net)
	}
	return
}


