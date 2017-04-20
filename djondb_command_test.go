package djondbgo

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestCommandDropNamespace(t *testing.T) {
	net := CreateNetwork();
	net.connect("localhost", 1243)

	cmd := CreateCommand()
	cmd.dropNamespace(net, "testdb", "testdrop")
}

func TestCommandShowDBs(t *testing.T) {
	net := CreateNetwork();
	net.connect("localhost", 1243)

	cmd := CreateCommand()
	dbs, _ := cmd.showDbs(net)
	fmt.Println(dbs)
}

func TestCommandShowNamespaces(t *testing.T) {
	net := CreateNetwork();
	net.connect("localhost", 1243)

	cmd := CreateCommand()
	nss, _ := cmd.showNamespaces(net, "testdb")
	fmt.Println(nss)
}

func TestCommandInsert(t *testing.T) {
	net := CreateNetwork()
	net.connect("localhost", 1243)

	cmd := CreateCommand()
	o := make(map[string]interface{})

	o["name"] = "John"
	o["lastName"] = "Smith"
	o["age"] = 10
	o["b"] = true
	o["salary"] = 12232.232

	inner := make(map[string]interface{})
	inner["test"] = "blah"
	o["inner"] = inner

	arr := make([]map[string]interface{}, 1, 1)
	element := make(map[string]interface{})
	element["arr"] = "Hello"
	arr[0] = element
	o["arr"] = arr
	cmd.insert(net, "testdb", "testinsert", o)
}

func TestCommandFind(t *testing.T) {
	net := CreateNetwork()
	net.connect("localhost", 1243)

	cmd := CreateCommand()

	cur, _ := cmd.find(net, "testdb", "testinsert", "*", "")

	fmt.Println(cur.Next())
}

func TestCommandInsertUpdateRemove(t *testing.T) {
	net := CreateNetwork()
	net.connect("localhost", 1243)

	cmd := CreateCommand()
	
	o := make(map[string]interface{})

	o["name"] = "John"
	o["lastName"] = "Smith"
	o["age"] = 10
	o["b"] = true
	o["salary"] = 12232.232

	cmd.insert(net, "testdb", "testupdate", o)

	c, _ := cmd.find(net, "testdb", "testupdate", "*", "")

	n, err := c.Next()
	if !n || err != nil {
		t.Logf("Error retrieving last insert")
		t.FailNow()
	}

	pou := c.Current()
	if pou == nil {
		t.Logf("Error current is null after find insert")
		t.FailNow()
	}
	ou := *pou
	ou["newvalue"] = 123
	cmd.update(net, "testdb", "testupdate", ou)

	c2, _ := cmd.find(net, "testdb", "testupdate", "*", "_id = '" + ou["_id"].(string) + "'")
	n, err = c2.Next()
	if !n || err != nil {
		t.Logf("Error id not found")
		t.FailNow()
	}
	pounew := c2.Current()
	if pounew == nil {
		t.Logf("Error c2 current nil")
		t.FailNow()
	}
	ounew := *pounew
	newvalue := ounew["newvalue"].(int32)
	if int(newvalue) != int(123) {
		t.Logf("Error newvalue is not 123")
		t.FailNow()
	}

	cmd.remove(net, "testdb", "testupdate", ounew["_id"].(string), ounew["_revision"].(string))
	fmt.Println("removed: ", ounew["_id"])
	c3, _ := cmd.find(net, "testdb", "testupdate", "*", "_id = '" + ounew["_id"].(string) + "'")
	n, err = c3.Next()
	if n || err != nil {
		t.Logf("Error id not removed")
		t.FailNow()
	}
}

func TestCommandTX(t *testing.T) {
	net := CreateNetwork()
	net.connect("localhost", 1243)

	cmd := CreateCommand()
	
	o := make(map[string]interface{})

	o["name"] = "John"
	o["lastName"] = "Smith"
	o["age"] = 10
	o["b"] = true
	o["salary"] = 12232.232

	cmd.insert(net, "testdb", "testtx", o)

	c, _ := cmd.find(net, "testdb", "testtx", "*", "")

	n, err := c.Next()
	if !n || err != nil {
		t.Logf("Error retrieving last insert")
		t.FailNow()
	}

	pou := c.Current()
	if pou == nil {
		t.Logf("Error current is null after find insert")
		t.FailNow()
	}
	cmd.beginTransaction()
	ou := *pou
	ou["newvalue"] = 123
	cmd.update(net, "testdb", "testtx", ou)

	c2, _ := cmd.find(net, "testdb", "testtx", "*", "_id = '" + ou["_id"].(string) + "'")
	n, err = c2.Next()
	if !n || err != nil {
		t.Logf("Error id not found")
		t.FailNow()
	}
	pounew := c2.Current()
	if pounew == nil {
		t.Logf("Error c2 current nil")
		t.FailNow()
	}
	ounew := *pounew
	newvalue := ounew["newvalue"].(int32)
	if int(newvalue) != int(123) {
		t.Logf("Error newvalue is not 123")
		t.FailNow()
	}
	cmd.commitTransaction(net)

	c2, _ = cmd.find(net, "testdb", "testtx", "*", "_id = '" + ou["_id"].(string) + "'")
	n, err = c2.Next()
	if !n || err != nil {
		t.Logf("Error id not found")
		t.FailNow()
	}
	pounew = c2.Current()
	if pounew == nil {
		t.Logf("Error c2 current nil")
		t.FailNow()
	}
	ounew = *pounew
	newvalue = ounew["newvalue"].(int32)
	if int(newvalue) != int(123) {
		t.Logf("Error newvalue is not 123")
		t.FailNow()
	}
	ounew["newvalue"] = 124
	cmd.beginTransaction()
	cmd.update(net, "testdb", "testtx", ounew)
	c2, _ = cmd.find(net, "testdb", "testtx", "*", "_id = '" + ou["_id"].(string) + "'")
	n, err = c2.Next()
	if !n || err != nil {
		t.Logf("Error id not found")
		t.FailNow()
	}
	pounew = c2.Current()
	if pounew == nil {
		t.Logf("Error c2 current nil")
		t.FailNow()
	}
	ounew = *pounew
	newvalue = ounew["newvalue"].(int32)
	if int(newvalue) != int(124) {
		t.Logf("Error newvalue is not 124")
		t.FailNow()
	}
	cmd.rollbackTransaction(net)
	c2, _ = cmd.find(net, "testdb", "testtx", "*", "_id = '" + ou["_id"].(string) + "'")
	n, err = c2.Next()
	if !n || err != nil {
		t.Logf("Error id not found")
		t.FailNow()
	}
	pounew = c2.Current()
	if pounew == nil {
		t.Logf("Error c2 current nil")
		t.FailNow()
	}
	ounew = *pounew
	newvalue = ounew["newvalue"].(int32)
	if int(newvalue) != int(123) {
		t.Logf("Error newvalue is not 123")
		t.FailNow()
	}
}

