package djondbgo

import (
	"testing"
)

func TestInsert(t *testing.T) {
	net := CreateNetwork();
	net.connect("localhost", 1243)
/* Header **/
	net.writeString("3.6.33333")
/* Command Type */
	net.writeInt(0) // INSERT
/* options */
	var options map[string]interface{}
	options = make(map[string]interface{})
	net.writeBSON(options) // Options

/* data */
	net.writeString("testdb")
	net.writeString("testns")
	var data map[string]interface{}
	data = make(map[string]interface{})
	data["name"] = "John"
	data["lastName"] = "Smith"
	net.writeBSON(data)

	net.flush()

	t.Logf("Done");
}

func TestShowDbs(t *testing.T) {
	net := CreateNetwork();
	net.connect("localhost", 1243)
/* Header **/
	net.writeString("3.6.33333")
/* Command Type */
	net.writeInt(6) // SHOWDBS
/* options */
	var options map[string]interface{}
	options = make(map[string]interface{})
	net.writeBSON(options) // Options
	net.flush()

	dbs := net.readInt()
	t.Logf("Dbs: %d", dbs)
	db := net.readString()
	t.Logf("DB " + db)
	db = net.readString()
	t.Logf("DB " + db)
	t.Logf("Done");
}
