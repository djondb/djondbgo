## djondb go
========

This package is the pure Go implementation to connect to djondb, this only depends on https://github.com/satori/go.uuid to generate the uuids,
everything else is included in pure Go language, so it benefits from multi platform, etc

Note: please keep in mind that the connection object is not thread safe therefore you will need to create a DjondbConnection per thread

## Installation 

Use the `go` command:
   $ go get github.com/djondb/djondbgo

## Requirements

DjondbGO has been tested in Go 1.8

## Dependencies

Djondb GO depends on satori go.uuid, to retrieve it you can use:

	$ go get github.com/satori/go.uuid


## Example using DQL

import (
	djondb
)

func main() {
	con := CreateConnection("localhost", 1243)
	con.open()

	cur, err := con.executeQuery("select * from testdb:testinsert")
	if err != nil {
		fmt.Println(err)
		return
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

## Example using DQL
import (
	djondb
)

func main() {
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
