// main project main.go
package main

import (
	"bytes"
	_ "code.google.com/p/goprotobuf/proto"
	"encoding/binary"
	"fmt"
	_ "github.com/pebbe/zmq4"
)

type Person struct {
	Age  uint32
	Name string
}

func main() {
	fmt.Println("Hello World!")
	buf := new(bytes.Buffer)
	p := Person{100, "hello"}
	fmt.Println(p.Name)
	e := binary.Write(buf, binary.LittleEndian, p.Age)
	buf.Write([]byte(p.Name))
	fmt.Println("size of buf: ", buf.Len())
	fmt.Println(e)
	fmt.Printf("% x", buf.Bytes())
}
