package djondb

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestConnectionDropNamespace(t *testing.T) {
	con := CreateConnection("localhost", 1243)
	con.open()

	con.dropNamespace("testdb", "testdrop")
}

func TestConnectionShowDBs(t *testing.T) {
	con := CreateConnection("localhost", 1243)
	con.open()

	dbs, _ := con.showDbs()
	fmt.Println(dbs)
}

func TestConnectionShowNamespaces(t *testing.T) {
	con := CreateConnection("localhost", 1243)
	con.open()

	nss, _ := con.showNamespaces("testdb")
	fmt.Println(nss)
}

func TestConnectionInsert(t *testing.T) {
	con := CreateConnection("localhost", 1243)
	con.open()

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
	con.insert("testdb", "testinsert", o)
}

func TestConnectionFind(t *testing.T) {
	con := CreateConnection("localhost", 1243)
	con.open()

	cur, _ := con.find("testdb", "testinsert", "*", "")

	fmt.Println(cur.next())
}

func TestConnectionInsertUpdateRemove(t *testing.T) {
	con := CreateConnection("localhost", 1243)
	con.open()

	o := make(map[string]interface{})

	o["name"] = "John"
	o["lastName"] = "Smith"
	o["age"] = 10
	o["b"] = true
	o["salary"] = 12232.232

	con.insert("testdb", "testupdate", o)

	c, _ := con.find("testdb", "testupdate", "*", "")

	n, err := c.next()
	if !n || err != nil {
		t.Logf("Error retrieving last insert")
		t.FailNow()
	}

	pou := c.current()
	if pou == nil {
		t.Logf("Error current is null after find insert")
		t.FailNow()
	}
	ou := *pou
	ou["newvalue"] = 123
	con.update("testdb", "testupdate", ou)

	c2, _ := con.find("testdb", "testupdate", "*", "_id = '" + ou["_id"].(string) + "'")
	n, err = c2.next()
	if !n || err != nil {
		t.Logf("Error id not found")
		t.FailNow()
	}
	pounew := c2.current()
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

	con.remove("testdb", "testupdate", ounew["_id"].(string), ounew["_revision"].(string))
	fmt.Println("removed: ", ounew["_id"])
	c3, _ := con.find("testdb", "testupdate", "*", "_id = '" + ounew["_id"].(string) + "'")
	n, err = c3.next()
	if n || err != nil {
		t.Logf("Error id not removed")
		t.FailNow()
	}
}

func TestConnectionTX(t *testing.T) {
	con := CreateConnection("localhost", 1243)
	con.open()
	
	o := make(map[string]interface{})

	o["name"] = "John"
	o["lastName"] = "Smith"
	o["age"] = 10
	o["b"] = true
	o["salary"] = 12232.232

	con.insert("testdb", "testtx", o)

	c, _ := con.find("testdb", "testtx", "*", "")

	n, err := c.next()
	if !n || err != nil {
		t.Logf("Error retrieving last insert")
		t.FailNow()
	}

	pou := c.current()
	if pou == nil {
		t.Logf("Error current is null after find insert")
		t.FailNow()
	}
	con.beginTransaction()
	ou := *pou
	ou["newvalue"] = 123
	con.update("testdb", "testtx", ou)

	c2, _ := con.find("testdb", "testtx", "*", "_id = '" + ou["_id"].(string) + "'")
	n, err = c2.next()
	if !n || err != nil {
		t.Logf("Error id not found")
		t.FailNow()
	}
	pounew := c2.current()
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
	con.commitTransaction()

	c2, _ = con.find("testdb", "testtx", "*", "_id = '" + ou["_id"].(string) + "'")
	n, err = c2.next()
	if !n || err != nil {
		t.Logf("Error id not found")
		t.FailNow()
	}
	pounew = c2.current()
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
	con.beginTransaction()
	con.update("testdb", "testtx", ounew)
	c2, _ = con.find("testdb", "testtx", "*", "_id = '" + ou["_id"].(string) + "'")
	n, err = c2.next()
	if !n || err != nil {
		t.Logf("Error id not found")
		t.FailNow()
	}
	pounew = c2.current()
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
	con.rollbackTransaction()
	c2, _ = con.find("testdb", "testtx", "*", "_id = '" + ou["_id"].(string) + "'")
	n, err = c2.next()
	if !n || err != nil {
		t.Logf("Error id not found")
		t.FailNow()
	}
	pounew = c2.current()
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

func TestConnectionCreateIndex(t *testing.T) {
	con := CreateConnection("localhost", 1243)
	con.open()

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

	err := con.createIndex(indexDef)
	if err != nil {
		t.Logf("Error %s", err)
		t.FailNow()
	}
}

func TestConnectionBackup(t *testing.T) {
	con := CreateConnection("localhost", 1243)
	con.open()

	con.backup("testdb", "temp.tar")
}

func TestConnectionExecuteQuery(t *testing.T) {
	con := CreateConnection("localhost", 1243)
	con.open()

	cur, err := con.executeQuery("select * from testdb:testinsert")
	if err != nil {
		t.Logf("Error %s", err)
		t.FailNow()
	}
	res, err := cur.next()
	if err != nil {
		t.Logf("Error %s", err)
		t.FailNow()
	}

	if res {
		fmt.Println("current: ", cur.current())
	}
}

func TestConnectionInsertUpdateRemoveDQL(t *testing.T) {
	con := CreateConnection("localhost", 1243)
	con.open()
	
	o := make(map[string]interface{})

	o["name"] = "John"
	o["lastName"] = "Smith"
	o["age"] = 10
	o["b"] = true
	o["salary"] = 12232.232
	b, _ := json.Marshal(o)
	
	insertDQL := fmt.Sprintf("insert %s into testdb:testupdate", b)

	_, err := con.executeUpdate(insertDQL)
	if err != nil {
		t.Logf("Error %s", err)
		t.FailNow()
	}

	findAllDQL := "select * from testdb:testupdate"
	c, _ := con.executeQuery(findAllDQL)

	n, err := c.next()
	if !n || err != nil {
		t.Logf("Error retrieving last insert")
		t.FailNow()
	}

	pou := c.current()
	if pou == nil {
		t.Logf("Error current is null after find insert")
		t.FailNow()
	}
	ou := *pou
	ou["newvalue"] = 123

	b, _ = json.Marshal(ou)
	updateDQL := fmt.Sprintf("update %s into testdb:testupdate", b)
	con.executeUpdate(updateDQL)

	findByIdDQL := fmt.Sprintf("select * from testdb:testupdate where _id = '%s'", ou["_id"].(string))
	c2, _ := con.executeQuery(findByIdDQL)
	n, err = c2.next()
	if !n || err != nil {
		t.Logf("Error id not found")
		t.FailNow()
	}
	pounew := c2.current()
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
	con.executeUpdate(removeDQL)
	fmt.Println("removed: ", ounew["_id"])
	c3, _ := con.executeQuery(findByIdDQL)
	n, err = c3.next()
	if n || err != nil {
		t.Logf("Error id not removed")
		t.FailNow()
	}
}