func TestCreateIndex(t *testing.T) {
	net := CreateNetwork()
	net.connect("localhost", 1243)

	cmd := CreateCommand()

	indexDef := make(map[string]interface{})

	indexDef["db"] = "testdb"
	indexDef["ns"] = "testns"
	indexDef["name"] = "testindex"

	fields := make([]map[string]interface{}, 1, 1)

	var field map[string]interface{}
	field = make(map[string]interface{})
	field["path"] = "name"
	fields[0] = field
	indexDef["fields"] = fields

	err := cmd.createIndex(net, indexDef)
	if err != nil {
		t.Logf("Error %s", err)
		t.FailNow()
	}
}

func TestCommandBackup(t *testing.T) {
	net := CreateNetwork()
	net.connect("localhost", 1243)

	cmd := CreateCommand()
	cmd.backup(net, "testdb", "temp.tar")
}

func TestCommandExecuteQuery(t *testing.T) {
	net := CreateNetwork()
	net.connect("localhost", 1243)

	cmd := CreateCommand()

	cur, err := cmd.executeQuery(net, "select * from testdb:testinsert")
	if err != nil {
		t.Logf("Error %s", err)
		t.FailNow()
	}
	res, err := cur.Next()
	if err != nil {
		t.Logf("Error %s", err)
		t.FailNow()
	}

	if res {
		fmt.Println("current: ", cur.Current())
	}
}

func TestCommandInsertUpdateRemoveDQL(t *testing.T) {
	net := CreateNetwork()
	net.connect("localhost", 1243)

	cmd := CreateCommand()
	
	o := make(map[string]interface{})

	o["name"] = "John"
	o["lastName"] = "Smith"
	o["age"] = 10
	o["b"] = true
	o["salary"] = 12232.232
	b, _ := json.Marshal(o)
	
	insertDQL := fmt.Sprintf("insert %s into testdb:testupdate", b)

	_, err := cmd.executeUpdate(net, insertDQL)
	if err != nil {
		t.Logf("Error %s", err)
		t.FailNow()
	}

	findAllDQL := "select * from testdb:testupdate"
	c, _ := cmd.executeQuery(net, findAllDQL)

	n, err := c.Next()
	if !n || err != nil {
		t.Logf("Error retrieving last insert")
		t.FailNow()
	}

	pou := c.Current()
	if pou == nil {
		t.Logf("Error current is null after find insert")
		t.FailNow()
	}
	ou := *pou
	ou["newvalue"] = 123

	b, _ = json.Marshal(ou)
	updateDQL := fmt.Sprintf("update %s into testdb:testupdate", b)
	cmd.executeUpdate(net, updateDQL)

	findByIdDQL := fmt.Sprintf("select * from testdb:testupdate where _id = '%s'", ou["_id"].(string))
	c2, _ := cmd.executeQuery(net, findByIdDQL)
	n, err = c2.Next()
	if !n || err != nil {
		t.Logf("Error id not found")
		t.FailNow()
	}
	pounew := c2.Current()
	if pounew == nil {
		t.Logf("Error c2 current nil")
		t.FailNow()
	}
	ounew := *pounew
	newvalue := ounew["newvalue"].(int32)
	if int(newvalue) != int(123) {
		t.Logf("Error newvalue is not 123")
		t.FailNow()
	}

	removeDQL := fmt.Sprintf("remove '%s' with '%s' from testdb:testupdate", ounew["_id"].(string), ounew["_revision"].(string))
	cmd.executeUpdate(net, removeDQL)
	fmt.Println("removed: ", ounew["_id"])
	c3, _ := cmd.executeQuery(net, findByIdDQL)
	n, err = c3.Next()
	if n || err != nil {
		t.Logf("Error id not removed")
		t.FailNow()
	}
}

