package djondbgo

import (
	"net"
	"fmt"
	"bytes"
	"encoding/binary"
)

type Network struct {
	buffer bytes.Buffer;
	bufferLen int32;
	bufferPos int32;
	conn net.Conn;
}

func CreateNetwork() *Network {
	return &Network {}
}

func (dnet *Network) connect(host string, port int) (err error) {
	address := fmt.Sprintf("%s:%d", host, port)
	dnet.bufferLen = 0
	dnet.bufferPos = 0
	dnet.buffer.Reset()
	dnet.conn, err = net.Dial("tcp", address)
	return
}

func (dnet *Network) writeString(data string) {
	bytes := []byte(data)
	dnet.writeInt(int32(len(bytes)))
	l, _ := dnet.buffer.Write(bytes)
	dnet.bufferLen += int32(l)
}

func (dnet *Network) writeInt(data int32) {
	binary.Write(&dnet.buffer, binary.LittleEndian, data)
	dnet.bufferLen += 4
}

func (dnet *Network) writeLong(data int64) {
	binary.Write(&dnet.buffer, binary.LittleEndian, data)
	dnet.bufferLen +=8 
}

func (dnet *Network) writeDouble(data float64) {
	binary.Write(&dnet.buffer, binary.LittleEndian, data)
	dnet.bufferLen +=8 
}

func (dnet *Network) writeBoolean(data bool) {
	var d []byte = make([]byte, 1, 1)
	if data {
		d[0] = 1
	} else {
		d[0] = 0
	}
	dnet.buffer.Write(d)
	dnet.bufferLen +=1
}

func (dnet *Network) writeBSONArray(arr []map[string]interface{}) {
	dnet.writeLong(int64(len(arr)))
	for i := 0; i < len(arr); i++ {
		b := arr[i]

		dnet.writeBSON(b)
	}
}

func (dnet *Network) writeBSON(data map[string]interface{}) {
	dnet.writeLong(int64(len(data)))
	for k, v := range data {
		dnet.writeString(k)
		switch v.(type) {
			case int:
				dnet.writeLong(0)
				dnet.writeInt(int32(v.(int)))
			case int32:
				dnet.writeLong(0)
				dnet.writeInt(v.(int32))
			case float64:
				dnet.writeLong(1)
				dnet.writeDouble(v.(float64))
			case int64:
				dnet.writeLong(2)
				dnet.writeLong(v.(int64))
			case string:
				dnet.writeLong(4)
				dnet.writeString(v.(string))
			case map[string]interface{}:
				dnet.writeLong(5)
				dnet.writeBSON(v.(map[string]interface{}))
			case bool:
				dnet.writeLong(10)
				dnet.writeBoolean(v.(bool))
			case []map[string]interface{}:
				dnet.writeLong(6)
				dnet.writeBSONArray(v.([]map[string]interface{}))
		}
	}
}

func (dnet *Network) flush() {
	dnet.conn.Write(dnet.buffer.Bytes())
	dnet.reset()
}

func (dnet *Network) check(size int32) {
	if dnet.bufferLen < dnet.bufferPos + size {
		var buffer []byte = make([]byte, 1024*100, 1024*100)
		read, _ := dnet.conn.Read(buffer)
		dnet.buffer.Write(buffer)
		dnet.bufferLen += int32(read)
	}
}

func (dnet *Network) readInt() int32 {
	dnet.check(4)
	var result int32
	binary.Read(&dnet.buffer, binary.LittleEndian, &result)

	dnet.bufferPos += 4

	return result
}

func (dnet *Network) readLong() int64 {
	dnet.check(8)
	var result int64
	binary.Read(&dnet.buffer, binary.LittleEndian, &result)

	dnet.bufferPos += 8 
	return result
}

func (dnet *Network) readDouble() float64 {
	dnet.check(8)
	var result float64
	binary.Read(&dnet.buffer, binary.LittleEndian, &result)

	dnet.bufferPos += 8 
	return result
}

func (dnet *Network) readBoolean() bool {
	dnet.check(1)

	b, _ := dnet.buffer.ReadByte()
	dnet.bufferPos += 1 
	if b == 0 {
		return true
	} else {
		return false
	}
}

func (dnet *Network) readChars(size int32) string {
	dnet.check(size)
	var b []byte = make([]byte, size, size)
	dnet.buffer.Read(b)
	var result string = string(b)
	dnet.bufferPos += size

	return result
}

func (dnet *Network) readString() string {
	size := dnet.readInt()
	result := dnet.readChars(size)
	return result
}

func (dnet *Network) readBSONArray() []map[string]interface{} {
	size := dnet.readLong()
	var result []map[string]interface{} = make([]map[string]interface{}, size, size)
	var i int64
	for i = 0; i < size; i++ {
		o := dnet.readBSON()
		result[i] = o
	}
	return result
}

func (dnet *Network) readBSON() map[string]interface{} {
	elements := dnet.readLong()

	var result map[string]interface{}
	result = make(map[string]interface{})
	for i := int64(0); i < elements; i++ {
		key := dnet.readString()
		datatype := dnet.readLong()
		var val interface{}

		switch datatype {
			case 0:
				val = dnet.readInt()
			case 1:
				val = dnet.readDouble()
			case 2:
				val = dnet.readLong()
			case 4:
				val = dnet.readString()
			case 5:
				val = dnet.readBSON()
			case 6:
				val = dnet.readBSONArray()
			case 10:
				val = dnet.readBoolean()
		}
		result[key] = val
	}
	return result
}

func (dnet *Network) reset() {
	dnet.bufferLen = 0
	dnet.bufferPos = 0
	dnet.buffer.Reset()
}
