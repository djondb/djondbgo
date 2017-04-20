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

Here's a simple example on how to use the driver, for the complete documentation please visit: http://djondb.com/docs/?section=go-driver

	import (
		"github.com/djondb/djondbgo"
	)

	func main() {
		con := djondbgo.CreateConnection("localhost", 1243)
		con.Open()

		cur, err := con.ExecuteQuery("select * from testdb:testinsert")
		if err != nil {
			fmt.Println(err)
			return
		}
		res, err := cur.Next()
		if err != nil {
			t.Logf("Error %s", err)
			t.FailNow()
		}

		if res {
			fmt.Println("current: ", cur.current())
		}
	}

## Documentation

http://djondb.com/docs/?section=go-driver
