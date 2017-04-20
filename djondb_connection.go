package djondb

import (
	"fmt"
)

type DjondbConnection struct {
	host string
	port int
	cmd Command
	net Network
}


func (con *DjondbConnection) checkError(cmd *Command) (err error) {
	if con.cmd.resultCode > 0 {
		err = fmt.Errorf("Error (err error) { %d, Message (err error) { %s", con.cmd.resultCode, con.cmd.resultMessage)
	}
	return
}

func CreateConnection(host string, port int) (DjondbConnection) {
	con := &DjondbConnection{}
	con.host = host
	con.port = port
	con.cmd = *CreateCommand()
	con.net = *CreateNetwork();
	return *con
}

func (con *DjondbConnection) open() (err error) {
	con.net.connect(con.host, con.port)
	return
}

func (con *DjondbConnection) showDbs() (res []string, err error) {
	res, err = con.cmd.showDbs(&con.net)
	err = con.checkError(&con.cmd)
	return 
}

func (con *DjondbConnection) showNamespaces(db string) (res []string, err error) {
	res, _ = con.cmd.showNamespaces(&con.net, db)
	err = con.checkError(&con.cmd)
	return
}

func (con *DjondbConnection) insert(db string, ns string, data map[string]interface{}) (err error) {
	con.cmd.insert(&con.net, db, ns, data)
	err = con.checkError(&con.cmd)
	return
}

func (con *DjondbConnection) update(db string, ns string, data map[string]interface{}) (err error) {
	con.cmd.update(&con.net, db, ns, data)
	err = con.checkError(&con.cmd)
	return 
}

func (con *DjondbConnection) find(db string, ns string, sselect string, filter string) (cur DjondbCursor, err error) {
	cur, _ = con.cmd.find(&con.net, db, ns, sselect, filter)
	err = con.checkError(&con.cmd)
	return
}

func (con *DjondbConnection) dropNamespace(db string, ns string) (err error) {
	con.cmd.dropNamespace(&con.net, db, ns)
	err = con.checkError(&con.cmd)
	return 
}

func (con *DjondbConnection) remove(db string, ns string, id string, revision string) (err error) {
	con.cmd.remove(&con.net, db, ns, id, revision)
	err = con.checkError(&con.cmd)
	return 
}

func (con *DjondbConnection) beginTransaction() (err error) {
	con.cmd.beginTransaction()
	return
}
	
func (con *DjondbConnection) commitTransaction() (err error) {
	con.cmd.commitTransaction(&con.net)
	err = con.checkError(&con.cmd)
	return
}
	
func (con *DjondbConnection) rollbackTransaction() (err error) {
	con.cmd.rollbackTransaction(&con.net)
	err = con.checkError(&con.cmd)
	return
}

func (con *DjondbConnection) createIndex(indexDef map[string]interface{}) (err error) {
	con.cmd.createIndex(&con.net, indexDef)
	err = con.checkError(&con.cmd)
	return
}

func (con *DjondbConnection) executeQuery(query string) (res DjondbCursor, err error) {
	res, _ = con.cmd.executeQuery(&con.net, query)
	err = con.checkError(&con.cmd)
	return
}

func (con *DjondbConnection) executeUpdate(query string) (res bool, err error) {
	res, _ = con.cmd.executeUpdate(&con.net, query)
	err = con.checkError(&con.cmd)
	return
}

func (con *DjondbConnection) backup(db string, destination string) (err error) {
	con.cmd.backup(&con.net, db, destination)
	err = con.checkError(&con.cmd)
	return
}

